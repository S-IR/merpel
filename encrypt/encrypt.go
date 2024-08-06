package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// EncryptData encrypts the data using AES GCM.
// It accepts a byte array (data), and a byte array (key).
// It returns the encrypted byte array and an error.
func EncryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, data, nil)
	return append(nonce, ciphertext...), nil
}

// DecryptData decrypts the data using AES GCM.
// It accepts the encrypted byte array, a byte array (key), and the nonce.
// It returns the decrypted byte array and an error.
func DecryptData(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return nil, errors.New("invalid ciphertext")
	}

	nonce := encrypted[:nonceSize]
	ciphertext := encrypted[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// EncryptFile takes a large file an encrypts it using AES GCM. It handles the encryption of both small and  large files.
func EncryptFile(file io.Reader, writer io.Writer, key []byte) error {

	buffer := make([]byte, 1024*1024) // 1 MB buffer
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			// Encrypt the chunk
			encryptedChunk, err := encryptChunk(buffer[:n], key)
			if err != nil {
				return err
			}

			// Write the encrypted chunk
			_, err = writer.Write(encryptedChunk)
			if err != nil {
				return err
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// EncryptChunk encrypts a chunk of data using AES GCM.
func encryptChunk(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, data, nil)
	return append(nonce, ciphertext...), nil
}

// ValidateAESKey checks if the provided key is valid for AES encryption.
// Valid key lengths are 16, 24, or 32 bytes.
func IsValidAESKey(key []byte) bool {
	switch len(key) {
	case 16, 24, 32:
		return true
	default:
		return false
	}
}

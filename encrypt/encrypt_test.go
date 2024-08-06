package encrypt_test

import (
	"bytes"
	"crypto/rand"
	"os"
	"testing"

	"github.com/s-ir/merpel/encrypt"
	"github.com/s-ir/merpel/mockups"
)

func TestEncrypt(t *testing.T) {
	// Create a random file
	path, err := mockups.CreateDeterministicFile(".", 1000*10) // Creating a 10KB file for testing
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(path)

	// Read the contents of the file
	data, err := os.ReadFile(path)
	if err != nil {
		t.Error(err)
	}

	// Define a random key for encryption
	key := make([]byte, 32) // AES-256 key length
	if _, err := rand.Read(key); err != nil {
		t.Error(err)
	}

	// Encrypt the data
	encrypted, err := encrypt.EncryptData(data, key)
	if err != nil {
		t.Errorf("Encryption error: %v", err)
	}

	// Decrypt the data
	decrypted, err := encrypt.DecryptData(encrypted, key)
	if err != nil {
		t.Errorf("Decryption error: %v", err)
	}

	// Compare the original and decrypted data
	if !bytes.Equal(data, decrypted) {
		t.Errorf("Decrypted data does not match original")
	}
}

package mockups

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// fileCounter keeps track of the number of files created
var fileCounter uint64

// CreateDeterministicFile generates a file with deterministic content of the specified size in KB.
func CreateDeterministicFile(path string, size uint64) (filename string, err error) {
	// Use the size as the seed for the random number generator
	r := rand.New(rand.NewSource(int64(size)))

	sizeInBytes := size * 1024

	// Generate a deterministic extension based on the size
	extensions := []string{".txt", ".png", ".svg", ".json", ".xml"}
	extension := extensions[r.Intn(len(extensions))]

	// Generate a unique identifier based on the size and file creation counter
	uniqueID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d-%d", size, fileCounter))))
	fileCounter++ // Increment the file counter for the next file
	filename = fmt.Sprintf("temp-%s%s", uniqueID[:6], extension)

	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return "", err
	}
	defer file.Close()

	const bufferSize = 1024
	buffer := make([]byte, bufferSize)
	for written := uint64(0); written < sizeInBytes; {
		remaining := sizeInBytes - written
		if remaining < bufferSize {
			for i := 0; i < int(remaining); i++ {
				buffer[i] = byte(r.Intn(256))
			}
			_, err = file.Write(buffer[:remaining])
		} else {
			for i := 0; i < bufferSize; i++ {
				buffer[i] = byte(r.Intn(256))
			}
			_, err = file.Write(buffer)
		}
		if err != nil {
			return "", err
		}
		written += uint64(len(buffer))
	}

	return filename, err
}

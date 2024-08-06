package mockups

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

var folderCounter uint64

func CreateDeterministicFolder(path string, size uint64) (folderPath string, err error) {
	r := rand.New(rand.NewSource(int64(size)))

	uniqueID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d-%d", size, folderCounter))))
	folderCounter++
	foldername := fmt.Sprintf("temp-%s", uniqueID[:6])
	folderPath = filepath.Join(path, foldername)
	err = os.MkdirAll(folderPath, 0777)
	if err != nil {
		return "", err
	}

	if size == 0 {
		return folderPath, err
	}

	const MIN_FOLDER_SIZE = 10 * 1024 // 10 KB
	for written := uint64(0); written < size; {
		remaining := size - written

		// there's a 15% chance that a nested folder will be created inside
		if r.Float32() < 0.15 && remaining > MIN_FOLDER_SIZE {
			// random folder size is either the remaining size
			randFolderSize := min(remaining, uint64(r.Intn(9)+1)*10*1024)
			_, err = CreateDeterministicFolder(folderPath, randFolderSize)
			if err != nil {
				return "", err
			}
			written += randFolderSize
			continue
		}

		// this is between 1 and 10 MB
		randFileSize := uint64((r.Intn(9) + 1) * 1024 * 1024)

		if remaining < randFileSize {
			_, err := CreateDeterministicFile(folderPath, remaining/1024)
			if err != nil {
				return "", err
			}
			written += remaining
		} else {
			_, err := CreateDeterministicFile(folderPath, randFileSize/1024)
			if err != nil {
				return "", err
			}
			written += randFileSize
		}
	}
	return folderPath, nil
}

// min is a helper function to return the smaller of two uint64 numbers
func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

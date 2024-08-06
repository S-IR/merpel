package lib

import (
	"io/fs"
	"os"
	"path/filepath"
)

// os.MkdirAll followed by os.Create on the filepath
func MkFileAll(p string, perm fs.FileMode) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}

	return os.Create(p)
}

package fileHandlers

import (
	"hash/fnv"
	"io"

	"github.com/s-ir/merpel/lib"
)

var h = fnv.New128a()

func HashFile(file io.Reader, filename string) ([]byte, error) {
	h := fnv.New128a() // Create a new hash for each call

	lib.Assert(len(filename) > 0, "filename cannot be empty")
	lib.Assert(file != nil, "file cannot be nil")
	if _, err := h.Write([]byte(filename)); err != nil {
		return nil, err
	}

	buffer := make([]byte, 1024*1024)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			if _, err := h.Write(buffer[:n]); err != nil {
				return nil, err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return h.Sum(nil), nil
}

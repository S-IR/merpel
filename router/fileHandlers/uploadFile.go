package fileHandlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/s-ir/merpel/cfg"
	"github.com/s-ir/merpel/encrypt"
	"github.com/s-ir/merpel/lib"
	"github.com/s-ir/merpel/pbs"
	"google.golang.org/protobuf/proto"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("user").(*pbs.User)
	if !ok {
		panic("this endpoint should have been authenticated")
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = os.MkdirAll("./uploads", 0755)
	if err != nil {
		panic(err)
	}

	metadata := r.FormValue("metadata")
	if metadata == "" {
		http.Error(w, "Missing metadata for your given file", http.StatusBadRequest)
		return
	}

	var postFileRequest pbs.PostFileRequest
	err = proto.Unmarshal([]byte(metadata), &postFileRequest)
	if err != nil {
		http.Error(w, "Error parsing proto data", http.StatusBadRequest)
		return
	}

	postFileRequest.Path = filepath.Clean(postFileRequest.Path)
	hash, err := HashFile(file, postFileRequest.Path)
	if err != nil {
		http.Error(w, "Error storing file", http.StatusInternalServerError)
		return
	}
	lib.Assert(len(hash) > 0, "hash cannot be empty")

	var realEncryptionKey []byte
	givenKey := postFileRequest.GetEncryptionKey()
	if len(givenKey) == 0 {
		realEncryptionKey = cfg.ENCRYPTION_KEY
	} else {
		isValid := encrypt.IsValidAESKey(givenKey)
		if !isValid {
			http.Error(w, "Invalid encryption key, key must be 16, 24 or 32 bytes long", http.StatusBadRequest)
			return
		}

		realEncryptionKey = givenKey
	}
	lib.Assert(len(realEncryptionKey) > 0, "encryption key cannot be empty")

	finalPath := filepath.Join("./_merple", user.Id, "files", filepath.Dir(postFileRequest.Path), string(hash))
	outFile, err := lib.MkFileAll(finalPath, cfg.RW_ONLY_BY_APP)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = encrypt.EncryptFile(file, outFile, realEncryptionKey)
	if err != nil {
		panic(err)
	}
	// Send a response indicating success
	w.WriteHeader(http.StatusOK)

}

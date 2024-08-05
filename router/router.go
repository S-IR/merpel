package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/s-ir/merpel/pbs"
)

func RouterInit() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Post("/files", UploadHandler)
	// r.Delete("/files/{filename}", DeleteHandler)
	// r.Get("/hello", helloHandler)
	fmt.Println("Merple : A simple file storage database LISTENING ON :19113")
	return r
}

// // helloHandler handles requests to the /hello endpoint
// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Hello, World!\n")
// }

func UploadHandler(w http.ResponseWriter, r *http.Request) {
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

	var fileInfo pbs.PostFileRequest
	err = proto.Unmarshal([]byte(metadata), &fileInfo)
	if err != nil {
		http.Error(w, "Error parsing proto data", http.StatusBadRequest)
		return
	}

	fileInfo.Path = filepath.Clean(fileInfo.Path)
	os.MkdirAll("./uploads/"+filepath.Dir(fileInfo.Path), 0755)
	outFile, err := os.Create("./uploads/" + fileInfo.Path)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		panic(err)
	}

	// Send a response indicating success
	w.WriteHeader(http.StatusOK)

}

func DeleteHandler() {

}

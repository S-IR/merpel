package router

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/s-ir/merpel/pbs"
	"google.golang.org/protobuf/proto"
)

func TestUploadHandler(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "testfile.txt")
	if err != nil {
		t.Fatalf("Error creating form file: %v", err)
	}

	_, err = part.Write([]byte("This is a test file"))
	if err != nil {
		t.Fatalf("Error writing to form file: %v", err)
	}

	fileInfo := &pbs.PostFileRequest{
		Path:       "./hello/example.txt",
		Permission: 0644,
	}

	protoData, err := proto.Marshal(fileInfo)
	if err != nil {
		fmt.Println("Error encoding proto:", err)
		return
	}

	protoPart, err := writer.CreateFormField("metadata")
	if err != nil {
		t.Fatalf("Error creating proto form field: %v", err)
	}

	_, err = protoPart.Write(protoData)
	if err != nil {
		t.Fatalf("Error writing proto data: %v", err)
	}

	writer.Close()
	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()
	// Call the handler
	UploadHandler(recorder, req)

	resp := recorder.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", resp.Status)
	}

	os.RemoveAll("./uploads")
}

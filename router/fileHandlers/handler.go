package fileHandlers

import (
	"github.com/go-chi/chi"
	"github.com/s-ir/merpel/router/auth"
)

func AttachFileHandler(r *chi.Mux) {
	r.Post("/upload", UploadHandler)
	r.Use(auth.Middleware)
}

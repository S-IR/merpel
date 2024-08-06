package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/s-ir/merpel/router/fileHandlers"
)

func RouterInit() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	fileHandlers.AttachFileHandler(r)

	// r.Delete("/files/{filename}", DeleteHandler)
	// r.Get("/hello", helloHandler)
	return r
}

// // helloHandler handles requests to the /hello endpoint
//
//	func helloHandler(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprintf(w, "Hello, World!\n")
//	}

package main

import (
	"log"
	"net/http"

	"github.com/s-ir/merpel/router"
)

func main() {
	r := router.RouterInit()
	log.Fatal(http.ListenAndServe(":19113", r))

}

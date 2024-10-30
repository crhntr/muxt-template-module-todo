package main

import (
	"cmp"
	"log"
	"net/http"
	"os"

	"github.com/crhntr/muxt-template-module-htmx-sqlc/internal/hypertext"
)

func main() {
	srv := &hypertext.Server{}
	mux := http.NewServeMux()
	hypertext.TemplateRoutes(mux, srv)
	log.Fatal(http.ListenAndServe(":"+cmp.Or(os.Getenv("PORT"), "8080"), mux))
}

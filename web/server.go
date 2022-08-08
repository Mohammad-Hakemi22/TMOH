package web

import (
	"log"
	"net/http"

	"github.com/Mohammad-Hakemi22/tmoh/routes"
)

func RunServer() {
	r := routes.Router()

	log.Fatalln(http.ListenAndServe(":8080", r))
}

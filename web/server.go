package web

import (
	"log"
	"net/http"
)

func RunServer() {
	r := Router()

	log.Fatalln(http.ListenAndServe(":8000", r))
}

package web

import (
	"log"
	"net/http"

	"github.com/Mohammad-Hakemi22/tmoh/routes"
	"github.com/Mohammad-Hakemi22/tmoh/config"
)

func RunServer() {
	r := routes.Router()
	log.Println("Listen And Serve on Port", config.AppConfig.Port)
	log.Fatalln(http.ListenAndServe(config.AppConfig.Port, r))
}

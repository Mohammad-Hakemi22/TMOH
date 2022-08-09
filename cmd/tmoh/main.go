package main

import (
	"fmt"

	"github.com/Mohammad-Hakemi22/tmoh/config"
	"github.com/Mohammad-Hakemi22/tmoh/web"
)

func init() {
	config.SetConfig()
}

func main() {
	fmt.Println("Hello to the Earth")
	web.RunServer()
}

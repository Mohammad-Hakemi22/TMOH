package main

import (
	"github.com/Mohammad-Hakemi22/tmoh/config"
	"github.com/Mohammad-Hakemi22/tmoh/web"
)

func init() {
	config.SetConfig()
}

func main() {
	web.RunServer()
}

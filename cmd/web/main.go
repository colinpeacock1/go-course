package main

import (
	"fmt"
	"net/http"
	"web-hello-world-v2/pkg/config"

	"github.com/colinpeacock1/go-course/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

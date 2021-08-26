package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/shanduur/simple-srv"
)

func main() {
	srv := server.New(":8080")

	srv.MainRouter.HandleFunc("/hello-world", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello World!"))
	})

	srv.MainRouter.HandleFunc("/hello-error", func(rw http.ResponseWriter, r *http.Request) {
		server.PrintError(rw, fmt.Errorf("error obtained"), 501)
	})

	if err := srv.Run(); err != nil {
		log.Fatalf("fatal error during running server: %s", err.Error())
	}
}

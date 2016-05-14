package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gokrokvertskhov/serverlib/apiserver"
)

func main() {
	LoadConfig("auth.toml")
	HandlersInit()

	router := apiserver.NewRouter(routes)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	server := &http.Server{
		Addr: Conf.Default.Bind,
		Handler: router,
		ErrorLog: logger,
	}

	log.Fatal(server.ListenAndServe())
}
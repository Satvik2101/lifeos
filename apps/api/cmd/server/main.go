package main

import (
	server "lifeos-api/internal/server"
	"log"
	"net/http"
)

func main(){
	addr:= ":8080"
	router := server.NewRouter()

	log.Printf("Starting server on port %s",addr)

	server := http.Server{
		Addr: addr,
		Handler: router,
	}
	
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
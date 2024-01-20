package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"tema-sprc-go/middleware"

	"tema-sprc-go/config"
	"time"

	"tema-sprc-go/router"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {

	if os.Getenv("ENV") == "development" {
		err := godotenv.Load(".env.development")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//initiate Ent Client
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	//set the client to the variable defined in package config
	//this will enable the client intance to be accessed anywhere through the accessor which is a function
	//named GetClient
	config.SetClient(client)

	//initiate router and register all the route
	r := mux.NewRouter()
	r.Use(middleware.Header)
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	
	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}

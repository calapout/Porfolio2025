package main

import (
	"Backend/Router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_, fileErr := os.Stat(".env")

	// LOAD IF FILE EXISTS, OTHERWISE DO NOTHING
	if fileErr == nil {
		loadErr := godotenv.Load()
		if loadErr != nil {
			log.Fatal("Error loading .env file: " + loadErr.Error())
		}
	}

	// CHECK ENVIRONMENT VARIABLES
	_, isBaseUrlSet := os.LookupEnv("BASE_URL")
	_, isApiTokenSet := os.LookupEnv("API_TOKEN")

	if !isBaseUrlSet {
		log.Fatal("BASE_URL environment variable is not set")
	}

	if !isApiTokenSet {
		log.Fatal("API_TOKEN environment variable is not set")
	}

	// SETUP ROUTER
	Router.Setup()

	// START SERVER
	port := ":8080"
	portValue, hasPortValue := os.LookupEnv("PORT")
	if hasPortValue {
		port = ":" + portValue
	}

	log.Println("Server started on port " + port)

	serverErr := http.ListenAndServe(port, nil)
	if serverErr != nil {
		log.Fatal("Error starting server: " + serverErr.Error())
	}
}

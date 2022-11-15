package main

import (
	"fmt"
	"go-microservice-example/pkg/api"
	"go-microservice-example/pkg/db"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("server has started")

	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := api.StartAPI(pgdb)
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}

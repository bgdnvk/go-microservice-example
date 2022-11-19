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
	//start the db
	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}
	//get the router of the API by passing the db
	router := api.StartAPI(pgdb)
	//get the port from the environment variable
	port := os.Getenv("PORT")
	//pass the router and start listening with the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}

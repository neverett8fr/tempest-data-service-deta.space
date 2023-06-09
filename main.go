package main

import (
	"log"

	"tempest-data-service-deta.space/cmd"
	application "tempest-data-service-deta.space/pkg/application/service"
	"tempest-data-service-deta.space/pkg/config"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes(conf config.Config) *mux.Router {
	r := mux.NewRouter()
	application.NewServiceRoutes(r, conf)

	return r
}

// Initiate web server
func main() {
	conf, err := config.Initialise()
	if err != nil {
		log.Fatalf("error initialising config, err %v", err)
		return
	}
	log.Println("config initialised")

	router := getRoutes(*conf)
	log.Println("API routes retrieved")

	err = cmd.StartServer(&conf.Service, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
		return
	}
	log.Println("server started")

}

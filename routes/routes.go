package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niketa/docker_orchestrator/controllers"
)

func CreateRoute() {

	r := mux.NewRouter()

	r.HandleFunc("/docker/config", controllers.DockerConfig).Methods("POST")
	r.HandleFunc("/docker/config/{id}", controllers.GetDockerConfig).Methods("GET")
	log.Fatal(http.ListenAndServe(":8585", r))
}

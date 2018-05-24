package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/niketa/docker_orchestrator/model"
	"github.com/niketa/docker_orchestrator/services"
	"gopkg.in/mgo.v2/bson"
)

func DockerConfig(w http.ResponseWriter, r *http.Request) {

	//defer r.Body.Close()
	var rootobject model.Root
	err := json.NewDecoder(r.Body).Decode(&rootobject)
	if err != nil {
		return
	}
	rootobject.ID = bson.NewObjectId()

	marshalData, _ := json.Marshal(rootobject)

	err = services.UnmarshalJsInsert(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalData)
}

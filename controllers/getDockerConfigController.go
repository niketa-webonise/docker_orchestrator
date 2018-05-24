package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niketa/docker_orchestrator/model"
	"github.com/niketa/docker_orchestrator/services"
	"gopkg.in/mgo.v2/bson"
)

func GetDockerConfig(w http.ResponseWriter, req *http.Request) {

	var rootobject model.Root
	vars := mux.Vars(req)
	rootobject.ID = bson.ObjectIdHex(vars["id"])
	//	marshljson := model.GetDockerItem(jsonobject)

	marshalData, _ := json.Marshal(rootobject)

	rootobject, err := services.UnmarshalJsGetItem(marshalData)
	if err != nil {
		log.Fatal(err)
		return
	}
	marshalResultData, _ := json.Marshal(rootobject)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", marshalResultData)
}

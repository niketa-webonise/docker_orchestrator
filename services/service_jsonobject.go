package services

import (
	"encoding/json"

	"github.com/niketa/docker_orchestrator/model"
)

func UnmarshalJsInsert(bytevalue []byte) error {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	return model.InsertJsonObject(rootobject)
}

func UnmarshalJsGetItem(bytevalue []byte) (model.Root, error) {

	var rootobject model.Root
	json.Unmarshal(bytevalue, &rootobject)
	rootobject, err := model.GetDockerItem(rootobject)
	return rootobject, err
}

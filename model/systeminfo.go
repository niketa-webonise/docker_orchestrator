package model

import ("github.com/niketa/gorepo/jsonToMongoDB/database"
		"log")


type JsonObject struct{
	SystemInfo SystemInfo `bson:"SystemInfo"`
}
type SystemInfo struct {
	GrorVersion string `bson:"grorVersion"`
	Name string `bson:"name"`
}



func InsertSystemInfo(jsonObject JsonObject) {

	
	c,err := database.ConnectDB()

	err = c.Insert(&SystemInfo{jsonObject.SystemInfo.GrorVersion,jsonObject.SystemInfo.Name})

	if err !=nil {
		log.Fatal(err)
	}

}


package main

import("fmt"
		"os"
        "io/ioutil"   
        "encoding/json"
        "github.com/niketa/gorepo/jsonToMongoDB/model"
        )


func main(){

	jsonPath := os.Getenv("JSON_FILE_NAME")

	if jsonPath == "" {
		
		fmt.Println("unable to identify path")
	}


	jsonFile,err := os.Open(jsonPath)

		if err!=nil {
			fmt.Println(err)
		}

	fmt.Println("Successfully opened sample_gror.json")
	defer jsonFile.Close()

	byteValue,_ := ioutil.ReadAll(jsonFile)

	var components model.Components
	var hosts model.Hosts
	var jsonObject model.JsonObject
	var authDatas model.AuthDatas


	 json.Unmarshal(byteValue,&components)
	 model.InsertComponent(components)

	json.Unmarshal(byteValue, &authDatas)
	model.InsertAuthData(authDatas)

	json.Unmarshal(byteValue,&jsonObject)
	model.InsertSystemInfo(jsonObject)
	
	json.Unmarshal(byteValue, &hosts)
	model.InsertHost(hosts)

}

	






package model

import ("github.com/niketa/gorepo/jsonToMongoDB/database"
		"log")

type Components struct {

	Components []Component `bson:"components"`
}

type Component struct{
	
	Instances []Instance  `bson:"instances"`
	Name string `bson:"name"`
}

type Instance struct {

	EnvMap EnvMap `bson:"EnvMap"`
	PortMapping string `bson:"portMapping"`
	AuthId string `bson:"authId"`
	HostId string `bson:"hostId"`
	VolumeMapping VolumeMapping `bson:"volumeMapping"`
	VolumesFrom []string `bson:"volumesFrom"`
	CommandToBeExecuted string `bson:"commandToBeExecuted"`
	Links string `bson:"links"`
	ImageName string `bson:"imageName"`
	Tag string `bson:"tag"`
	HostsMapping string `bson:"hostsMapping"`
	Name string `bson:"name"`

}

type EnvMap struct {
	
	CASSANDRA_BROADCAST_ADDRESS string `bson:"CASSANDRA_BROADCAST_ADDRESS"`
	CASSANDRA_SEEDS string `bson:"CASSANDRA_SEEDS"`

}


type VolumeMapping struct{

	 CassData string `bson:"/home/ubuntu/cass-data"`
	 CassConfig string `bson:"/home/ubuntu/cass-config"`

}

func InsertComponent(components Components){

 	c,err := database.ConnectDB()
	 
	
	for i:=0;i<len(components.Components);i++{

		err = c.Insert(&Component{components.Components[i].Instances,components.Components[i].Name})

		if err!=nil{
		log.Fatal(err)
		}
	
	}

}
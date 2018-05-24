package model

import (
	"github.com/niketa/docker_orchestrator/database"
	"gopkg.in/mgo.v2/bson"
)

type Instance struct {
	EnvMap              EnvMap        `json:"EnvMap" bson:"EnvMap"`
	PortMapping         PortMapping   `json:"PortMapping" bson:"portMapping"`
	AuthId              string        `json:"authId" bson:"authId"`
	HostId              string        `json:"hostId" bson:"hostId"`
	VolumeMapping       VolumeMapping `json:"volumeMapping" bson:"volumeMapping"`
	VolumesFrom         []string      `json:"volumesFrom" bson:"volumesFrom"`
	CommandToBeExecuted string        `json:"commandToBeExecuted" bson:"commandToBeExecuted"`
	Links               Links         `json:"links" bson:"links"`
	ImageName           string        `json:"imageName" bson:"imageName"`
	Tag                 string        `json:"tag" bson:"tag"`
	HostsMapping        HostsMapping  `json:"hostsMapping" bson:"hostsMapping"`
	Name                string        `json:"name" bson:"name"`
}

type HostsMapping struct {
}

type PortMapping struct {
}

type Links struct {
}

type EnvMap struct {
	CASSANDRA_BROADCAST_ADDRESS string `json:"CASSANDRA_BROADCAST_ADDRESS" bson:"CASSANDRA_BROADCAST_ADDRESS"`
	CASSANDRA_SEEDS             string `json:"CASSANDRA_SEEDS" bson:"CASSANDRA_SEEDS"`
}

type VolumeMapping struct {
	CassData   string `json:"/home/ubuntu/cass-data" bson:"/home/ubuntu/cass-data"`
	CassConfig string `json:"/home/ubuntu/cass-config" bson:"/home/ubuntu/cass-config"`
}

type Component struct {
	Instances []Instance `json:"instances" bson:"instances"`
	Name      string     `json:"name" bson:"name"`
}

type Host struct {
	Protocol                string `json:"protocol" bson:"protocol"`
	ApiVersion              string `json:"apiVersion" bson:"apiVersion"`
	HostType                string `json:"hostType" bson:"hostType"`
	DockerVersion           string `json:"dockerVersion" bson:"dockerVersion"`
	Alias                   string `json:"alias" bson:"alias"`
	CertPathForDockerDaemon string `json:"certPathForDockerDaemon" bson:"certPathForDockerDaemon"`
	IP                      string `json:"ip" bson:"ip"`
	DockerPort              string `json:"dockerPort" bson:"dockerPort"`
}

type AuthData struct {
	Password string `json:"password" bson:"password"`
	Key      string `json:"key" bson:"key"`
	Username string `json:"username" bson:"username"`
	Auth     string `json:"auth" bson:"auth"`
	Email    string `json:"email" bson:"email"`
}

type SystemInfo struct {
	GrorVersion string `json:"grorVersion" bson:"grorVersion"`
	Name        string `json:"name" bson:"name"`
}

type JsonObject struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	SystemInfo SystemInfo    `json:"systemInfo" bson:"systemInfo"`
	AuthDatas  []AuthData    `json:"authData" bson:"authData"`
	Hosts      []Host        `json:"hosts" bson:"hosts"`
	Components []Component   `json:"components" bson:"components"`
}

func InsertJsonObject(jsonObject JsonObject) error {
	c := database.Collection()
	return c.Insert(jsonObject)
}
package model

import ("github.com/niketa/gorepo/jsonToMongoDB/database"	
		"log")


type Hosts struct{

	Hosts []Host `bson:"hosts"`
}

type Host struct {

	Protocol string `bson:"protocol"`
	ApiVersion string `bson:"apiVersion"`
	HostType string `bson:"hostType"`
	DockerVersion string `bson:"dockerVersion"`
	Alias string `bson:"alias"`
	CertPathForDockerDaemon string `bson:"certPathForDockerDaemon"`
	IP string `bson:"ip"`
	DockerPort string `bson:"dockerPort"`
}

func InsertHost(hosts Hosts){

	

	 c,err := database.ConnectDB()

	for i:=0;i<len(hosts.Hosts);i++{

	
		err = c.Insert(&Host{hosts.Hosts[i].Protocol,hosts.Hosts[i].ApiVersion,hosts.Hosts[i].HostType,hosts.Hosts[i].DockerVersion,hosts.Hosts[i].Alias,hosts.Hosts[i].CertPathForDockerDaemon,hosts.Hosts[i].IP,hosts.Hosts[i].DockerPort})

		if err != nil {
                log.Fatal(err)
        }

       
	}

}
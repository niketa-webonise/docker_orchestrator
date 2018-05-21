package model

import ("github.com/niketa/gorepo/jsonToMongoDB/database"
		"log")


type AuthDatas struct{
	AuthDatas []AuthData `bson:"authData"`

}
type AuthData struct {  
	Password string `bson:"password"`
    Key string `bson:"key"`
    Username string `bson:"username"`
    Auth string `bson:"auth"`
    Email string  `bson:"email"`
}

func InsertAuthData(authDatas AuthDatas) {
   
   c,err := database.ConnectDB()

	for i:=0;i<len(authDatas.AuthDatas);i++ {

		err = c.Insert(&AuthData{authDatas.AuthDatas[i].Password,authDatas.AuthDatas[i].Key,authDatas.AuthDatas[i].Username,authDatas.AuthDatas[i].Auth,authDatas.AuthDatas[i].Email})

		if err!=nil{
			log.Fatal(err)
		}
	}

}
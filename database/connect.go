package database

import ("gopkg.in/mgo.v2")

func ConnectDB() (*mgo.Collection,error){

	
	Session,err := mgo.Dial("mongodb://127.0.0.1:27017/")
	 if err != nil {
            panic(err)
        }
        Session.SetMode(mgo.Monotonic, true)
        c := Session.DB("demoDB").C("samples")

       return c,err
}


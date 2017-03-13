package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
	"github.com/nu7hatch/gouuid"
)

type DbHandler struct {
	ConnectionString string
}

func (db *DbHandler) GetUserByUsername (username string) UserModel  {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("NewUsers").C("users")

	result := UserModel{}
	err = c.Find(bson.M{"credentials.username": username}).One(&result)

	fmt.Print("db out: ")
	fmt.Print(result)
	fmt.Println();
	if err != nil {
		log.Print(err)
	}
	if(result == UserModel{}){
		var all []UserModel
		err:= c.Find(nil).All(&all)
		fmt.Println(all)
		if err != nil {
			panic(err)
		}
	}
	return result;
}

func (db *DbHandler) GetUserById (id string) UserModel  {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("NewUsers").C("users")

	result := UserModel{}
	err = c.Find(bson.M{"credentials.id": id}).One(&result)

	fmt.Print("db out: ")
	fmt.Print(result)
	fmt.Println();
	if err != nil {
		log.Print(err)
	}
	if(result == UserModel{}){
		var all []UserModel
		err:= c.Find(nil).All(&all)
		fmt.Println(all)
		if err != nil {
			panic(err)
		}
	}

	return result;
}
func (db *DbHandler) InsertUserIntoUsers(user UserModel) bool  {
	session, err := mgo.Dial(db.ConnectionString)
	if err != nil {
		panic(err)
	}
	id,_:= uuid.NewV4();
	user.Id = id.String()
	defer session.Close()
	c := session.DB("NewUsers").C("users")
	err = c.Insert(&user)
	if(err != nil){
		log.Println(err)
		return  false;
	}
	return true;
}

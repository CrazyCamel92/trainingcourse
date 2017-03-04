package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string
	Password string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// db name and collection name, will create if the collection does not exists
	c := session.DB("test").C("users")
	// insert key value pointer
	err = c.Insert(&User{"Erez", "mypassword"})
	if err != nil {
		log.Fatal(err)
	}

	result := User{}
	// find in the collection using key value pair of the specific props inside the
	// collection is like a Table in SQL
	err = c.Find(bson.M{"username": "Erez"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("password:", result.Password)
}
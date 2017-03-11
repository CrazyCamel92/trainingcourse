package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func LoginEndpoint(w http.ResponseWriter,req *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin","*");
	var loginModel LoginModel
	_= json.NewDecoder(req.Body).Decode(&loginModel)
	fmt.Println(loginModel)
	result := GetUserByUsername(loginModel.Username)
	if(result != UserModel{}){
		fmt.Fprint(w,"Welcome back "+result.Name)
	}else{
		fmt.Fprint(w,"false")
	}
}

func RegisterEndpoint(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	var registerModel UserModel
	_ = json.NewDecoder(req.Body).Decode(&registerModel)
	fmt.Print("input model: ")
	fmt.Print(registerModel);
	fmt.Println()

	if (registerModel.Name != "" && registerModel.Credentials != LoginModel{}){

		result:= GetUserByUsername(registerModel.Credentials.Username)
		if (result == UserModel{}) {
			InsertUserIntoUsers(registerModel)
		}
		fmt.Fprint(w, "true")
	}else{
		fmt.Println("bad request");
		fmt.Fprint(w,"bad request");
}
}
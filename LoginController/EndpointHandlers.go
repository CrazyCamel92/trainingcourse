package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"github.com/nu7hatch/gouuid"
)

func LoginEndpoint(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	var userModel UserModel;

	cookie, err := req.Cookie("login-website")
	if (err != nil){
		log.Println("coockie extraction failed");
		log.Println(err);
	}
	log.Println(cookie)
	if(cookie != nil){
		var cookieModel CookieModel
		json.Unmarshal([]byte(cookie.Value),&cookieModel)
		log.Println("searching for id")
		log.Println(cookieModel.Id)
		userModel= GetUserById(cookieModel.Id)
	}else if(cookie == nil){
		var loginModel LoginModel
		_ = json.NewDecoder(req.Body).Decode(&loginModel)
		userModel := GetUserByUsername(loginModel.Username)
		cookieModel:= CookieModel{Id:userModel.Id}
		jString,_:= json.Marshal(&cookieModel)
		cookieStr := string(jString)
		http.SetCookie(w, &http.Cookie{
			Name: "login-website",
			Value: cookieStr,
		})
	}
	if(userModel != UserModel{}){
		fmt.Fprint(w,"Welcome back "+userModel.Name)
	}else{
		fmt.Fprint(w,"false")
	}
}

func RegisterEndpoint(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var registerModel UserModel
	_ = json.NewDecoder(req.Body).Decode(&registerModel)

	if (registerModel.Name != "" && registerModel.Credentials != LoginModel{}){

		result:= GetUserByUsername(registerModel.Credentials.Username)
		if (result == UserModel{}) {
			id,_:= uuid.NewV4()
			registerModel.Id = id.String()
			InsertUserIntoUsers(registerModel)
		}
		fmt.Fprint(w, "true")
	}else{
		fmt.Println("bad request")
		fmt.Fprint(w,"bad request")
}
}
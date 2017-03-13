package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"github.com/nu7hatch/gouuid"
	"unicode"
)

type EndpointsManager struct {
	dbHandler DbHandler
}

func NewEndpointsManager () *EndpointsManager {
	_dbHandler :=DbHandler{ConnectionString:"localhost"}
	var endpointsManager = new(EndpointsManager);
	endpointsManager.dbHandler = _dbHandler;
	return  endpointsManager
}

func (manager *EndpointsManager) LoginEndpoint(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	var userModel UserModel;

	cookie, err := req.Cookie("login-website")
	if (err != nil){
		log.Println("coockie extraction failed");
		log.Println(err);
	}
	log.Println(cookie)
	if(cookie != nil){
		cookieModel:= CookieModel{}
		json.Unmarshal([]byte(cookie.Value),&cookieModel)
		userModel= manager.dbHandler.GetUserById(cookieModel.Id)
	}else if(cookie == nil){
		var loginModel LoginModel
		_ = json.NewDecoder(req.Body).Decode(&loginModel)
		userModel := manager.dbHandler.GetUserByUsername(loginModel.Username)
		cookieModel:= CookieModel{Id:userModel.Id}
		///////////////////////////
		// need to convert to unicode before using marshal
		////////////////////////////////////////
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

func (manager *EndpointsManager) RegisterEndpoint(w http.ResponseWriter,req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var registerModel UserModel
	_ = json.NewDecoder(req.Body).Decode(&registerModel)

	if (registerModel.Name != "" && registerModel.Credentials != LoginModel{}){

		result:= manager.dbHandler.GetUserByUsername(registerModel.Credentials.Username)
		if (result == UserModel{}) {
			id,_:= uuid.NewV4()
			registerModel.Id = id.String()
			manager.dbHandler.InsertUserIntoUsers(registerModel)
		}
		fmt.Fprint(w, "true")
	}else{
		fmt.Println("bad request")
		fmt.Fprint(w,"bad request")
}
}
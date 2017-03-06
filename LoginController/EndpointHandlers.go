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
	if(loginModel.Username == "erez" && loginModel.Password =="mypassword"){
		fmt.Fprint(w,"true")
	}else{
		fmt.Fprint(w,"false")
	}
}
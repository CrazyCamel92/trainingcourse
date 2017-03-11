package main

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserModel struct {
	Credentials LoginModel `json:"credentials"`
	Name string `json:"name"`

}

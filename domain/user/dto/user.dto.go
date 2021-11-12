package dto

type User struct {
	UserName     string `json:"userName"`
	UserAccount  string `json:"userAccount"`
	UserEmail    string `json:"userEmail"`
	UserPassword string `json:"userPassword"`
}
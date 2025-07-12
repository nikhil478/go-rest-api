package models

type User struct {
	UserID string `json:"userID"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

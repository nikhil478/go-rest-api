package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nikhil478/go-rest-api/internal/common"
	"github.com/nikhil478/go-rest-api/internal/models"
)

var users []*models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		common.SendErrorResponse(w, "Error while decoding body")
	}
	users = append(users, user)
	common.SendPlainResponse(w, user)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	common.SendPlainResponse(w, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	for _, user := range users {
		if user.UserID == userID {
			common.SendPlainResponse(w, user)
			return
		}
	}
	common.SendErrorResponse(w, "Error while fetching user by ID")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	for _, user := range users {
		if user.UserID == userID {
			common.SendPlainResponse(w, user)
			return
		}
	}
	common.SendErrorResponse(w, "No user found")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	for i, user := range users {
		if user.UserID == userID {
			users = append(users[:i], users[i+1:]...)
			common.SendPlainResponse(w, "User deleted successfully")
			return
		}
	}
	common.SendErrorResponse(w, "User not found")
}

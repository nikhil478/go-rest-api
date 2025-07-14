package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nikhil478/go-rest-api/internal/common"
	"github.com/nikhil478/go-rest-api/internal/models"
)

var apps []*models.App

func CreateApp(w http.ResponseWriter, r *http.Request) {
	app := models.App{}
	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		common.SendErrorResponse(w, "Error while decoding body")
	}
	apps = append(apps, &app)
	common.SendResponse(w, app)
}

func GetAllApp(w http.ResponseWriter, r *http.Request) {
	common.SendResponse(w, apps)
}

func GetAppByName(w http.ResponseWriter, r *http.Request) {
	appID := r.URL.Query().Get("appID")
	for _, app := range apps {
		if app.AppID == appID {
			common.SendResponse(w, app)
			return
		}
	}
	common.SendErrorResponse(w, "Error while fetching app by ID")
}

func UpdateApp(w http.ResponseWriter, r *http.Request) {
	appID := r.URL.Query().Get("appID")
	for _, app := range apps {
		if app.AppID == appID {
			common.SendResponse(w, app)
			return
		}
	}
	common.SendErrorResponse(w, "No app found")
}

func DeleteApp(w http.ResponseWriter, r *http.Request) {
	appID := r.URL.Query().Get("appID")
	for i, app := range apps {
		if app.AppID == appID {
			apps = append(apps[:i], apps[i+1:]...)
			common.SendResponse(w, "User deleted successfully")
			return
		}
	}
	common.SendErrorResponse(w, "User not found")
}

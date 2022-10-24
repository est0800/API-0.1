package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"API-0.1/model"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

func MapID(id string) (int, string, model.User) {

	for index, item := range model.Users {
		if item.Id == id {
			return index, item.Id, item
		}
	}
	return 0, "Can't find the user", model.User{}
}

func UserPost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var save model.User
	json.Unmarshal(reqBody, &save)

	user := model.ReturnUser(save.Nickname, save.Password)
	user.MakeUser(w, r)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.Users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, item_id, item := MapID(id)

	if item_id == "Can't find the user" {
		json.NewEncoder(w).Encode(map[string]string{"error": item_id})
	} else {
		json.NewEncoder(w).Encode(map[string]model.User{"user": item})
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var update model.User

	index, item_id, _ := MapID(id)

	if item_id == "Can't find the user" {
		json.NewEncoder(w).Encode(map[string]string{"error": item_id})
	} else {
		model.Users = append(model.Users[:index], model.Users[index+1:]...)

		update.Id = xid.New().String()
		json.NewDecoder(r.Body).Decode(&update)

		model.Users = append(model.Users, update)
		json.NewEncoder(w).Encode(update)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	index, item_id, _ := MapID(id)
	if item_id == "Can't find the user" {
		json.NewEncoder(w).Encode(map[string]string{"error": item_id})
	} else {
		model.Users = append(model.Users[:index], model.Users[index+1:]...)
		json.NewEncoder(w).Encode(map[string]string{"state": "User deleted"})
	}
}

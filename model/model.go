package model

import (
	"encoding/json"
	"net/http"

	"github.com/rs/xid"
)

type User struct {
	Id       string
	Nickname string
	Password string
}

var Users []User

func ReturnUser(items ...string) *User {
	var data []string
	data = append(data, items...)

	return &User{
		Id:       xid.New().String(),
		Nickname: data[0],
		Password: data[1],
	}
}

type UserReqs interface {
	MakeUser(w http.ResponseWriter, r *http.Request) (string, User)
	//GetUser()
}

func (u User) MakeUser(w http.ResponseWriter, r *http.Request) {
	Users = append(Users, u)

	json.NewEncoder(w).Encode(map[string]User{"user": u})
}

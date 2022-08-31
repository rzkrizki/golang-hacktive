package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int
	Nama string
}

type UserService struct {
	db []*User
}

type UserIface interface {
	Register(u *User) string
	GetUser() []*User
	GetUserbyID(number int) *User
	RegisterUser(w http.ResponseWriter, r *http.Request)
	GetUserRegister(w http.ResponseWriter, r *http.Request)
	GetUserRegisterbyID(w http.ResponseWriter, r *http.Request)
}

func NewUserService(db []*User) UserIface {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() []*User {
	return u.db
}

func (u *UserService) GetUserbyID(number int) *User {
	for _, v := range u.db {
		if v.ID == number {
			return v
		}
	}
	return nil
}

func (u *UserService) GetUserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		resGet := u.GetUser()
		json.NewEncoder(w).Encode(resGet)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func (u *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		user := u.Register(&User{ID: id, Nama: r.FormValue("name")})
		fmt.Println(user)
		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func (u *UserService) GetUserRegisterbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		user := u.GetUserbyID(id)
		json.NewEncoder(w).Encode(user)
		return

	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

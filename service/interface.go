package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	RegisterUser(c *gin.Context)
	GetUserRegister(c *gin.Context)
	GetUserRegisterbyID(c *gin.Context)
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

func (u *UserService) GetUserRegister(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	if c.Request.Method == "GET" {
		resGet := u.GetUser()
		c.JSON(http.StatusOK, resGet)
		return
	}

	c.JSON(http.StatusBadRequest, "Invalid method")
}

func (u *UserService) RegisterUser(c *gin.Context) {

	if c.Request.Method == "POST" {
		id, _ := strconv.Atoi(c.Params.ByName("id"))
		user := u.Register(&User{ID: id, Nama: c.Params.ByName("name")})
		fmt.Println(user)
		c.JSON(http.StatusOK, user)
		return
	}

	c.JSON(http.StatusBadRequest, "Invalid method")
}

func (u *UserService) GetUserRegisterbyID(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	if c.Request.Method == "GET" {
		id, _ := strconv.Atoi(c.Params.ByName("id"))
		user := u.GetUserbyID(id)
		c.JSON(http.StatusOK, user)
		return

	}

	c.JSON(http.StatusBadRequest, "Invalid method")
}

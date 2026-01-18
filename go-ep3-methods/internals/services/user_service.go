package service

import (
	"fmt"

	"github.com/wrong2learn/go-ep3-methods/internals/models"
)

type UserService struct{}

func (s UserService) Create() { //Public method
	s.log("user created")
}

func (s UserService) log(msg string) { //Private method
	fmt.Println(msg)
}

func (s UserService) Create2(u *models.User) {
	u.SetName("Default")
}

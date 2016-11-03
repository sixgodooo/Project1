package Server

import (

)

type User interface {
	UserName() string
	UserId() string
}

type UserImpl struct {
	Name string
	Id string
}

func (user *UserImpl) UserName() string {
	return user.Name
}

func (user *UserImpl) UserId() string {
	return user.Id
}

func CreateUser(name string, id string) User {
	userImpl := new(UserImpl)
	userImpl.Name = name
	userImpl.Id = id
	return userImpl
}
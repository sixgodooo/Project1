package Server

import (

)

type User interface {
	UserName() string
	UserId() string
	SetUserName(string)
	SetUserPasswd(string)
	UserPasswd()string
}

type UserImpl struct {
	_name string
	_id string
	_passwd string
}

func (user *UserImpl) UserName() string {
	return user._name
}

func (user *UserImpl) UserId() string {
	return user._id
}

//这是一个测试借口
func (user *UserImpl) SetUserName(name string) {
	user._name = name
}

func (user *UserImpl) SetUserPasswd(passwd string) {
	user._passwd = passwd
}

func (user *UserImpl) UserPasswd() string {
	return user._passwd
}

func CreateUser(name string, id string, passwd string) User {
	userImpl := new(UserImpl)
	userImpl._name = name
	userImpl._id = id
	userImpl._passwd = passwd
	return userImpl
}
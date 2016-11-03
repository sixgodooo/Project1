package Server

import (
)

type UserManager interface{
	AddUser(name string, id string) (bool, error)
	FindUser(user User) (bool, error)
	Check(user User) bool
}

type UserManagerImpl struct {
	//TODO
}

func (m *UserManagerImpl) AddUser(name string, id string) (bool, error) {
	return true, nil//TODO
}

func (m *UserManagerImpl) FindUser(user User) (bool, error) {
	return true, nil//TODO
}

func (m *UserManagerImpl) Check(user User) bool{
	return true//TODO
}

func (m *UserManagerImpl) Init() {
	//TODO
}

func CreateUserManager() UserManager {
	userMgr := new(UserManagerImpl)
	userMgr.Init();
	return userMgr
}
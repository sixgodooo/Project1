package Server

import (
)

type UserManager interface{
	AddUser(user User) (bool, error)
	FindUser(name string) (User, bool)
	Check(user User) bool
}

type UserManagerImpl struct {
	//TODO
	_userMap map[string]User
}

func (m *UserManagerImpl) AddUser(user User) (bool, error) {
	m._userMap[user.UserName()] = user
	return true, nil//TODO
}

func (m *UserManagerImpl) FindUser(name string) (User, bool) {
	user, exist := m._userMap[name]
	return user, exist
}

func (m *UserManagerImpl) Check(user User) bool{
	_, exist := m._userMap[user.UserName()]
	if exist == false {
		return false
	} else {
		userInMem, _ := m._userMap[user.UserName()]
		if userInMem.UserPasswd() == user.UserPasswd() {
			return true
		} else {
			return false
		}
	}
}

func (m *UserManagerImpl) Init() {
	//TODO
	//从数据库或者文件加载用户数据
	m._userMap = make(map[string]User)
}

func CreateUserManager() UserManager {
	userMgr := new(UserManagerImpl)
	userMgr.Init();
	return userMgr
}
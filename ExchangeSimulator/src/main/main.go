package main

import (
	"Server"
	"fmt"
)

func main() {
	//exchange := Server.CreateExchange()
	//exchange.Init()
	user := Server.CreateUser("zz", "1")
	userMgr := Server.CreateUserManager()
	userMgr.AddUser(user)
	user1, found := userMgr.FindUser("1")
	if found == true {
		fmt.Println("found")
		user1.SetUserName("ZZZ")
		user2, found2 := userMgr.FindUser("1")
		fmt.Println(user2.UserName())
		if found2 == false {
		}
	} else {
		fmt.Println("not found")
	}
	fmt.Println("123")
	
}

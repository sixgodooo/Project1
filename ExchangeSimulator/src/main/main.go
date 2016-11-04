package main

import (
	"Server"
	"fmt"
)

func main() {
	exchange := Server.CreateExchange()
	exchange.Init()
	fmt.Println("123")
	
}

package Server

import (

)

func main() {
	exchange := CreateExchange()
	//TODO程序启动开始监听端口，接受请求，解析报文，如果是新增ordr，调用exchange的addorder，如果是cancelorder，调用exchange的cancelorder
	user := CreateUser("zz", "zz")
	order := CreateOrder(1, 1, 1, 1, user)
	exchange.AddOrder(order)
	exchange.CancelOrder(order)
}

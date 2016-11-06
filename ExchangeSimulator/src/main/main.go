package main

import (
	"Server"
	"fmt"
)

func printOrder(order Server.Order) {
		fmt.Println("OrderId:" + order.OrderId())
		fmt.Println("Price:")
		fmt.Println(order.Price())
		fmt.Println("Amount:")
		fmt.Println(order.Amount())
}
func main() {
	//exchange := Server.CreateExchange()
	//exchange.Init()
	
	/*
	fmt.Println("测试User和UserManager User的创建，UserManager对于User的增加，查找和校验")
	user := Server.CreateUser("zz", "1")
	userMgr := Server.CreateUserManager()
	userMgr.AddUser(user)
	user1, found := userMgr.FindUser("1")
	if found == true {
		fmt.Println(user1.UserName() + " " + user1.UserId())
	} else {
		fmt.Println("not found")
	}
	if (userMgr.Check(user1) == true) {
		fmt.Println("user:" + user1.UserName() + " check OK")
	} else {
		fmt.Println("user:" + user1.UserName() + " check Failed")
	}
	user2 := Server.CreateUser("yy", "2")
	if (userMgr.Check(user2) == true) {
		fmt.Println("user:" + user2.UserName() + " check OK")
	} else {
		fmt.Println("user:" + user2.UserName() + " check Failed")
	}
	*/
	
	fmt.Println("测试OrderBook，包括order的增删改查，还有遍历，还有排序是否正确")
	testUser := Server.CreateUser("zz", "testCount")
	orderBook := Server.CreateOrderBook(1)
	fmt.Println("Empty orderBook productId:")
	fmt.Println(orderBook.ProductId())
	fmt.Println("BestOfferPrice:")
	fmt.Println(orderBook.BestOfferPrice())
	fmt.Println("BestBidPrice:")
	fmt.Println(orderBook.BestBidPrice())
	bidOrder1 := Server.CreateOrder(10, Server.Bid, 100, 1, testUser)
	bidOrder2 := Server.CreateOrder(11, Server.Bid, 10, 1, testUser)
	
	offerOrder1 := Server.CreateOrder(12, Server.Offer, 15, 1, testUser)
	offerOrder2 := Server.CreateOrder(13, Server.Offer, 100, 1, testUser)
	
	orderBook.AddOrder(bidOrder1)
	orderBook.AddOrder(bidOrder2)
	orderBook.AddOrder(offerOrder1)
	orderBook.AddOrder(offerOrder2)
	
	fmt.Println("BestOfferPrice:")
	fmt.Println(orderBook.BestOfferPrice())
	fmt.Println("BestBidPrice:" )
	fmt.Println(orderBook.BestBidPrice())
	
	bidOrders := orderBook.BidOrders()
	offerOrders := orderBook.OfferOrders()
	
	fmt.Println("BidOrders:")
	for i := 0; i < len(bidOrders); i++ {
		printOrder(bidOrders[i])
	}
	fmt.Println("OfferOrders:")
	for j := 0; j < len(offerOrders); j++ {
		printOrder(offerOrders[j])
	}
	
	fmt.Println("BestBidOrder:")
	printOrder(orderBook.BestBidOrder())
	fmt.Println("BestOfferOrder:")
	printOrder(orderBook.BestOfferOrder())
}

//接下来要做的
//加日志
//测试
//加互斥锁
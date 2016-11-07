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

func printOrderBook(orderBook Server.OrderBook){
	bidOrders := orderBook.BidOrders()
	offerOrders := orderBook.OfferOrders()
	fmt.Println(" ===============================Bid Orders=======================================")
	for i := 0; i < len(bidOrders); i++ {
		printOrder(bidOrders[i])
	}
	fmt.Println("===============================Offer Orders=======================================")
	for j := 0; j < len(offerOrders); j++ {
		printOrder(offerOrders[j])
	}

}

func testUserAndUserManager() {
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
}

func testOrderAndOrderBook() {
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
	
	exist1, _ := orderBook.FindOrder(bidOrder1.OrderId())
	if (exist1 == true) {
		fmt.Println("Find exist Order OK")
	}
	exist2, _ := orderBook.FindOrder("xxx")
	if (exist2 == false) {
		fmt.Println("Find not exist Order OK")
	}
	
	fmt.Println("测试删除Order")
	orderBook.DelOrder(offerOrder1)
	newOfferOrders := orderBook.OfferOrders()
	for k := 0; k < len(newOfferOrders); k++ {
		printOrder(newOfferOrders[k])
	}
	fmt.Println("new BestOfferPrice")
	fmt.Println(orderBook.BestOfferPrice())
}

func testOrderBookManager() {
	orderBook1 := Server.CreateOrderBook(1)
	orderBook2 := Server.CreateOrderBook(2)
	
	orderBookMgr := Server.CreateOrderBookManager()
	orderBookMgr.AddOrderBook(orderBook1)
	orderBookMgr.AddOrderBook(orderBook2)
	
	exist,_:= orderBookMgr.FindOrderBook(1)
	if (exist == true) {
		fmt.Println("Find exist OK")
	}
	orderBookMgr.DelOrderBook(1)
	exist,_ = orderBookMgr.FindOrderBook(1)
	if (exist == false) {
		fmt.Println("Find not exist OK")
		fmt.Println("Delete OK")
	}
	
}

//为了测试这一部分，在OrderBookManager初始化时手动添加了OrderBook，正常需要从数据库加载，测试结束后需要将之删除
func testExecutionSystem() {
	//测试用例1 添加不满足交易条件的Bid和Offer，仅仅加入队列
	testUser := Server.CreateUser("zz", "testCount")
	bidOrder1 := Server.CreateOrder(10, Server.Bid, 100, 1, testUser)
	bidOrder2 := Server.CreateOrder(11, Server.Bid, 10, 1, testUser)
	offerOrder1 := Server.CreateOrder(15, Server.Offer, 15, 1, testUser)
	offerOrder2 := Server.CreateOrder(16, Server.Offer, 100, 1, testUser)
	execSystem := Server.CreateExecutionSystem(Server.OrderDriven)
	execSystem.AddOrder(bidOrder1)
	execSystem.AddOrder(bidOrder2)
	execSystem.AddOrder(offerOrder1)
	execSystem.AddOrder(offerOrder2)
	fmt.Println("After initialization")
	printOrderBook(execSystem.QueryOrderBook(1))
	
	//测试用例2 添加满足交易条件的Bid，全部满足
	bidOrder3 := Server.CreateOrder(15, Server.Bid, 5, 1, testUser)
	execSystem.AddOrder(bidOrder3)
	fmt.Println("After add a bid order at price 15 with amount 5")
	printOrderBook(execSystem.QueryOrderBook(1))

	//测试用例3 添加满足交易条件的Offer，全部满足
	offerOrder3 := Server.CreateOrder(11, Server.Offer, 5, 1, testUser)
	execSystem.AddOrder(offerOrder3)
	fmt.Println("After add a offer order at price 11 with amount 5")
	printOrderBook(execSystem.QueryOrderBook(1))

	//测试用例4 添加满足交易条件的Bid，只能满足一部分
	bidOrder4 := Server.CreateOrder(15, Server.Bid, 100, 1, testUser)
	execSystem.AddOrder(bidOrder4)
	fmt.Println("After add a bid order at price 15 with amount 100")
	printOrderBook(execSystem.QueryOrderBook(1))

	//测试用例5 添加满足交易条件的Offer，只能满足一部分
	offerOrder4 := Server.CreateOrder(10, Server.Offer, 100, 1, testUser)
	bidOrder5 := Server.CreateOrder(5, Server.Bid, 20, 1, testUser)
	execSystem.AddOrder(bidOrder5)
	execSystem.AddOrder(offerOrder4)
	fmt.Println("After add a bid order at price 5 with amount 20")
	fmt.Println("After add a offer order at price 10 with amount 100")
	printOrderBook(execSystem.QueryOrderBook(1))

	//测试用例6 取消没有交易的Bid和Offer，从队列中移除
	execSystem.CancelOrder(bidOrder5)
	fmt.Println("After cancel the bid order at price 5 with amount 20")
	printOrderBook(execSystem.QueryOrderBook(1))
	
	//测试用例7 取消已经交易的Bid和Offer，报错
	result, exception := execSystem.CancelOrder(bidOrder5)
	fmt.Println("After cancel a processed order")
	if result == false {
		fmt.Println(exception.Error())
	}
}

func main() {
	//exchange := Server.CreateExchange()
	//exchange.Init()
	//testUserAndUserManager()
	//testOrderAndOrderBook()
	//testOrderBookManager()
	testExecutionSystem()
	

}

//接下来要做的
//加日志
//测试
//加互斥锁
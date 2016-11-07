package Server

import (

)

type OrderBookManager interface {
	AddOrderBook(orderBook OrderBook)
	DelOrderBook(productId int)
	FindOrderBook(productId int) (bool,OrderBook)
	//Init()
}

type OrderBookManagerImpl struct {
	//TODO
	_orderBookMap map[int]OrderBook
}

func (o *OrderBookManagerImpl) AddOrderBook(orderBook OrderBook) {
	//TODO 新建一张数据库表将数据写入
	o._orderBookMap[orderBook.ProductId()] = orderBook
}

func (o *OrderBookManagerImpl) DelOrderBook(productId int) {
	//TODO 在数据库中将对应的表清空
	delete(o._orderBookMap, productId)
}

func (o *OrderBookManagerImpl) FindOrderBook(productId int) (bool, OrderBook){
	//若果没有，要做异常处理
	orderBook, exist := o._orderBookMap[productId]
	return exist, orderBook
}

func (o *OrderBookManagerImpl) Init() {
	o._orderBookMap = make(map[int]OrderBook)
	//TODO 从数据库加载各个orderbook
	//为了测试方便，需要先手工创建几个orderbook
	//测试代码，后续需要删除
	orderBook1 := CreateOrderBook(1)
	o._orderBookMap[1] = orderBook1
	orderBook2 := CreateOrderBook(2)
	o._orderBookMap[2] = orderBook2
	orderBook3 := CreateOrderBook(3)
	o._orderBookMap[3] = orderBook3
}

func CreateOrderBookManager() OrderBookManager {
	orderBookManager := new(OrderBookManagerImpl)
	orderBookManager.Init()
	return orderBookManager
}
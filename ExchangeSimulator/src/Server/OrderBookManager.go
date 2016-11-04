package Server

import (

)

type OrderBookManager interface {
	AddOrderBook(orderBook OrderBook)
	DelOrderBook(productId int)
	FindOrderBook(productId int) OrderBook
	Init()
}

type OrderBookManagerImpl struct {
	//TODO
	_orderBookMap map[int]OrderBook
}

func (o *OrderBookManagerImpl) AddOrderBook(orderBook OrderBook) {
	//TODO
}

func (o *OrderBookManagerImpl) DelOrderBook(productId int) {
	//TODO
}

func (o *OrderBookManagerImpl) FindOrderBook(productId int) OrderBook{
	//TODO
	return o._orderBookMap[productId]
}

func (o *OrderBookManagerImpl) Init() {
	//TODO
	o._orderBookMap = make(map[int]OrderBook)
}

func CreateOrderBookManager() OrderBookManager {
	orderBookManager := new(OrderBookManagerImpl)
	orderBookManager.Init()
	return orderBookManager
}
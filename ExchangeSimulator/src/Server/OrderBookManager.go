package Server

import (

)

type OrderBookManager interface {
	AddOrderBook(orderBook OrderBook)
	DelOrderBook(productId int)
	FindOrderBook(productId int)	
	Init()
}

type OrderBookManagerImpl struct {
	//TODO
}

func (o *OrderBookManagerImpl) AddOrderBook(orderBook OrderBook) {
	//TODO
}

func (o *OrderBookManagerImpl) DelOrderBook(productId int) {
	//TODO
}

func (o *OrderBookManagerImpl) FindOrderBook(productId int) {
	//TODO
}

func (o *OrderBookManagerImpl) Init() {
	//TODO
}

func CreateOrderBookManager() OrderBookManager {
	return new(OrderBookManagerImpl)
}
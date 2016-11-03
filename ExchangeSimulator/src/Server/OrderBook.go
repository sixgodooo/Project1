package Server

import (

)

type OrderBook interface {
	AddOrder(order Order)
	DelOrder(order Order)
	FindOrder(orderId int)
}

type OrderBookImpl struct {
	//TODO
}

func (o *OrderBookImpl) AddOrder(order Order) {
	//TODO
}

func (o *OrderBookImpl) DelOrder(order Order) {
	//TODO
}

func (o *OrderBookImpl) FindOrder(orderId int) {
	//TODO
}

func CreateOrderBook() OrderBook{
	return new(OrderBookImpl)
}
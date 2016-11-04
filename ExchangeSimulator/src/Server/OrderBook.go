package Server

import (

)

type OrderBook interface {
	AddOrder(order Order)
	DelOrder(order Order)
	FindOrder(orderId string) Order
	AllOrders() map[string]Order
	ProductId() int
	SetProductId(int)
	Init()
}

type OrderBookImpl struct {
	//TODO
	_productId int
	_orderMap map[string]Order
}

func (o *OrderBookImpl) AddOrder(order Order) {
	//TODO 将order数据存入数据库中的orderBook 数据库操作失败是要抛异常的
	o._orderMap[order.OrderId()] = order
}

func (o *OrderBookImpl) DelOrder(order Order) {
	//TODO 将order数据从数据库中的orderBook删除 数据库操作失败是要抛异常的
	delete(o._orderMap, order.OrderId())
}

func (o *OrderBookImpl) FindOrder(orderId string) Order {
	order, exist := o._orderMap[orderId]
	if (exist) {
		return order
	} else {
		//TODO不存在应该是要 抛异常的
		return order
	}
}

func (o *OrderBookImpl) AllOrders() map[string]Order {
	return o._orderMap
}

func (o *OrderBookImpl) ProductId() int {
	return o._productId
}

func (o *OrderBookImpl) SetProductId(productId int) {
	o._productId = productId
}

func (o *OrderBookImpl) Init() {
	//TODO 从数据库加载数据
}
func CreateOrderBook(productId int) OrderBook{
	orderBook := new(OrderBookImpl)
	orderBook.SetProductId(productId)
	orderBook.Init()
	return orderBook
}
package Server

import (

)

const (
	Bid = iota
	Offer
)

type Order interface {
    //CreateOrder(price int, orderType int, amount int, productId int) *Order
	Price() int
	//SetPrice(price uint)
	
	OrderType() int
	//SetOrderType(orderType int)
	
	Amount() uint
	//SetAmount(amount uint) 
	
	ProductId() int
	
	OrderId() int
	SetOrderId(int)
	
	User() User
}

type OrderImpl struct {
	_price int
	_orderType int
	_amount uint
	_productId int
	_orderId int
	_user User
}

func (s *OrderImpl) Price() int {
	return s._price
}

func (s *OrderImpl) OrderType() int {
	return s._orderType
}

func (s *OrderImpl) Amount() uint {
	return s._amount
}

func (s *OrderImpl) ProductId() int {
	return s._productId
}

func (s *OrderImpl) OrderId() int {
	return s._orderId
}

func (s *OrderImpl) SetOrderId(id int) {
	s._orderId = id
}

func (s *OrderImpl) User() User{
	return s._user
}
func CreateOrder(price int, orderType int, amount uint, productId int, user User) Order {
	order := new(OrderImpl)
	order._price = price
	order._orderType = orderType
	order._amount = amount
	order._productId = productId
	order._user = user
	return order
}
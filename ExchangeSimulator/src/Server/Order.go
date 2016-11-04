package Server

import (
	"time"
)

const (
	Bid = iota
	Offer
)

type Order interface {
	Price() int	
	OrderType() int	
	Amount() uint
	ProductId() int
	OrderId() string
	SetOrderId(string)
	User() User
}

type OrderImpl struct {
	_price int
	_orderType int
	_amount uint
	_productId int
	_orderId string
	_user User
	_time time.Time
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

func (s *OrderImpl) OrderId() string {
	return s._orderId
}

func (s *OrderImpl) SetOrderId(id string) {
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
	order._time = time.Now()
	order._orderId = time.Now().String()
	return order
}
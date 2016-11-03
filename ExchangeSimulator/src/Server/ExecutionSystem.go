package Server

import (

)

const (
	OrderDriven = iota
	Brokered
)

type ExecutionSystem interface {
	AddOrder(order Order) (bool, error)
	CancelOrder(order Order)(bool, error)
	QueryOrderBook(productId int) OrderBook
}

type OrderDrivenSystem struct {
	//TODO
	_orderBookMap map[int]OrderBook
}

func (s *OrderDrivenSystem) AddOrder(order Order) (bool, error){
	//TODO
	return true, nil
}

func (s *OrderDrivenSystem) CancelOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

func (s *OrderDrivenSystem) QueryOrderBook(productId int) OrderBook {
	return s._orderBookMap[productId]
}

func (s *OrderDrivenSystem) Init() {
	s._orderBookMap = make(map[int]OrderBook)
	//从数据库将Orderbook数据加载至内存
}

type BrokeredSystem struct {
	//TODO	
	_orderBookMap map[int]OrderBook
}

func (s *BrokeredSystem) AddOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

func (s *BrokeredSystem) CancelOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

func (s *BrokeredSystem) QueryOrderBook(productId int) OrderBook {
	return s._orderBookMap[productId]
}

func (s *BrokeredSystem) Init() {
	s._orderBookMap = make(map[int]OrderBook)
	//从数据库将Orderbook数据加载至内存
}

func CreateExecutionSystem(systemType int) ExecutionSystem {
	if systemType == OrderDriven {
		system := new(OrderDrivenSystem)
		system.Init()
		return system
	} else {
		system := new(BrokeredSystem)
		system.Init()
		return system
	}

}
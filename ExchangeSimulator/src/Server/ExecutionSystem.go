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
	_orderBookManager OrderBookManager
	_log Log
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
	return s._orderBookManager.FindOrderBook(productId)
}

func (s *OrderDrivenSystem) Init() {
	s._orderBookManager = CreateOrderBookManager()
	s._log = CreateLog("Server.txt")
}

type BrokeredSystem struct {
	//TODO	
	_orderBookManager OrderBookManager
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
	return s._orderBookManager.FindOrderBook(productId)
}

func (s *BrokeredSystem) Init() {
	s._orderBookManager = CreateOrderBookManager()
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
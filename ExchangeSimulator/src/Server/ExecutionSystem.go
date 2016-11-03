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
}

type OrderDrivenSystem struct {
	//TODO
}

func (s *OrderDrivenSystem) AddOrder(order Order) (bool, error){
	//TODO
	return true, nil
}

func (s *OrderDrivenSystem) CancelOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

type BrokeredSystem struct {
	//TODO
}

func (s *BrokeredSystem) AddOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

func (s *BrokeredSystem) CancelOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

func CreateExecutionSystem(systemType int) ExecutionSystem {
	if systemType == OrderDriven {
		return new(OrderDrivenSystem)
	} else {
		return new(BrokeredSystem)
	}

}
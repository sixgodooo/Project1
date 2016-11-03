package Server

import (

)

const (
	ContinousMarket = iota
	CallMarket
)

type TradingSession interface {
	Init() 
	
	AddOrder(order Order) (bool, error)
	CancelOrder(order Order) (bool, error)
	QueryOrderBook(productId int) OrderBook
	
	SessionId() string
	SessionType() int
	
	SetExecutionSystemType(int)
	ExecutionSystemType() int
	
	Start()
	Suspend()
	Resume()
	End()//期间要检查有没有过期的Order,干掉
}

type ContinousTradingSession struct {
	//TODO
	_executionSystemType int
	_executionSystem ExecutionSystem
}

func (s *ContinousTradingSession) Init() {
	//TODO
	s._executionSystem = CreateExecutionSystem(s._executionSystemType)
}

func (s *ContinousTradingSession) AddOrder(order Order) (bool, error) {
	//TODO
	return true, nil
}

func (s *ContinousTradingSession) CancelOrder(order Order)(bool, error) {
	//TODO
	return true, nil
} 

func (s *ContinousTradingSession) SessionId() string {
	return ""//TODO
}

func (s *ContinousTradingSession) SessionType() int {
	return 0;//TODO
}

func (s *ContinousTradingSession) SetExecutionSystemType(executionSystemType int) {
	s._executionSystemType = executionSystemType
} 

func (s *ContinousTradingSession) ExecutionSystemType() int {
	return s._executionSystemType
}

func (s *ContinousTradingSession) Start() {
	//TODO
}

func (s *ContinousTradingSession) Suspend() {
	//TODO
}

func (s *ContinousTradingSession) Resume() {
	//TODO
}

func (s *ContinousTradingSession) End() {
}

type CallTradingSession struct {
	_executionSystemType int
	_executionSystem ExecutionSystem
}

func (s *CallTradingSession) Init() {
	//TODO
	s._executionSystem = CreateExecutionSystem(s._executionSystemType)
}

func (s *CallTradingSession) AddOrder(order Order) (bool, error){
	//TODO
	return true, nil
}

func (s *CallTradingSession) CancelOrder(order Order) (bool, error){
	//TODO
	return true, nil
} 

func (s *CallTradingSession) SessionId() string {
	return ""//TODO
}

func (s *CallTradingSession) SessionType() int {
	return 0;//TODO
}

func (s *CallTradingSession) SetExecutionSystemType(executionSystemType int) {
	s._executionSystemType = executionSystemType
} 

func (s *CallTradingSession) ExecutionSystemType() int {
	return s._executionSystemType
}

func (s *CallTradingSession) Start() {
	//TODO
}

func (s *CallTradingSession) Suspend() {
	//TODO
}

func (s *CallTradingSession) Resume() {
	//TODO
}

func (s *CallTradingSession) End() {
	//TODO
}


func CreateTradingSession(sessionType int) TradingSession {
	if sessionType == ContinousMarket {
		return new(ContinousTradingSession)//TODO
	} else {
		return new(CallTradingSession)//TODO
	}
}

package Server

import (

)

const (
	ContinousMarket = iota
	CallMarket
)

type TradingSession interface {
	Init() 
	
	AddOrder(order Order)
	CancelOrder(order Order)
	
	SessionId() string
	SessionType() int
	
	Start()
	Suspend()
	Resume()
	End()//期间要检查有没有过期的Order,干掉
}

type ContinousTradingSession struct {
	//TODO
}

func (s *ContinousTradingSession) Init() {
	//TODO
}

func (s *ContinousTradingSession) AddOrder(order Order) {
	//TODO
}

func (s *ContinousTradingSession) CancelOrder(order Order) {
	//TODO
} 

func (s *ContinousTradingSession) SessionId() string {
	return ""//TODO
}

func (s *ContinousTradingSession) SessionType() int {
	return 0;//TODO
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
}

func (s *CallTradingSession) Init() {
	//TODO
}

func (s *CallTradingSession) AddOrder(order Order) {
	//TODO
}

func (s *CallTradingSession) CancelOrder(order Order) {
	//TODO
} 

func (s *CallTradingSession) SessionId() string {
	return ""//TODO
}

func (s *CallTradingSession) SessionType() int {
	return 0;//TODO
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

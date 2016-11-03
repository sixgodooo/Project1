package Server

import (

)
 
type Exchange struct {
	_userManager UserManager
	_tradingSessionManager TradingSessionManager 
	_envConfig EnvConfig
}

func (exchange *Exchange) AddOrder(order Order) (bool, error) {
	return true, nil//TODO
}

func (exchange *Exchange) CancelOrder(order Order) (bool, error) {
	return true, nil
}

func (exchange *Exchange) Init() {
	exchange._userManager = CreateUserManager()
	exchange._envConfig = CreateEnvConfig()
	exchange._tradingSessionManager = CreateTradingSessionManager()
	
	tradingSessionType := exchange._envConfig.TradingSessionType()
	executionSystemType := exchange._envConfig.ExecutionSystemType()
}
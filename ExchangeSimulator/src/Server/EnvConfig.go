package Server

import (

)

type EnvConfig interface {
	LoadGlobalSetting()
	TradingSessionType() int 
	ExecutionSystemType() int
}

type EnvConfigImpl struct {
	//TODO
	_tradingSessionType int
	_executionSystemType int
}

func (config *EnvConfigImpl) LoadGlobalSetting() {
	//TODO
	config._tradingSessionType = ContinousMarket
	config._executionSystemType = OrderDriven
}

func (config *EnvConfigImpl) TradingSessionType() int {
	//TODO
	return config._tradingSessionType
}

func (config *EnvConfigImpl) ExecutionSystemType() int {
	//TODO
	return config._executionSystemType
}

func CreateEnvConfig() EnvConfig {
	envConfig := new(EnvConfigImpl)
	envConfig.LoadGlobalSetting()
	return envConfig
}
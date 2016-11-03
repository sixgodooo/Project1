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
	//从数据库或者配置文件加载配置数据
	config._tradingSessionType = ContinousMarket
	config._executionSystemType = OrderDriven
}

func (config *EnvConfigImpl) TradingSessionType() int {
	//TODO 暂时只用一个session，所以只返回一个类型
	return config._tradingSessionType
}

func (config *EnvConfigImpl) ExecutionSystemType() int {
	//TODO 暂时只用一个session，所以只返回一个类型
	return config._executionSystemType
}

func CreateEnvConfig() EnvConfig {
	envConfig := new(EnvConfigImpl)
	envConfig.LoadGlobalSetting()
	return envConfig
}
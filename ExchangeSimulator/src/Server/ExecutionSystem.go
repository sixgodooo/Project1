package Server

import (

)

const (
	OrderDriven = iota
	Brokered
)

//存在一个选择，成功处理的order是否需要记录
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
	productId := order.ProductId()
	orderBook := s._orderBookManager.FindOrderBook(productId)
	if (order.OrderType() == Bid) {//如果是要买
		bestBidPrice := orderBook.BestBidPrice()
		bestOfferPrice := orderBook.BestOfferPrice()
		if (order.Price() <= bestBidPrice) {//先于当前最高出价比较，如果没有当前最高出价低，则加入队列，不需要其他处理
			orderBook.AddOrder(order)
			return true, nil
		} else {//如果比当前最高出价高，则与当前最低Offer比较，如果不到最低offer，则加入对立，不需要其他处理
			if (order.Price() < bestOfferPrice) {
				orderBook.AddOrder(order)
				return true, nil
			} else {
				//当前出价达到了最低Offer，这个order直接被处理掉，相应修改Offer的数据，甚至直接将offer的order删除，甚至可能一次向将几个order买光
				//TODO
			}
		}
	} else {//卖的情况
		bestBidPrice := orderBook.BestBidPrice()
		bestOfferPrice := orderBook.BestOfferPrice()
		if (order.Price() >= bestOfferPrice) {//如果没有超过当前的最低Offer则直接加入队列，不需要其他处理
			orderBook.AddOrder(order)
		} else {
			if (order.Price() > bestBidPrice) {//如果超出了当前最低价，但是还是高于最高报价，加入队列，不需要其他处理
				orderBook.AddOrder(order)
			} else {
				//当前出价低于了最高的bid，可能买了一个order的一部分，也可能买了一整个，也可能买了好几个
				//TODO
			}
		} 
	}
	return true, nil
}

func (s *OrderDrivenSystem) CancelOrder(order Order) (bool, error) {//如果这个Order在orderbook

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
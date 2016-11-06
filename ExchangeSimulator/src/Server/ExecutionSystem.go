package Server

import (
	"errors"
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
	_, orderBook := s._orderBookManager.FindOrderBook(productId)
	s._log.Log("BEGIN")
	s._log.LogInt("BestBidPrice:", orderBook.BestBidPrice())
	s._log.LogInt("BestOfferPrice:", orderBook.BestOfferPrice())
	s._log.LogInt("currentOrder Amount:", order.Amount())
	s._log.LogInt("currentOrder Price:", order.Price())
	s._log.Log("BEGIN")
	//目前假定一定能够找到OrderBOok
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
				//如果要买的数量没有买完，修改order之后加入队列
				//TODO
				//s._log.Log("1")
				for {
					if (order.Amount() == 0 || order.Price() < orderBook.BestOfferPrice()) {//如果数量买够了，或者买不起了，则不买了
						break
					}
					//s._log.Log("2")
					if (order.Amount() >= orderBook.BestOfferOrder().Amount()) {//当前的最低的Offer可以买光
						//s._log.Log("3")
						order.SetAmount(order.Amount() - orderBook.BestOfferOrder().Amount())
						//s._log.Log("order current status:")
						//s._log.LogInt("Amount:", order.Amount())
						//s._log.LogInt("Price:", order.Price())
						//s._log.LogInt("Type:", order.OrderType())
						orderBook.DelOrder(orderBook.BestOfferOrder())
					} else {//当前的最低offer买不光，说明这个order已经全部满足，不需要进入队列，另外修改OfferOrder的数量
						orderBook.BestOfferOrder().SetAmount(orderBook.BestOfferOrder().Amount() - order.Amount())
						order.SetAmount(0)
					}
				}
				//s._log.Log("4")
				//s._log.Log("order final status:")
				//s._log.LogInt("Amount:", order.Amount())
				//s._log.LogInt("Price:", order.Price())
				//s._log.LogInt("Type:", order.OrderType())
				if (order.Amount() > 0) {//如果还没有买光，就加入队列继续买
					//s._log.Log("5")
					orderBook.AddOrder(order)
				}
			}
		}
	} else {//卖的情况
					s._log.Log("begin")
					s._log.LogInt("order Amount:", order.Amount())
					s._log.LogInt("order Price:", order.Price())
					s._log.LogInt("bestBidOrder Amount:", orderBook.BestBidOrder().Amount())
					s._log.LogInt("bestBidOrder Price:",orderBook.BestBidOrder().Price())
					s._log.Log("begin")
		bestBidPrice := orderBook.BestBidPrice()
		bestOfferPrice := orderBook.BestOfferPrice()
		if (order.Price() >= bestOfferPrice) {//如果没有超过当前的最低Offer则直接加入队列，不需要其他处理
			orderBook.AddOrder(order)
			return true, nil
		} else {
			if (order.Price() > bestBidPrice) {//如果超出了当前最低价，但是还是高于最高报价，加入队列，不需要其他处理
				orderBook.AddOrder(order)
				return true, nil
			} else {
				//当前出价低于了最高的bid，可能买了一个order的一部分，也可能买了一整个，也可能买了好几个
				s._log.Log("1")
				for {
					if (order.Amount() == 0 || order.Price() > orderBook.BestBidPrice()) {//如果卖光了，或者卖不动了，就不卖了，
						break
					}
					s._log.Log("2")
					s._log.LogInt("order Amount:", order.Amount())
					s._log.LogInt("order Price:", order.Price())
					s._log.LogInt("bestBidOrder Amount:", orderBook.BestBidOrder().Amount())
					s._log.LogInt("bestBidOrder Price:",orderBook.BestBidPrice())
					if (order.Amount() >= orderBook.BestBidOrder().Amount()) {//能消化掉一个完整的bid order，再去消化下一个
						s._log.Log("3")
						order.SetAmount(order.Amount() - orderBook.BestBidOrder().Amount())
						orderBook.DelOrder(orderBook.BestBidOrder())
						s._log.LogInt("after del, size of bid seq:", len(orderBook.BidOrders()))
					} else {//消化不掉一个完整的bid order，就说明已经卖光了，另外修改BId order的量
						s._log.Log("3.5")
						orderBook.BestBidOrder().SetAmount(orderBook.BestBidOrder().Amount()- order.Amount())
						order.SetAmount(0)
					}
				}
				s._log.Log("4")
				if (order.Amount() > 0) {//如果还没有卖光，就加入队列继续卖
					s._log.Log("5")
					orderBook.AddOrder(order)
				}
			}
		} 
	}
	s._log.Log("END")
	s._log.LogInt("BestBidPrice:", orderBook.BestBidPrice())
	s._log.LogInt("BestOfferPrice:", orderBook.BestOfferPrice())
	s._log.LogInt("currentOrder Amount:", order.Amount())
	s._log.LogInt("currentOrder Price:", order.Price())
	s._log.Log("END")
	//TODO 怎么返回结果
	return true, nil
}

func (s *OrderDrivenSystem) CancelOrder(order Order) (bool, error) {//如果这个Order在orderbook中，就删掉，否则就不能删
	//假设OrderBook存在
	_,orderBook := s._orderBookManager.FindOrderBook(order.ProductId())
	exist, order := orderBook.FindOrder(order.OrderId())
	if (exist) {
		orderBook.DelOrder(order)
		return true, nil
	} else {
		return false, errors.New("Already processed")
	}

}

func (s *OrderDrivenSystem) QueryOrderBook(productId int) OrderBook {
	//TODO假设OrderBook存在
	_, orderBook := s._orderBookManager.FindOrderBook(productId)
	return orderBook
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
	//假设一定能找到
	_, orderBook := s._orderBookManager.FindOrderBook(productId)
	return orderBook
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
package Server

import (
	//"sort"
)

type OrderBook interface {
	AddOrder(order Order)
	DelOrder(order Order)
	FindOrder(orderId string) Order
	AllOrders() map[string]Order
	BidOrders() []Order
	OfferOrders() []Order
	ProductId() int
	SetProductId(int)
	Init()
}

type OrderBookImpl struct {
	//TODO
	_productId int
	_orderMap map[string]Order
	_bidOrderSeq []Order
	_offerOrderSeq []Order
}

func (o *OrderBookImpl) AddOrder(order Order) {
	//TODO 将order数据存入数据库中的orderBook 数据库操作失败是要抛异常的
	o._orderMap[order.OrderId()] = order
	if (order.OrderType() == Bid) {
		o._bidOrderSeq = append(o._bidOrderSeq, order)
	} else {
		o._offerOrderSeq = append(o._offerOrderSeq, order)
	}
}

func (o *OrderBookImpl) DelOrder(order Order) {
	//TODO 将order数据从数据库中的orderBook删除 数据库操作失败是要抛异常的
	_, exist := o._orderMap[order.OrderId()]
	if (exist == false) {
		return 
	}
	delete(o._orderMap, order.OrderId())
	//TODO 从序列中将Order删除
	if (order.OrderType() == Bid) {
		var index = 0
		for ; index < len(o._bidOrderSeq); index++ {
			if (o._bidOrderSeq[index].OrderId() == order.OrderId()) {
				break
			}
		}
		o._bidOrderSeq = append(o._bidOrderSeq[:index], o._bidOrderSeq[index:]...)
	} else {
		var index = 0
		for ; index < len(o._offerOrderSeq); index++ {
			if (o._offerOrderSeq[index].OrderId() == order.OrderId()) {
				break
			}
		}
		o._offerOrderSeq = append(o._offerOrderSeq[:index], o._offerOrderSeq[index:]...)
	}
}

func (o *OrderBookImpl) FindOrder(orderId string) Order {
	order, exist := o._orderMap[orderId]
	if (exist) {
		return order
	} else {
		//TODO不存在应该是要 抛异常的
		return order
	}
}

func (o *OrderBookImpl) AllOrders() map[string]Order {
	return o._orderMap
}

func (o *OrderBookImpl) BidOrders() []Order {
	return o._bidOrderSeq
}

func (o *OrderBookImpl) OfferOrders() []Order {
	return o._offerOrderSeq
}

func (o *OrderBookImpl) ProductId() int {
	return o._productId
}

func (o *OrderBookImpl) SetProductId(productId int) {
	o._productId = productId
}

func (o *OrderBookImpl) Init() {
	o._orderMap = make(map[string]Order)
	o._bidOrderSeq = make([]Order, 0)//TODO 正确性有待检验
	o._offerOrderSeq = make([]Order, 0)//TODO 正确性有待检验
	//TODO 从数据库加载数据
}

func CreateOrderBook(productId int) OrderBook{
	orderBook := new(OrderBookImpl)
	orderBook.SetProductId(productId)
	orderBook.Init()
	return orderBook
}

/*
func (o *OrderBookImpl) SortOrder(compare func (order1, order2 *Order) bool) {
	o._orderSeq = make([]Order, len(o._orderMap))
	i := 0
	for _, value := range o._orderMap {
		o._orderSeq[i] = value
		i++	
	}
	sort.Sort(OrderSorter{o._orderSeq, compare})
}
*/

type OrderSorter struct {
	_orders []Order
	_compare func(order1, order2 *Order) bool
}

func (orderSorter *OrderSorter) Len() int {
	return len(orderSorter._orders)
}

func (orderSorter *OrderSorter) Swap(i, j int) {
	orderSorter._orders[i], orderSorter._orders[j] = orderSorter._orders[j], orderSorter._orders[i]
}

func (orderSorter *OrderSorter) Less(i, j int) bool {
	return orderSorter._compare(&orderSorter._orders[i], &orderSorter._orders[j])
}

//OrderSorter排序的使用方法，sort.Sort( OrderSorter{orders, func (order1, order2 *Order) bool { 这里是对两个order的比较} } )
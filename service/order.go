package service

import "fmt"

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) GetOrderInfo(id int) {
	fmt.Printf("order: %d\n", id)
}

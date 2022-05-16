package service

import "fmt"

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) GetOrderInfo(uid int) {
	fmt.Printf("user:%d orders", uid)
}

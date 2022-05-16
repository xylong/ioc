package service

import "fmt"

type OrderService struct {
	Version string
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "1.0.0"}
}

func (s *OrderService) GetOrderInfo(uid int) {
	fmt.Printf("user:%d orders", uid)
}

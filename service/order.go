package service

import "fmt"

type OrderService struct {
	Version string
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "1.0"}
}

func (o *OrderService) GetOrderInfo(id int) {
	fmt.Printf("order: %d\n", id)
}

package service

import "fmt"

type IOrder interface {
	GetByUser(int2 int)
}

type OrderService struct {
	DB      *DBService `inject:"-"`
	Version string
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "1.0.0 "}
}

func (s *OrderService) GetByUser(uid int) {
	fmt.Printf("uid=%d的订单", uid)
}

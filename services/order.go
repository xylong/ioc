package services

import "fmt"

type OrderService struct {

}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (u *OrderService) GetOrderInfo(id int) {
	fmt.Println(id)
}
package service

import "fmt"

type UserService struct {
	order *OrderService
}

func NewUserService(order *OrderService) *UserService {
	return &UserService{order: order}
}

func (u *UserService) GetUserInfo(id int) {
	fmt.Printf("user: %d\n", id)
}

func (u *UserService) GetOrderInfo(id int) {
	u.order.GetOrderInfo(id)
}

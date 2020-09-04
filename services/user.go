package services

import "fmt"

type UserService struct {
	order *OrderService
}

func NewUserService(order *OrderService) *UserService {
	return &UserService{order: order}
}

func (u *UserService) GetUserInfo(id int) {
	fmt.Println(id)
}

func (u *UserService) GetOrderInfo(id int) {
	u.order.GetOrderInfo(id)
}
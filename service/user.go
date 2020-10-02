package service

import "fmt"

type UserService struct {
	Order *OrderService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserInfo(id int) {
	fmt.Printf("user: %d\n", id)
}

func (u *UserService) GetOrderInfo(id int) {
	u.Order.GetOrderInfo(id)
}

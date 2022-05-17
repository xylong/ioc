package service

import (
	"fmt"
	"strconv"
)

type UserService struct {
	Order *OrderService `inject`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUserInfo(uid int) {
	fmt.Println("userid is " + strconv.Itoa(uid))
}

func (s *UserService) GetOrderInfo(uid int) {
	s.Order.GetOrderInfo(uid)
}

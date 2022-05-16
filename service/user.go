package service

import (
	"fmt"
	"strconv"
)

type UserService struct {
	os *OrderService
}

func NewUserService(service *OrderService) *UserService {
	return &UserService{os: service}
}

func (s *UserService) GetUserInfo(uid int) {
	fmt.Println("userid is " + strconv.Itoa(uid))
}

func (s *UserService) GetOrderInfo(uid int) {
	s.os.GetOrderInfo(uid)
}

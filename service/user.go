package service

import "fmt"

type UserService struct {
	order *OrderService
}

func NewUserService(order *OrderService) *UserService {
	return &UserService{order: order}
}

func (s *UserService) Get(id int) {
	fmt.Println(id)
}

func (s *UserService) GetOrder(uid int) {
	s.order.GetByUser(uid)
}

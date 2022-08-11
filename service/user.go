package service

import "fmt"

type UserService struct {
	Order *OrderService `inject:"Service.Order()"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Get(id int) {
	fmt.Println(id)
}

func (s *UserService) GetOrder(uid int) {
	s.Order.GetByUser(uid)
}

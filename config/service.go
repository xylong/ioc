package config

import (
	"fmt"
	"github.com/xylong/ioc/service"
)

type Service struct {
}

func NewServiceConfig() *Service {
	return &Service{}
}

func (s *Service) Order() *service.OrderService {
	fmt.Println("a")
	return service.NewOrderService()
}

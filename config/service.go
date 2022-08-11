package config

import "github.com/xylong/ioc/service"

type Service struct {
}

func NewServiceConfig() *Service {
	return &Service{}
}

func (s *Service) Order() *service.OrderService {
	return service.NewOrderService()
}
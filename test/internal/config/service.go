package config

import (
	"github.com/xylong/ioc/test/internal/service"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Order() *service.OrderService {
	return service.NewOrderService()
}

func (s *Service) DB() *service.DBService {
	return service.NewDBService()
}

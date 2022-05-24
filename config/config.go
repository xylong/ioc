package config

import (
	"ioc/service"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) OrderService() *service.OrderService {
	return service.NewOrderService()
}

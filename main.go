package main

import (
	"fmt"
	"github.com/xylong/ioc/config"
	"github.com/xylong/ioc/injector"
	"github.com/xylong/ioc/service"
)

func main() {
	// 直接灌入
	//injector.Factory.Set(service.NewOrderService())
	//fmt.Println(injector.Factory.Get((*service.OrderService)(nil)))
	injector.Factory.Load(config.NewServiceConfig())
	//injector.Factory.Set(service.NewOrderService())

	{
		user := service.NewUserService()
		injector.Factory.Apply(user)
		fmt.Println(user.Order.DB)
	}

	{
		admin := service.NewAdminService()
		injector.Factory.Apply(admin)
		fmt.Println(admin.Order)
	}
}

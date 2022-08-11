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
	injector.Factory.Expr(config.NewServiceConfig())
	//injector.Factory.Set(service.NewOrderService())

	{
		user := service.NewUserService()
		injector.Factory.Apply(user)
		fmt.Println(user.Order)
	}

	{
		admin := service.NewAdminService()
		injector.Factory.Apply(admin)
		fmt.Println(admin.Order)
	}
}

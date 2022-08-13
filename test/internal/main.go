package main

import (
	"fmt"
	"github.com/xylong/ioc"
	"github.com/xylong/ioc/test/internal/config"
	service "github.com/xylong/ioc/test/internal/service"
)

func main() {
	// 直接灌入
	//ioc.Factory.Set(service.NewOrderService(), service.NewDBService())
	ioc.Factory.Unwrap(config.NewService())
	//fmt.Println(ioc.Factory.Get((*service.OrderService)(nil)))

	{
		user := service.NewUserService()
		ioc.Factory.Apply(user)
		fmt.Println(user.Order)
	}

	{
		admin := service.NewAdminService()
		ioc.Factory.Apply(admin)
		fmt.Println(admin.Order)
	}
}

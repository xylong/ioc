package main

import (
	"fmt"
	"github.com/xylong/ioc/injector"
	"github.com/xylong/ioc/service"
)

func main() {
	//user := service.NewUserService(service.NewOrderService())
	//user.Get(1)
	//user.GetOrder(1)

	injector.Factory.Set(service.NewOrderService())
	order := injector.Factory.Get((*service.OrderService)(nil))
	fmt.Println(order)
}

package main

import (
	"fmt"
	"github.com/xylong/ioc/injector"
	"github.com/xylong/ioc/service"
)

func main() {
	injector.Factory.Set(service.NewOrderService())
	order := injector.Factory.Get((*service.OrderService)(nil))
	fmt.Println(order)

	user := service.NewUserService()
	injector.Factory.Apply(user)

	fmt.Println(user.Order)
}

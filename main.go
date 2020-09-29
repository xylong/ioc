package main

import (
	"fmt"
	"ioc/injector"
	"ioc/service"
)

func main() {
	//userService := service.NewUserService(service.NewOrderService())
	//userService.GetUserInfo(1)
	//userService.GetOrderInfo(2)

	injector.Factory.Set(service.NewOrderService())
	order := injector.Factory.Get((*service.OrderService)(nil))
	fmt.Println(order)
}

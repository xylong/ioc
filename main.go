package main

import (
	"fmt"
	"ioc/injector"
	"ioc/service"
)

func main() {
	//uid := 123

	//userService := service.NewUserService(service.NewOrderService())
	//userService.GetUserInfo(uid)
	//userService.GetOrderInfo(uid)

	injector.BeanFactory.Set(service.NewOrderService())
	order := injector.BeanFactory.Get((*service.OrderService)(nil))
	fmt.Println(order)
}

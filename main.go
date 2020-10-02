package main

import (
	"fmt"
	"ioc/injector"
	"ioc/service"
)

func main() {
	orderService := service.NewOrderService()
	injector.Factory.Set(orderService)

	userService := service.NewUserService()
	injector.Factory.Apply(userService)

	fmt.Println(userService.Order)
}

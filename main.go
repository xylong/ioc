package main

import (
	"fmt"
	"ioc/injector"
	"ioc/service"
)

func main() {
	injector.Factory.Set(service.NewOrderService())

	userService := service.NewUserService()
	injector.Factory.Apply(userService)

	fmt.Println(userService.Order)
	fmt.Println(injector.Factory)
}

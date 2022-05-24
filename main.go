package main

import (
	"fmt"
	"ioc/config"
	"ioc/injector"
	"ioc/service"
)

func main() {
	uid := 123

	//userService := service.NewUserService(service.NewOrderService())
	//userService.GetUserInfo(uid)
	//userService.GetOrderInfo(uid)
	serverConfig := config.NewConfig()
	injector.BeanFactory.ExprMap = map[string]interface{}{
		"ServerConfig": serverConfig,
	}

	injector.BeanFactory.Set(serverConfig)

	userService := service.NewUserService()
	injector.BeanFactory.Apply(userService)
	//order := injector.BeanFactory.Get((*service.OrderService)(nil))
	fmt.Println(userService.Order)
	userService.Order.GetOrderInfo(uid)
}

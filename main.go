package main

import "ioc/service"

func main() {
	uid := 123

	userService := service.NewUserService(service.NewOrderService())
	userService.GetUserInfo(uid)
	userService.GetOrderInfo(uid)
}

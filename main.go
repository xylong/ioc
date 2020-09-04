package main

import "ioc/services"

func main() {
	id:=1

	userService:=services.NewUserService(services.NewOrderService())

	userService.GetUserInfo(id)
	userService.GetOrderInfo(id)
}

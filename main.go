package main

import "ioc/service"

func main() {
	user := service.NewUserService(service.NewOrderService())

	user.GetUserInfo(1)
	user.GetOrderInfo(2)
}

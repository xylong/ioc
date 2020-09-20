package main

<<<<<<< HEAD
import "ioc/service"

func main() {
	user := service.NewUserService(service.NewOrderService())

	user.GetUserInfo(1)
	user.GetOrderInfo(2)
=======
import "ioc/services"

func main() {
	id:=1

	userService:=services.NewUserService(services.NewOrderService())

	userService.GetUserInfo(id)
	userService.GetOrderInfo(id)
>>>>>>> d8e8ba4e9ccb19cb07595e1887d39ddb547ab403
}

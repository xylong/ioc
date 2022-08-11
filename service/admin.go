package service

type AdminService struct {
	Order *OrderService `inject:"Service.Order()"`
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

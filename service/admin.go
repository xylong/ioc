package service

type AdminService struct {
	Order *OrderService `inject:"-"`
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

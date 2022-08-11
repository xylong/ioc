package service

type AdminService struct {
	Order IOrder `inject:"-"`
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

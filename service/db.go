package service

type DBService struct {
	Name string
}

func NewDBService() *DBService {
	return &DBService{Name: "db"}
}

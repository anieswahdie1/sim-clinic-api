package service

import "sim-clinic-api/internal/model"

type AuthService interface {
	Register(request model.RegisterRequest) (*model.User, error)
	Login(request model.LoginRequest) (*model.LoginResponse, error)
	Logout(tokenString string, userID uint) error
	ValidateToken(tokenString string) (*model.User, error)
}

type UserService interface {
	GetAllUsers(currentUserRole string) ([]model.User, error)
	GetUserByID(id uint, currentUserRole string) (*model.User, error)
	UpdateUser(id uint, request model.UpdateUserRequest, currentUserRole string, currentUserID uint) (*model.User, error)
	DeleteUser(id uint, currentUserRole string, currentUserID uint) error
}

type MasterDataService interface {
	CreateLayananTerapi(request model.LayananTerapiRequest) (*model.LayananTerapi, error)
	GetAllLayananTerapi() ([]model.LayananTerapi, error)
	GetLayananTerapiByID(id uint) (*model.LayananTerapi, error)
	UpdateLayananTerapi(id uint, request model.LayananTerapiRequest) (*model.LayananTerapi, error)
	DeleteLayananTerapi(id uint) error

	CreateRiwayatPenyakit(request model.RiwayatPenyakitRequest) (*model.RiwayatPenyakit, error)
	GetAllRiwayatPenyakit() ([]model.RiwayatPenyakit, error)
	GetRiwayatPenyakitByID(id uint) (*model.RiwayatPenyakit, error)
	UpdateRiwayatPenyakit(id uint, request model.RiwayatPenyakitRequest) (*model.RiwayatPenyakit, error)
	DeleteRiwayatPenyakit(id uint) error

	CreateTeknikTerapi(request model.TeknikTerapiRequest) (*model.TeknikTerapi, error)
	GetAllTeknikTerapi() ([]model.TeknikTerapi, error)
	GetTeknikTerapiByID(id uint) (*model.TeknikTerapi, error)
	UpdateTeknikTerapi(id uint, request model.TeknikTerapiRequest) (*model.TeknikTerapi, error)
	DeleteTeknikTerapi(id uint) error
}

type CustomerService interface {
	CreateCustomer(request model.Customer) (*model.Customer, error)
	GetCustomer(requst model.RequestPagination) (*[]model.Customer, error)
	//UpdateCustomer(request *model.Customer) (*model.Customer, error)
}

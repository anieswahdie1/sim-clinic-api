package repository

import "sim-clinic-api/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindAll() ([]model.User, error)
	FindByRole(roleName string) ([]model.User, error)
	FindByRoles(roleNames []string) ([]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}

type RoleRepository interface {
	FindByID(id uint) (*model.Role, error)
}

type TokenRepository interface {
	BlacklistToken(token *model.BlacklistedToken) error
	IsTokenBlacklisted(token string) (bool, error)
	CleanExpiredTokens() error
	GetUserActiveTokens(userID uint) ([]model.BlacklistedToken, error)
}

type MasterDataRepository interface {
	// Layanan Terapi
	CreateLayananTerapi(layanan *model.LayananTerapi) error
	FindAllLayananTerapi() ([]model.LayananTerapi, error)
	FindLayananTerapiByID(id uint) (*model.LayananTerapi, error)
	FindLayananTerapiByCode(code string) ([]model.LayananTerapi, error)
	UpdateLayananTerapi(layanan *model.LayananTerapi) error
	DeleteLayananTerapi(id uint) error

	// Riwayat Penyakit
	CreateRiwayatPenyakit(riwayat *model.RiwayatPenyakit) error
	FindAllRiwayatPenyakit() ([]model.RiwayatPenyakit, error)
	FindRiwayatPenyakitByID(id uint) (*model.RiwayatPenyakit, error)
	FindRiwayatPenyakitByCode(code string) (*model.RiwayatPenyakit, error)
	UpdateRiwayatPenyakit(riwayat *model.RiwayatPenyakit) error
	DeleteRiwayatPenyakit(id uint) error

	// Teknik Terapi
	CreateTeknikTerapi(teknik *model.TeknikTerapi) error
	FindAllTeknikTerapi() ([]model.TeknikTerapi, error)
	FindTeknikTerapiByID(id uint) (*model.TeknikTerapi, error)
	FindTeknikTerapiByCode(code string) (*model.TeknikTerapi, error)
	UpdateTeknikTerapi(teknik *model.TeknikTerapi) error
	DeleteTeknikTerapi(id uint) error
}

type CustomerRepository interface {
	CreateCustomer(customer *model.Customer) error
	FindCustomerByID(id string) (*model.Customer, error)
	FindCustomerByPhoneNumber(phoneNumber string) (*model.Customer, error)
}

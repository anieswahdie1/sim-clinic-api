package repository

import (
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
)

type masterDataRepository struct {
	db *gorm.DB
}

func NewMasterDataRepository(db *gorm.DB) MasterDataRepository {
	return &masterDataRepository{db: db}
}

func (r *masterDataRepository) CreateLayananTerapi(layanan *model.LayananTerapi) error {
	return r.db.Create(layanan).Error
}

func (r *masterDataRepository) FindAllLayananTerapi() ([]model.LayananTerapi, error) {
	var layanans []model.LayananTerapi
	err := r.db.Find(&layanans).Error
	return layanans, err
}

func (r *masterDataRepository) FindLayananTerapiByID(id uint) (*model.LayananTerapi, error) {
	var layanan model.LayananTerapi
	err := r.db.First(&layanan, id).Error
	return &layanan, err
}

func (r *masterDataRepository) FindLayananTerapiByCode(code string) ([]model.LayananTerapi, error) {
	var layanan []model.LayananTerapi
	err := r.db.Where("code = ?", code).Find(&layanan).Error
	return layanan, err
}

func (r *masterDataRepository) UpdateLayananTerapi(layanan *model.LayananTerapi) error {
	return r.db.Save(layanan).Error
}

func (r *masterDataRepository) DeleteLayananTerapi(id uint) error {
	return r.db.Delete(&model.LayananTerapi{}, id).Error
}

func (r *masterDataRepository) CreateRiwayatPenyakit(riwayat *model.RiwayatPenyakit) error {
	return r.db.Create(riwayat).Error
}

func (r *masterDataRepository) FindAllRiwayatPenyakit() ([]model.RiwayatPenyakit, error) {
	var riwayats []model.RiwayatPenyakit
	err := r.db.Find(&riwayats).Error
	return riwayats, err
}

func (r *masterDataRepository) FindRiwayatPenyakitByID(id uint) (*model.RiwayatPenyakit, error) {
	var riwayat model.RiwayatPenyakit
	err := r.db.First(&riwayat, id).Error
	return &riwayat, err
}

func (r *masterDataRepository) FindRiwayatPenyakitByCode(code string) (*model.RiwayatPenyakit, error) {
	var riwayat []model.RiwayatPenyakit
	err := r.db.Where("code = ?", code).Find(&riwayat).Error
	if len(riwayat) == 0 {
		return nil, err
	}
	riwayats := model.RiwayatPenyakit{
		Code:        riwayat[0].Code,
		Name:        riwayat[0].Name,
		Description: riwayat[0].Description,
	}
	return &riwayats, err
}

func (r *masterDataRepository) UpdateRiwayatPenyakit(riwayat *model.RiwayatPenyakit) error {
	return r.db.Save(riwayat).Error
}

func (r *masterDataRepository) DeleteRiwayatPenyakit(id uint) error {
	return r.db.Delete(&model.RiwayatPenyakit{}, id).Error
}

// Teknik Terapi implementations (similar structure)
func (r *masterDataRepository) CreateTeknikTerapi(teknik *model.TeknikTerapi) error {
	return r.db.Create(teknik).Error
}

func (r *masterDataRepository) FindAllTeknikTerapi() ([]model.TeknikTerapi, error) {
	var teks []model.TeknikTerapi
	err := r.db.Find(&teks).Error
	return teks, err
}

func (r *masterDataRepository) FindTeknikTerapiByID(id uint) (*model.TeknikTerapi, error) {
	var teknik model.TeknikTerapi
	err := r.db.First(&teknik, id).Error
	return &teknik, err
}

func (r *masterDataRepository) FindTeknikTerapiByCode(code string) (*model.TeknikTerapi, error) {
	var teknik []model.TeknikTerapi
	err := r.db.Where("code = ?", code).Find(&teknik).Error
	if len(teknik) == 0 {
		return nil, err
	}

	teknikTerapi := model.TeknikTerapi{
		Code:        teknik[0].Code,
		Name:        teknik[0].Name,
		Description: teknik[0].Description,
	}
	return &teknikTerapi, err
}

func (r *masterDataRepository) UpdateTeknikTerapi(teknik *model.TeknikTerapi) error {
	return r.db.Save(teknik).Error
}

func (r *masterDataRepository) DeleteTeknikTerapi(id uint) error {
	return r.db.Delete(&model.TeknikTerapi{}, id).Error
}

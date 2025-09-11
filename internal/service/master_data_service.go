package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/repository"
)

type masterDataService struct {
	masterRepo repository.MasterDataRepository
}

func NewMasterDataService(masterRepo repository.MasterDataRepository) MasterDataService {
	return &masterDataService{masterRepo: masterRepo}
}

func (s *masterDataService) CreateLayananTerapi(request model.LayananTerapiRequest) (*model.LayananTerapi, error) {
	// Check if code already exists
	existing, _ := s.masterRepo.FindLayananTerapiByCode(request.Code)

	if len(existing) != 0 {
		return nil, &ServiceError{Message: "code already exists", Code: 400}
	}

	layanan := &model.LayananTerapi{
		Code: request.Code,
		Name: request.Name,
	}

	if err := s.masterRepo.CreateLayananTerapi(layanan); err != nil {
		return nil, err
	}

	logrus.Infof("Layanan terapi created: %s (%s)", request.Name, request.Code)
	return layanan, nil
}

func (s *masterDataService) GetAllLayananTerapi() ([]model.LayananTerapi, error) {
	return s.masterRepo.FindAllLayananTerapi()
}

func (s *masterDataService) GetLayananTerapiByID(id uint) (*model.LayananTerapi, error) {
	layanan, err := s.masterRepo.FindLayananTerapiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "layanan terapi not found", Code: 404}
		}
		return nil, err
	}
	return layanan, nil
}

func (s *masterDataService) UpdateLayananTerapi(id uint, request model.LayananTerapiRequest) (*model.LayananTerapi, error) {
	layanan, err := s.masterRepo.FindLayananTerapiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "layanan terapi not found", Code: 404}
		}
		return nil, err
	}

	// Check if new code conflicts with others
	if request.Code != layanan.Code {
		existing, _ := s.masterRepo.FindLayananTerapiByCode(request.Code)
		if existing != nil {
			return nil, &ServiceError{Message: "code already exists", Code: 400}
		}
	}

	layanan.Code = request.Code
	layanan.Name = request.Name

	if err := s.masterRepo.UpdateLayananTerapi(layanan); err != nil {
		return nil, err
	}

	logrus.Infof("Layanan terapi updated: %s (%s)", request.Name, request.Code)
	return layanan, nil
}

func (s *masterDataService) DeleteLayananTerapi(id uint) error {
	_, err := s.masterRepo.FindLayananTerapiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &ServiceError{Message: "layanan terapi not found", Code: 404}
		}
		return err
	}

	if err := s.masterRepo.DeleteLayananTerapi(id); err != nil {
		return err
	}

	logrus.Infof("Layanan terapi deleted: %d", id)
	return nil
}

func (s *masterDataService) CreateRiwayatPenyakit(request model.RiwayatPenyakitRequest) (*model.RiwayatPenyakit, error) {
	// Check if code already exists
	existing, _ := s.masterRepo.FindRiwayatPenyakitByCode(request.Code)
	if existing != nil {
		return nil, &ServiceError{Message: "code already exists", Code: 400}
	}

	riwayat := &model.RiwayatPenyakit{
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
	}

	if err := s.masterRepo.CreateRiwayatPenyakit(riwayat); err != nil {
		return nil, err
	}

	logrus.Infof("Riwayat penyakit created: %s (%s)", request.Name, request.Code)
	return riwayat, nil
}

func (s *masterDataService) GetAllRiwayatPenyakit() ([]model.RiwayatPenyakit, error) {
	return s.masterRepo.FindAllRiwayatPenyakit()
}

func (s *masterDataService) GetRiwayatPenyakitByID(id uint) (*model.RiwayatPenyakit, error) {
	riwayat, err := s.masterRepo.FindRiwayatPenyakitByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "riwayat penyakit not found", Code: 404}
		}
		return nil, err
	}
	return riwayat, nil
}

func (s *masterDataService) UpdateRiwayatPenyakit(id uint, request model.RiwayatPenyakitRequest) (*model.RiwayatPenyakit, error) {
	riwayat, err := s.masterRepo.FindRiwayatPenyakitByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "riwayat penyakit not found", Code: 404}
		}
		return nil, err
	}

	// Check if new code conflicts with others
	if request.Code != riwayat.Code {
		existing, _ := s.masterRepo.FindRiwayatPenyakitByCode(request.Code)
		if existing != nil {
			return nil, &ServiceError{Message: "code already exists", Code: 400}
		}
	}

	riwayat.Code = request.Code
	riwayat.Name = request.Name
	riwayat.Description = request.Description

	if err := s.masterRepo.UpdateRiwayatPenyakit(riwayat); err != nil {
		return nil, err
	}

	logrus.Infof("Riwayat penyakit updated: %s (%s)", request.Name, request.Code)
	return riwayat, nil
}

func (s *masterDataService) DeleteRiwayatPenyakit(id uint) error {
	_, err := s.masterRepo.FindRiwayatPenyakitByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &ServiceError{Message: "riwayat penyakit not found", Code: 404}
		}
		return err
	}

	if err := s.masterRepo.DeleteRiwayatPenyakit(id); err != nil {
		return err
	}

	logrus.Infof("Riwayat penyakit deleted: %d", id)
	return nil
}

func (s *masterDataService) CreateTeknikTerapi(request model.TeknikTerapiRequest) (*model.TeknikTerapi, error) {
	// Check if code already exists
	existing, _ := s.masterRepo.FindTeknikTerapiByCode(request.Code)
	if existing != nil {
		return nil, &ServiceError{Message: "code already exists", Code: 400}
	}

	teknik := &model.TeknikTerapi{
		Code:        request.Code,
		Name:        request.Name,
		Description: request.Description,
	}

	if err := s.masterRepo.CreateTeknikTerapi(teknik); err != nil {
		return nil, err
	}

	logrus.Infof("Teknik terapi created: %s (%s)", request.Name, request.Code)
	return teknik, nil
}

func (s *masterDataService) GetAllTeknikTerapi() ([]model.TeknikTerapi, error) {
	return s.masterRepo.FindAllTeknikTerapi()
}

func (s *masterDataService) GetTeknikTerapiByID(id uint) (*model.TeknikTerapi, error) {
	teknik, err := s.masterRepo.FindTeknikTerapiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "teknik terapi not found", Code: 404}
		}
		return nil, err
	}
	return teknik, nil
}

func (s *masterDataService) UpdateTeknikTerapi(id uint, request model.TeknikTerapiRequest) (*model.TeknikTerapi, error) {
	teknik, err := s.masterRepo.FindTeknikTerapiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "teknik terapi not found", Code: 404}
		}
		return nil, err
	}

	// Check if new code conflicts with others
	if request.Code != teknik.Code {
		existing, _ := s.masterRepo.FindTeknikTerapiByCode(request.Code)
		if existing != nil {
			return nil, &ServiceError{Message: "code already exists", Code: 400}
		}
	}

	teknik.Code = request.Code
	teknik.Name = request.Name
	teknik.Description = request.Description

	if err := s.masterRepo.UpdateTeknikTerapi(teknik); err != nil {
		return nil, err
	}

	logrus.Infof("Teknik terapi updated: %s (%s)", request.Name, request.Code)
	return teknik, nil
}

func (s *masterDataService) DeleteTeknikTerapi(id uint) error {
	_, err := s.masterRepo.FindTeknikTerapiByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &ServiceError{Message: "teknik terapi not found", Code: 404}
		}
		return err
	}

	if err := s.masterRepo.DeleteTeknikTerapi(id); err != nil {
		return err
	}

	logrus.Infof("Teknik terapi deleted: %d", id)
	return nil
}

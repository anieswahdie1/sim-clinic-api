package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
	"strconv"
)

var tagRepoUser = "internal.repository.user_repository."

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(page, limit, search string) ([]model.User, int64, error) {
	var (
		users               []model.User
		currPage, currLimit int
		total               int64
		tag                 = tagRepoUser + "FindAll."
	)

	query := r.db.Model(&model.User{}).Preload("Role")
	if page == "" {
		currPage = 1
	} else {
		currPage, _ = strconv.Atoi(page)
	}

	if limit == "" {
		currLimit = 5
	} else {
		currLimit, _ = strconv.Atoi(limit)
	}

	offset := (currPage - 1) * currLimit

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("username LIKE ?", searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(currLimit).Offset(offset).Find(&users).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err,
		})
		return nil, 0, err
	}
	return users, total, nil
}

func (r *userRepository) FindByRole(roleName string) ([]model.User, error) {
	var users []model.User
	err := r.db.Preload("Role").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("roles.name = ?", roleName).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByRoles(roleNames []string) ([]model.User, error) {
	var users []model.User
	err := r.db.Preload("Role").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("roles.name IN ?", roleNames).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

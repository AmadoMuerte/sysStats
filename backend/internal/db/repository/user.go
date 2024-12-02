package repository

import (
	"errors"
	"fmt"

	"github.com/AmadoMuerte/sysStats/internal/db"
	"github.com/AmadoMuerte/sysStats/internal/db/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) (uint, error)
	Update(user *models.User) error
	Delete(id int) error
}

type UserRepositoryImpl struct {
	storage *db.Storage
}

func NewUserRepository(storage *db.Storage) UserRepository {
	return &UserRepositoryImpl{storage: storage}
}

func (r *UserRepositoryImpl) GetUserByID(id int) (*models.User, error) {
	var user models.User
	if err := r.storage.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.storage.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *models.User) (uint, error) {
	var existingUser models.User
	if err := r.storage.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return 0, fmt.Errorf("user with email %s already exists", user.Email)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if err := r.storage.DB.Create(user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	return r.storage.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return r.storage.DB.Delete(&models.User{}, id).Error
}

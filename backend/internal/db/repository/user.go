package repository

import (
	"github.com/AmadoMuerte/FlickSynergy/internal/db"
	"github.com/AmadoMuerte/FlickSynergy/internal/db/models"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	Create(user *models.User) error
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

func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.storage.DB.Create(user).Error
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	return r.storage.DB.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return r.storage.DB.Delete(&models.User{}, id).Error
}

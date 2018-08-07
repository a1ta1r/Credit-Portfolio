package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/jinzhu/gorm"
)

type userStorage struct {
	DB gorm.DB
}

func (us userStorage) Create(user models.User) error {
	return us.DB.Create(&user).Error
}

func (us userStorage) Exists(user models.User) (bool, error) {
	result := us.DB.First(&user, user.ID)
	return result.RowsAffected != 0, result.Error
}

func (us userStorage) GetByID(id uint) (models.User, error) {
	var user models.User
	err := us.DB.First(&user, id).Error
	return user, err
}

func (us userStorage) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := us.DB.Where("Username = ?", username).First(&user).Error
	return user, err
}

func (us userStorage) Update(user models.User) error {
	return us.DB.Update(&user, user.ID).Error
}

func (us userStorage) Delete(user models.User) error {
	return us.DB.Delete(&user, user.ID).Error
}

func (us userStorage) GetAll(offset int64, limit int64) ([]models.User, error) {
	var users []models.User
	err := us.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

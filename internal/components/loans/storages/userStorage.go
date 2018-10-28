package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/jinzhu/gorm"
)

type UserStorage struct {
	DB gorm.DB
}

func (us UserStorage) Create(user entities.User) error {
	return us.DB.Create(&user).Error
}

func (us UserStorage) Exists(user entities.User) (bool, error) {
	result := us.DB.First(&user, user.ID)
	return result.RowsAffected != 0, result.Error
}

func (us UserStorage) GetByID(id uint) (entities.User, error) {
	var user entities.User
	err := us.DB.First(&user, id).Error
	return user, err
}

func (us UserStorage) GetByUsername(username string) (entities.User, error) {
	var user entities.User
	err := us.DB.Where("Username = ?", username).First(&user).Error
	return user, err
}

func (us UserStorage) Update(user entities.User) error {
	return us.DB.Save(&user).Error
}

func (us UserStorage) Delete(user entities.User) error {
	return us.DB.Delete(&user, user.ID).Error
}

func (us UserStorage) GetAll() ([]entities.User, error) {
	var users []entities.User
	err := us.DB.Find(&users).Error
	return users, err
}

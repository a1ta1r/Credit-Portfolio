package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/storages"
	"github.com/jinzhu/gorm"
)

func NewUserService(storageContainer storages.StorageContainer) UserService {
	return UserService{storageContainer: storageContainer}
}

type UserService struct {
	storageContainer storages.StorageContainer
}

func (us UserService) GetUsers() []models.User {
	if users, err := us.storageContainer.UserStorage.GetAll(); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		return users
	}
}

func (us UserService) CreateUser(user models.User) models.User {
	if err := us.storageContainer.UserStorage.Create(user); err != nil {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) UpdateUser(user models.User) models.User {
	if err := us.storageContainer.UserStorage.Update(user); err != nil {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) GetUserByID(id uint) models.User {
	if user, err := us.storageContainer.UserStorage.GetByID(id); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) GetUserByUsername(username string) models.User {
	if user, err := us.storageContainer.UserStorage.GetByUsername(username); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) DeleteUser(user models.User) {
	if err := us.storageContainer.UserStorage.Delete(user); err != nil {
		panic(err)
	}
}

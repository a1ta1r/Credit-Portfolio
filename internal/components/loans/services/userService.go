package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/common"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/jinzhu/gorm"
)

func NewUserService(storageContainer common.StorageContainer) UserService {
	return UserService{storageContainer: storageContainer}
}

type UserService struct {
	storageContainer common.StorageContainer
}

func (us UserService) GetUsers() []entities.User {
	if users, err := us.storageContainer.UserStorage.GetAll(); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		return users
	}
}

func (us UserService) CreateUser(user entities.User) (entities.User, bool) {
	users, _ := us.storageContainer.UserStorage.GetAll()

	exists := false
	for _, u := range users {
		if u.Username == user.Username || u.Email == user.Email {
			exists = true
		}
	}

	advertisers, _ := us.storageContainer.AdvertiserStorage.GetAdvertisers()
	for _, a := range advertisers {
		if a.Email == user.Email {
			exists = true
		}
	}


	if exists {
		return user, false
	}

	if err := us.storageContainer.UserStorage.Create(user); err != nil {
		panic(err)
	} else {
		return user, true
	}
}

func (us UserService) UpdateUser(user entities.User) entities.User {
	if err := us.storageContainer.UserStorage.Update(user); err != nil {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) GetUserByID(id uint) entities.User {
	if user, err := us.storageContainer.UserStorage.GetByID(id); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) GetUserByUsername(username string) entities.User {
	if user, err := us.storageContainer.UserStorage.GetByUsername(username); err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	} else {
		return user
	}
}

func (us UserService) DeleteUser(user entities.User) {
	if err := us.storageContainer.UserStorage.Delete(user); err != nil {
		panic(err)
	}
}

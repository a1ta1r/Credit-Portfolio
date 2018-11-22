package storages

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	userEntities "github.com/a1ta1r/Credit-Portfolio/internal/components/user/entities"
	"github.com/jinzhu/gorm"
	"time"
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

func (us UserStorage) GetCountByCreatedAt(from time.Time, to time.Time) (int, error) {
	var users []entities.User
	err := us.DB.Where("created_at > ? AND created_at < ?", from, to).Find(&users).Error
	return len(users), err
}

func (us UserStorage) GetCountsByCreatedAt(from time.Time, to time.Time) ([]userEntities.DayCount, error) {
	var dayCounts []userEntities.DayCount
	err := us.DB.Table("users").Select("date_trunc('day', created_at) as date, count(id) as count").Where("created_at is not null").Group("date_trunc('day', created_at)").Scan(&dayCounts).Error

	return dayCounts, err
}

func (us UserStorage) GetCountsByDeletedAt(from time.Time, to time.Time) ([]userEntities.DayCount, error) {
	var dayCounts []userEntities.DayCount
	err := us.DB.Table("users").Select("date_trunc('day', deleted_at) as date, count(id) as count").Where("deleted_at is not null").Group("date_trunc('day', deleted_at)").Scan(&dayCounts).Error

	return dayCounts, err
}

func (us UserStorage) GetCountsByLastSeen(from time.Time, to time.Time) ([]userEntities.DayCount, error) {
	var dayCounts []userEntities.DayCount
	err := us.DB.Table("users").Select("date_trunc('day', last_seen) as date, count(id) as count").Where("last_seen is not null").Group("date_trunc('day', last_seen)").Scan(&dayCounts).Error

	return dayCounts, err
}

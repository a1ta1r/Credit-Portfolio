package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/storages"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/user/entities"
	"time"
)

type UserStatisticsService struct {
	storages.UserStorage
}

func (uss UserStatisticsService) GetRegisteredUsersCount(from time.Time, to time.Time) (int, error) {
	return uss.GetCountByCreatedAt(from, to)
}

func (uss UserStatisticsService) GetRegisteredUsersDayCounts(from time.Time, to time.Time) ([]entities.DayCount, error) {
	return uss.GetCountsByCreatedAt(from, to)
}
func (uss UserStatisticsService) GetLastSeenUsersDayCounts(from time.Time, to time.Time) ([]entities.DayCount, error) {
	return uss.GetCountsByLastSeen(from, to)
}
func (uss UserStatisticsService) GetDeletedUsersDayCounts(from time.Time, to time.Time) ([]entities.DayCount, error) {
	return uss.GetCountsByDeletedAt(from, to)
}

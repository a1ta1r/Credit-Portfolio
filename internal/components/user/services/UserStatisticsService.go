package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/storages"
	"time"
)

type UserStatisticsService struct {
	storages.UserStorage
}

func (uss UserStatisticsService) GetRegisteredUsersCount(from time.Time, to time.Time) (int, error) {
	return uss.GetCountByCreatedAt(from, to)
}

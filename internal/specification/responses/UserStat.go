package responses

import "github.com/a1ta1r/Credit-Portfolio/internal/components/user/entities"

type UserStat struct {
	Status string `json:"status"`
	Count  []entities.DayCount    `json:"dayCounts"`
}

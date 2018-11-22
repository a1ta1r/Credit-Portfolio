package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/user/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserStatisticsController struct {
	userStat services.UserStatisticsService
}

func NewUserStatisticsController(userStatService services.UserStatisticsService) UserStatisticsController {
	return UserStatisticsController{userStatService}
}

// @Tags Statistics
// @Summary Получить количество зарегистрированных за данный период пользователей
// @Description Метод возвращает число новых пользователей за данный период
// @Param from query string false "Начало периода"
// @Param to query string false "Конец периода"
// @Produce  json
// @Success 200 {object} responses.UserStat
// @Failure 422
// @Router /stats/users/registered [get]
func (usc UserStatisticsController) GetRegisteredUsersCount(ctx *gin.Context) {
	from, to, dateError := usc.parseTime(ctx)
	if dateError != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadTimeFormat})
		return
	}
	users, err := usc.userStat.GetRegisteredUsersDayCounts(from, to)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"dayCounts": users,
	})
}

// @Tags Statistics
// @Summary Получить количество удаленных за данный период пользователей
// @Description Метод возвращает число удаленных пользователей за данный период
// @Param from query string false "Начало периода"
// @Param to query string false "Конец периода"
// @Produce  json
// @Success 200 {object} responses.UserStat
// @Failure 422
// @Router /stats/users/deleted [get]
func (usc UserStatisticsController) GetDeletedUsersCount(ctx *gin.Context) {
	from, to, dateError := usc.parseTime(ctx)
	if dateError != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadTimeFormat})
		return
	}
	users, err := usc.userStat.GetDeletedUsersDayCounts(from, to)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"dayCounts": users,
	})
}

// @Tags Statistics
// @Summary Получить количество активных за данный период пользователей
// @Description Метод возвращает число активных пользователей за данный период
// @Param from query string false "Начало периода"
// @Param to query string false "Конец периода"
// @Produce  json
// @Success 200 {object} responses.UserStat
// @Failure 422
// @Router /stats/users/active [get]
func (usc UserStatisticsController) GetLastSeenUsersCount(ctx *gin.Context) {
	from, to, dateError := usc.parseTime(ctx)
	if dateError != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadTimeFormat})
		return
	}
	users, err := usc.userStat.GetLastSeenUsersDayCounts(from, to)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"dayCounts": users,
	})
}


func (usc UserStatisticsController) parseTime(ctx *gin.Context) (time.Time, time.Time, error) {
	layout := time.RFC3339
	fromString := ctx.Query("from")
	toString := ctx.Query("to")
	var from time.Time
	var to time.Time
	var dateError error

	if toString != "" {
		to, dateError = time.Parse(layout, toString)
	} else {
		to = time.Now()
	}

	if fromString != "" {
		from, dateError = time.Parse(layout, fromString)
	} else {
		from = to.AddDate(0, 0, -7)
	}

	return from, to, dateError
}

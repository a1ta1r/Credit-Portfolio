package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"time"
)

func NewAgendaController(agendaService services.AgendaService) AgendaController {
	return AgendaController{agendaService: agendaService}
}

type AgendaController struct {
	agendaService services.AgendaService
}

func (ac AgendaController) GetAgendaElements(context *gin.Context) {
	layout := time.RFC3339
	fromString := context.Query("from")
	toString := context.Query("to")
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

	if dateError != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadTimeFormat})
		return
	}

	userId := uint(jwt.ExtractClaims(context)["user_id"].(float64))

	elements := ac.agendaService.GetElementsByTimeAndUserID(from, to, userId)

	context.JSON(http.StatusOK, gin.H{"dateFrom": from, "dateTo": to, "count": len(elements), "elements": elements})
}

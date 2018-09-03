package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
	"github.com/gin-gonic/gin"
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

	if fromString != "" {
		from, dateError = time.Parse(layout, fromString)
	} else {
		from = time.Now().AddDate(0, 0, -7)
	}

	if toString != "" {
		to, dateError = time.Parse(layout, toString)
	} else {
		time.Now()
	}

	if dateError != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadTimeFormat})
		return
	}

	elements := ac.agendaService.GetElementsByTime(from, to)

	context.JSON(http.StatusOK, gin.H{"count": len(elements), "elements": elements})
}

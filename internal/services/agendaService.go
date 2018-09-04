package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/jinzhu/gorm"
	"time"
)

func NewAgendaService(db gorm.DB) AgendaService {
	return AgendaService{db: db}
}

type AgendaService struct {
	db gorm.DB
}

func (as AgendaService) GetElementsByTimeAndUserID(from time.Time, to time.Time, userId uint) []models.AgendaElement {
	var elements []models.AgendaElement
	var incomes []models.Income
	var expenses []models.Expense
	var payments []models.Payment

	as.db.Where(&models.Income{UserID: userId}).Find(&incomes)
	as.db.Where(&models.Expense{UserID: userId}).Find(&expenses)
	//хз чё тут
	as.db.Where("user_id = ?", userId).Find(&payments)

	for _, income := range incomes {
		elements = append(elements, income.TransformWithPeriod(from, to)...)
	}

	for _, expense := range expenses {
		elements = append(elements, expense.TransformWithPeriod(from, to)...)
	}

	for _, payment := range payments {
		elements = append(elements, payment.TransformWithPeriod(from, to)...)
	}

	return elements
}
package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"github.com/jinzhu/gorm"
	"time"
)

func NewAgendaService(db gorm.DB) AgendaService {
	return AgendaService{db: db}
}

type AgendaService struct {
	db gorm.DB
}

func (as AgendaService) GetElementsByTimeAndUserID(from time.Time, to time.Time, userId uint) []entities.AgendaElement {
	var elements []entities.AgendaElement
	var incomes []entities.Income
	var expenses []entities.Expense
	var paymentPlans []entities.PaymentPlan
	var payments []entities.Payment

	as.db.Where(&entities.Income{UserID: userId}).Find(&incomes)
	as.db.Where(&entities.Expense{UserID: userId}).Find(&expenses)
	as.db.Where(&entities.PaymentPlan{UserID: userId}).Find(&paymentPlans)

	for _, paymentPlan := range paymentPlans {
		if paymentPlan.UserID == userId {
			as.db.Where(&entities.Payment{PaymentPlanID: paymentPlan.ID}).Find(&payments)
			for _, payment := range payments {
				elements = append(elements, payment.TransformWithPeriod(from, to)...)
			}
		}
	}

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

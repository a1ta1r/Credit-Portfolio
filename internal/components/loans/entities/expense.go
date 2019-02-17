package entities

import "time"

type Expense struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	User           User       `json:"-"`
	UserID         uint       `json:"userId"`
	Reason         string     `json:"reason"`
	Amount         float64    `json:"amount"`
	StartDate      time.Time  `json:"startDate"`
	IsRepeatable   bool       `json:"isRepeatable"` //рекуррентный платеж или нет
	Frequency      int        `json:"frequency"`
	PaymentPeriod  TimePeriod `json:"paymentPeriod"`
	RecurrentCount int        `json:"recurrentCount"` //число повторений в TimePeriod(4 недели, 12 лет)
}

func (e Expense) TransformSingle() AgendaElement {
	singleElement := AgendaElement{
		"Expense",
		e.ID,
		e.UserID,
		e.Reason,
		e.Amount,
		e.StartDate,
	}
	return singleElement
}

func (e Expense) TransformWithPeriod(from time.Time, to time.Time) []AgendaElement {
	if !e.IsRepeatable {
		return []AgendaElement{e.TransformSingle()}
	}
	var elements []AgendaElement
	currentIncomeTime := e.StartDate
	counter := 0
	for currentIncomeTime.Before(to) {
		if e.RecurrentCount - counter > 0 {
			if currentIncomeTime.After(from) {
				element := AgendaElement{
					"Expense",
					e.ID,
					e.UserID,
					e.Reason,
					e.Amount,
					currentIncomeTime,
				}
				elements = append(elements, element)
			}
				currentIncomeTime = e.getNextPeriod(currentIncomeTime)
				counter++
		} else {
			break
		}
	}
	return elements
}

func (e Expense) getNextPeriod(currentTime time.Time) time.Time {
	switch e.PaymentPeriod {
	case Day:
		return currentTime.AddDate(0, 0, e.Frequency)
	case Week:
		return currentTime.AddDate(0, 0, 7*e.Frequency)
	case Month:
		return currentTime.AddDate(0, e.Frequency, 0)
	case Quarter:
		return currentTime.AddDate(0, 3*e.Frequency, 0)
	case Year:
		return currentTime.AddDate(e.Frequency, 0, 0)
	default:
		return currentTime
	}
}

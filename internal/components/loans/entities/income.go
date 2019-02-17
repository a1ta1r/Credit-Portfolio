package entities

import "time"

type Income struct {
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

func (i Income) TransformSingle() AgendaElement {
	singleElement := AgendaElement{
		"Income",
		i.ID,
		i.UserID,
		i.Reason,
		i.Amount,
		i.StartDate,
	}
	return singleElement
}

func (i Income) TransformWithPeriod(from time.Time, to time.Time) []AgendaElement {
	if !i.IsRepeatable {
		return []AgendaElement{i.TransformSingle()}
	}
	var elements []AgendaElement
	currentIncomeTime := i.StartDate
	counter := 0
	for currentIncomeTime.Before(to) {
		if i.RecurrentCount - counter > 0 {
			if currentIncomeTime.After(from) {
				element := AgendaElement{
					"Income",
					i.ID,
					i.UserID,
					i.Reason,
					i.Amount,
					currentIncomeTime,
				}
				elements = append(elements, element)
			}
				currentIncomeTime = i.getNextPeriod(currentIncomeTime)
				counter++
		} else {
			break
		}
	}
	return elements
}

func (i Income) getNextPeriod(currentTime time.Time) time.Time {
	switch i.PaymentPeriod {
	case Day:
		return currentTime.AddDate(0, 0, i.Frequency)
	case Week:
		return currentTime.AddDate(0, 0, 7*i.Frequency)
	case Month:
		return currentTime.AddDate(0, i.Frequency, 0)
	case Quarter:
		return currentTime.AddDate(0, 3*i.Frequency, 0)
	case Year:
		return currentTime.AddDate(i.Frequency, 0, 0)
	default:
		return currentTime
	}
}

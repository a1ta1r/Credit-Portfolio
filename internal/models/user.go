package models

type User struct {
	name  string
	email string
	loans []PaymentPlan
}

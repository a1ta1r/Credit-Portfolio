package entities

const (
	Even           PaymentType = 0
	Differentiated PaymentType = 1
)

type PaymentType int

func (paymentType PaymentType) String() string {
	names := [...]string{
		"Even",
		"Differentiated"}
	if paymentType < Even || paymentType > Differentiated {
		return "Unknown"
	}
	return names[paymentType]
}

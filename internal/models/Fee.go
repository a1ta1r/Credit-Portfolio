package models

const (
	percent = "percent"
	total   = "total"
)

type Fee struct {
	fee     float64
	feeType string
}

func NewTotalFee(fee float64) Fee {
	return Fee{
		fee,
		total,
	}
}

func NewPercentFee(fee float64) Fee {
	return Fee{
		fee,
		percent,
	}
}

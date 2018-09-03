package models

const (
	Day     TimePeriod = 0
	Week    TimePeriod = 1
	Month   TimePeriod = 2
	Quarter TimePeriod = 3
	Year    TimePeriod = 4
)

type TimePeriod int

func (period TimePeriod) String() string {
	names := [...]string{
		"Day",
		"Week",
		"Month",
		"Quarter",
		"Year"}
	if period < Day || period > Year {
		return "Unknown"
	}
	return names[period]
}

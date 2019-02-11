package entities

const (
	Day     TimePeriod = 0
	Week    TimePeriod = 1
	Month   TimePeriod = 2
	Quarter TimePeriod = 3
	Year    TimePeriod = 4
	OneTime TimePeriod = 5
)

type TimePeriod int

func (period TimePeriod) String() string {
	names := [...]string{
		"Day",
		"Week",
		"Month",
		"Quarter",
		"Year",
		"OneTime",
	}
	if period < Day || period > OneTime {
		return "Unknown"
	}
	return names[period]
}

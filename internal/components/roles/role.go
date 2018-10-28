package roles

const (
	Basic Role = 0
	Admin Role = 1
	Ads   Role = 2
)

type Role int

func (role Role) String() string {
	names := [...]string{
		"User",
		"admin",
		"Advertiser"}
	if role < Basic || role > Ads {
		return "Unknown"
	}
	return names[role]
}

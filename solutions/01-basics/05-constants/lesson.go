//go:build ignore

package constants

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// DaysOfWeek returns a string representation of the day number.
// 0 = Sunday, 1 = Monday, ..., 6 = Saturday
func DaysOfWeek(day int) string {
	switch day {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return ""
	}
}

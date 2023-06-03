package enum

type BookingStatus string

const (
	Waiting  BookingStatus = "waiting"
	Accepted BookingStatus = "accepted"
	Rejected BookingStatus = "rejected"
)

func (e BookingStatus) String() string {
	switch e {
	case Waiting:
		return "waiting"
	case Accepted:
		return "accepted"
	case Rejected:
		return "rejected"
	default:
		return "unknown"
	}
}

package enum

type UnitStatus string

const (
	Idle             UnitStatus = "idle"
	InUse            UnitStatus = "in_use"
	Booked           UnitStatus = "booked"
	BookedWhileInUse UnitStatus = "booked_while_in_use"
)

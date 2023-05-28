package lib

import (
	"math"
	"time"
)

func CalculateTarif(start time.Time, end time.Time, hourlyPrice int) int {
	var secondsStep float64 = 300 // increment price every 5 minutes

	return int(math.Floor((end.Sub(start).Seconds()))/secondsStep) * hourlyPrice * int(secondsStep) / 3600
}

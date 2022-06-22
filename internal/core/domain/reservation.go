package domain

import "time"

type Reservation struct {
	Purchase        Purchase
	ReservationTime time.Time
	PickUpTime      time.Time
}

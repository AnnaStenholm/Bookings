package repository

import (
	"time"

	"github.com/AnnaStenholm/bookings/internal/models"
)

type DatabaseRepo interface {
	Allusers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error)
}

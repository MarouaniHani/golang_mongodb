package booking

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type BookingStatus string

const (
	StatusPending   BookingStatus = "PENDING"
	StatusConfirmed BookingStatus = "CONFIRMED"
	StatusCancelled BookingStatus = "CANCELLED"
)

var validStatuses = map[BookingStatus]struct{}{
	StatusPending:   {},
	StatusConfirmed: {},
	StatusCancelled: {},
}

func IsValidStatus(status BookingStatus) bool {
	_, ok := validStatuses[status]
	return ok
}

type Booking struct {
	ID           bson.ObjectID `json:"id" bson:"_id,omitempty"`
	CustomerName string        `json:"customer_name" bson:"customer_name"`
	HotelName    string        `json:"hotel_name" bson:"hotel_name"`
	Status       BookingStatus `json:"status" bson:"status"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

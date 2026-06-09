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

type Booking struct {
	ID           bson.ObjectID `json:"id" bson:"_id,omitempty"`
	CustomerName string        `json:"customer_name" bson:"customer_name"`
	HotelName    string        `json:"hotel_name" bson:"hotel_name"`
	Status       BookingStatus `json:"status" bson:"status"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

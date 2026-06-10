package booking

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ErrBookingNotFound = errors.New("booking not found")

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client) *Repository {
	db := os.Getenv("MONGO_DB_NAME")
	col := os.Getenv("MONGO_COLLECTION_NAME")
	return &Repository{
		collection: client.Database(db).Collection(col),
	}
}
func (r *Repository) CreateBooking(ctx context.Context, booking Booking) (string, error) {
	if booking.CreatedAt.IsZero() {
		booking.CreatedAt = time.Now()
	}
	if booking.Status == "" {
		booking.Status = StatusPending
	}
	res, err := r.collection.InsertOne(ctx, booking)
	id := res.InsertedID.(bson.ObjectID)
	return id.Hex(), err
}
func (r *Repository) GetBookingByID(ctx context.Context, id string) (*Booking, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking Booking
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&booking)
	if err != nil {
		return nil, err
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrBookingNotFound
	}
	return &booking, nil
}
func (r *Repository) GetAllBookings(ctx context.Context) ([]Booking, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var bookings []Booking
	if err = cursor.Decode(&bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *Repository) GetAllBookingsByStatus(ctx context.Context, status BookingStatus, page int64, limit int64) ([]Booking, error) {
	skip := (page - 1) * limit
	log.Printf(
		"Mongo query status=%q page=%d limit=%d",
		status,
		page,
		limit,
	)
	cursor, err := r.collection.Find(ctx, bson.M{"status": status},
		options.Find().
			SetSkip(skip).
			SetLimit(limit).
			SetSort(bson.D{{Key: "created_at", Value: -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	bookings := make([]Booking, 0)
	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

type UpdateBookingRequest struct {
	HotelName string        `json:"hotel_name"`
	Status    BookingStatus `json:"status"`
}

func (r *Repository) UpdateBooking(ctx context.Context, id string, updateReq UpdateBookingRequest) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID},
		bson.M{"$set": bson.M{"status": updateReq.Status, "hotel_name": updateReq.HotelName}})

	if result.ModifiedCount == 0 {
		return ErrBookingNotFound
	}
	return err
}
func (r *Repository) DeleteBooking(ctx context.Context, id string) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrBookingNotFound
	}
	return err
}

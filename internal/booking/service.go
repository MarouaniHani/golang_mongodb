package booking

import (
	"context"
	"errors"
)

var ErrHotelNameRequired = errors.New("hotel name is required")
var ErrStatusRequired = errors.New("status is required")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}
func (s *Service) CreateBooking(ctx context.Context, booking Booking) (string, error) {
	if booking.CustomerName == "" {
		return "", errors.New("customer name is required")
	}
	if booking.HotelName == "" {
		return "", errors.New("hotel name is required")
	}
	return s.repository.CreateBooking(ctx, booking)

}
func (s *Service) GetAllBookingsByStatus(ctx context.Context, status BookingStatus, page, limit int64) ([]Booking, error) {
	if page < 1 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 20
	}
	return s.repository.GetAllBookingsByStatus(ctx, status, page, limit)
}
func (s *Service) GetBookingByID(ctx context.Context, id string) (*Booking, error) {
	return s.repository.GetBookingByID(ctx, id)
}
func (s *Service) UpdateBooking(ctx context.Context, id string, updateReq UpdateBookingRequest) error {
	if updateReq.HotelName == "" {
		return ErrHotelNameRequired
	}
	if !IsValidStatus(updateReq.Status) {
		return ErrStatusRequired
	}
	return s.repository.UpdateBooking(ctx, id, updateReq)
}

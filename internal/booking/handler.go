package booking

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

var ErrInvalidBooking = errors.New("invalid booking data")

func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"CreateBooking called method=%s remote=%s",
		r.Method,
		r.RemoteAddr,
	)
	var booking Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if errors.Is(err, ErrInvalidBooking) {
		log.Printf("CreateBooking decode error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		log.Printf("CreateBooking internal error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	id, err := h.service.CreateBooking(r.Context(), booking)
	if err != nil {
		log.Printf("CreateBooking service error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
func (h *Handler) GetAllBookingsByStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"GetAllBookingsByStatus called method=%s remote=%s",
		r.Method,
		r.RemoteAddr,
	)
	status := BookingStatus(r.URL.Query().Get("status"))
	page := int64(1)
	limit := int64(20)
	if rawPage := r.URL.Query().Get("page"); rawPage != "" {
		parsedPage, err := strconv.ParseInt(rawPage, 10, 64)
		if err != nil {
			log.Printf("GetAllBookingsByStatus invalid page parameter: %v", err)
			http.Error(w, "invalid page parameter", http.StatusBadRequest)
			return
		}
		page = parsedPage
	}
	if rawLimit := r.URL.Query().Get("limit"); rawLimit != "" {
		parsedLimit, err := strconv.ParseInt(rawLimit, 10, 64)
		if err != nil {
			log.Printf("GetAllBookingsByStatus invalid limit parameter: %v", err)
			http.Error(w, "invalid limit parameter", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}
	bookings, err := h.service.GetAllBookingsByStatus(
		r.Context(),
		status,
		page,
		limit,
	)
	if err != nil {
		log.Printf("GetAllBookingsByStatus service error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)

}
func (h *Handler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"GetBookingByID called method=%s remote=%s",
		r.Method,
		r.RemoteAddr,
	)
	id := r.PathValue("id")
	booking, err := h.service.GetBookingByID(r.Context(), id)
	if errors.Is(err, ErrBookingNotFound) {
		log.Printf("GetBookingByID not found: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("GetBookingByID service error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(booking)
}

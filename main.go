package main

import (
	"context"
	"log"
	"net/http"
	"perfectstay-booking-api/internal/booking"
	"perfectstay-booking-api/internal/config"
)

func main() {
	ctx := context.Background()
	client, err := config.NewMongoClient(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("disconnect error: %v", err)
		}
	}()

	repo := booking.NewRepository(client)
	service := booking.NewService(repo)
	handler := booking.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /bookings", handler.CreateBooking)
	mux.HandleFunc("GET /bookings", handler.GetAllBookingsByStatus)
	mux.HandleFunc("GET /bookings/{id}", handler.GetBookingByID)
	mux.HandleFunc("PUT /bookings/{id}", handler.UpdateBooking)
	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

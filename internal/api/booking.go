package api

import (
	"encoding/json"
	"log"
	"net/http"
	"project/internal/app/booking"
	"project/internal/app/pkg/dto"
)

func CreateBooking(bookSvc booking.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var booking_details dto.BookingDetails

		err := json.NewDecoder(r.Body).Decode(&booking_details)

		if err != nil {
			log.Fatal("Error Occured during decoding", err)
			return
		}

		response, err := bookSvc.CreateUserBooking(booking_details)

		if err != nil {
			log.Fatal("Error Occured during decoding", err)
			return
		}

		json.NewEncoder(w).Encode(response)
	}
}

func GetSlots(bookSvc booking.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// details := dto.SlotResponse{}
		response, err := bookSvc.GetSlots(ctx)

		log.Println(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}

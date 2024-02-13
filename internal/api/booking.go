package api

import (
	"encoding/json"
	"fmt"
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
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding Update data from json into struct !!")
			return
		}

		err = booking_details.ValidateBooking()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}

		response, err := bookSvc.CreateUserBooking(booking_details)

		if err != nil {
			log.Fatal("Error Occured during decoding", err)
			return
		}
		if response.BookingID == 0 || response.CustomerName == "" {
			fmt.Fprint(w, "No Available slots")
		} else {
			fmt.Fprint(w, "Booking Done Successfully")
		}

		// json.NewEncoder(w).Encode(response)

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

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/internal/app/adminBookings"
	"project/internal/app/pkg/dto"
	"strconv"

	"github.com/gorilla/mux"
)

func AssignTableHandler(ab adminBookings.AdminBookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		assignReq := dto.AdminAssignTable{}

		err := json.NewDecoder(r.Body).Decode(&assignReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding Update data from json into struct !!")
			return
		}

		err = assignReq.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}
		_, err = ab.AdminAssignTable(assignReq)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusAccepted)

	}
}

func CancelTableHandler(ct adminBookings.AdminBookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["table_id"])
		req := dto.CancelTable{}
		if err != nil {
			fmt.Println("error occured in parsing int in CancelTableHandler " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resp, err := ct.AdminCancelTable(req, id)

		if err != nil {
			fmt.Println("Error Occured at AdminCancelTable", err.Error())
			return
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func UpdateTableHandler(ut adminBookings.AdminBookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateReq := dto.UpdateTable{}
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["table_id"])

		if err != nil {
			fmt.Println("error occured in parsing int in UpdateTableHandler " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewDecoder(r.Body).Decode(&updateReq)
		response, err := ut.AdminUpdateTable(updateReq, id)

		if err != nil {
			fmt.Println("Error in update table Handler: " + err.Error())
			w.WriteHeader(404)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func GetBookingsHandler(gbh adminBookings.AdminBookingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// details := dto.GetTable{}
		ctx := r.Context()
		response, err := gbh.AdminGetDetails(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}

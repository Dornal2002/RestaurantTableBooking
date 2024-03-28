package api

import (
	"net/http"
	"project/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(dep app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/getslots", GetSlots(dep.BookingService)).Methods(http.MethodGet)
	router.HandleFunc("/bookings", CreateBooking(dep.BookingService)).Methods(http.MethodPost)
	router.HandleFunc("/admin/create", SignUpHandler(dep.AdminService)).Methods(http.MethodPost)
	router.HandleFunc("/admin/login", LoginHandler(dep.AdminService)).Methods(http.MethodPost)
	router.HandleFunc("/admin/getuser", GetUsersHandler(dep.AdminService)).Methods(http.MethodGet)
	router.HandleFunc("/admin/assign_table", AssignTableHandler(dep.AdminBookingService)).Methods(http.MethodPut)
	router.HandleFunc("/admin/cancel_table/{booking_id}", CancelTableHandler(dep.AdminBookingService)).Methods(http.MethodDelete)
	router.HandleFunc("/admin/update_table/{booking_id}", UpdateTableHandler(dep.AdminBookingService)).Methods(http.MethodPut)
	router.HandleFunc("/admin/get_details", GetBookingsHandler(dep.AdminBookingService)).Methods(http.MethodGet)


	// router.Handle("/getslots",middleware.RequireAuth(GetSlots(dep.BookingService),[]string{"user,admin"}))
	return router

}

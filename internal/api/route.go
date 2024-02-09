package api

import (
	"net/http"
	"project/internal/app"

	"github.com/gorilla/mux"
)

func NewRouter(dep app.Dependencies) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/bookings", CreateBooking(dep.BookingService)).Methods(http.MethodPost)
	router.HandleFunc("/admin/create", CreateAdminAccount(dep.AdminService)).Methods(http.MethodPost)
	// router.HandleFunc("/admin/login", adminLogin).Methods(http.MethodPost)
	// // router.HandleFunc("/admin/bookings", getAdminBookings).Methods("GET")
	// // router.HandleFunc("/admin/tables", getTables).Methods("GET")
	// router.HandleFunc("/admin/assign_table/{table_id}", assignTable).Methods(http.MethodPut)
	// router.HandleFunc("/admin/cancel_table/{table_id}", cancelTable).Methods(http.MethodPut)
	// router.HandleFunc("/admin/update_table/{table_id}", updateTable).Methods(http.MethodPut)

	return router

}

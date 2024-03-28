package api

import (
	"net/http"
	"project/internal/app"
	"project/internal/app/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(dep app.Dependencies) *mux.Router {
	router := mux.NewRouter()


	// Get slots
	router.Handle("/getslots", middleware.RequireAuth(GetSlots(dep.BookingService), []string{"admin", "user"})).Methods(http.MethodGet)

	// Create booking
	router.Handle("/bookings", middleware.RequireAuth(CreateBooking(dep.BookingService), []string{"user"})).Methods(http.MethodPost)

	// User authentication
	router.HandleFunc("/signup", SignUpHandler(dep.AdminService)).Methods(http.MethodPost)
	router.HandleFunc("/login", LoginHandler(dep.AdminService)).Methods(http.MethodPost)

	// Admin endpoints
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Handle("/getusers", middleware.RequireAuth(GetUsersHandler(dep.AdminService), []string{"admin"})).Methods(http.MethodGet)
	adminRouter.Handle("/assign_table", middleware.RequireAuth(AssignTableHandler(dep.AdminBookingService), []string{"admin"})).Methods(http.MethodPut)
	adminRouter.Handle("/cancel_table/{booking_id}", middleware.RequireAuth(CancelTableHandler(dep.AdminBookingService), []string{"admin"})).Methods(http.MethodDelete)
	adminRouter.Handle("/update_table/{booking_id}", middleware.RequireAuth(UpdateTableHandler(dep.AdminBookingService), []string{"admin"})).Methods(http.MethodPut)
	adminRouter.Handle("/get_details", middleware.RequireAuth(GetBookingsHandler(dep.AdminBookingService), []string{"admin", "user"})).Methods(http.MethodGet)

	// router.HandleFunc("/getslots", GetSlots(dep.BookingService)).Methods(http.MethodGet)//
	// router.HandleFunc("/bookings", CreateBooking(dep.BookingService)).Methods(http.MethodPost)//
	// router.HandleFunc("/signup", SignUpHandler(dep.AdminService)).Methods(http.MethodPost)
	// router.HandleFunc("/login", LoginHandler(dep.AdminService)).Methods(http.MethodPost)
	// router.HandleFunc("/admin/getusers", GetUsersHandler(dep.AdminService)).Methods(http.MethodGet)//admin
	// router.HandleFunc("/admin/assign_table", AssignTableHandler(dep.AdminBookingService)).Methods(http.MethodPut)//admin
	// router.HandleFunc("/admin/cancel_table/{booking_id}", CancelTableHandler(dep.AdminBookingService)).Methods(http.MethodDelete)//admin
	// router.HandleFunc("/admin/update_table/{booking_id}", UpdateTableHandler(dep.AdminBookingService)).Methods(http.MethodPut)//admin
	// router.HandleFunc("/admin/get_details", GetBookingsHandler(dep.AdminBookingService)).Methods(http.MethodGet)//user admin 

	// write a route for the user so that he can get his booking details 
	// write a route to edit slots //admin


	// router.Handle("/getslots",middleware.RequireAuth(GetSlots(dep.BookingService),[]string{"user,admin"}))
	return router

}

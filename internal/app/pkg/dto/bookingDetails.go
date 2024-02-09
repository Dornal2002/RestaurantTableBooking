package dto

type BookingDetails struct {
	BookingID    int    `json:"booking_id"`
	CustomerName string `json:"customer_name"`
	ContactNo    string `json:"contact_no"`
	NoOfPeople   int    `json:"no_of_people"`
	SelectDate   string `json:"select_date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}

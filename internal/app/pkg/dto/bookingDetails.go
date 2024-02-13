package dto

type BookingDetails struct {
	BookingID    int    `json:"booking_id"`
	CustomerName string `json:"customer_name"`
	ContactNo    string `json:"contact_no"`
	Date         string `json:"date"`
	SlotId       int    `json:"slot_id"`
	TableId      int    `json:"table_id"`
}

type SlotResponse struct {
	SlotId    int    `json:"slot_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	TableId   []int  `json:"table_id"`
}

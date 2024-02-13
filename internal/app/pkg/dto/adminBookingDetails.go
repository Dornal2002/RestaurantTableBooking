package dto

type AdminAssignTable struct {
	BookingID    int    `json:"booking_id"`
	CustomerName string `json:"customer_name"`
	ContactNo    string `json:"contact_no"`
	Date         string `json:"date"`
	SlotId       int    `json:"slot_id"`
	TableID      int    `json:"table_id"`
}

type CancelTable struct {
	TableID int `json:"table_id"`
}

type UpdateTable struct {
	BookingID int `json:"booking_id"`
	TableID   int `json:"table_id"`
}

type GetTable struct {
	BookingID    int    `json:"booking_id"`
	CustomerName string `json:"customer_name"`
	ContactNo    string `json:"contact_no"`
	Date         string `json:"date"`
	SlotId       int    `json:"slot_id"`
	TableID      int    `json:"table_id"`
}

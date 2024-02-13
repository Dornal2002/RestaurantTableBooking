package dto

type TableDetails struct {
	TableID      int  `json:"table_id"`
	TableNo      int  `json:"table_no"`
	Availability bool `json:"availability"`
	TableSize    int  `json:"table_size"`
	BookingID    int  `json:"booking_id"`
}

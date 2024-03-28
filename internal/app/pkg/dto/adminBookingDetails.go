package dto

import (
	"fmt"
	"unicode"
)

type AdminAssignTable struct {
	BookingID    int    `json:"booking_id"`
	CustomerName string `json:"customer_name"`
	ContactNo    string `json:"contact_no"`
	Date         string `json:"date"`
	SlotId       int    `json:"slot_id"`
	TableID      int    `json:"table_id"`
}

type CancelTable struct {
	BookingID int `json:"booking_id"`
}

type UpdateTable struct {
	BookingID    int    `json:"booking_id"`
	CustomerName string `json:"customer_name"`
	ContactNo    string `json:"contact_no"`
	Date         string `json:"date"`
	SlotId       int    `json:"slot_id"`
	TableID      int    `json:"table_id"`
}



func (bd *AdminAssignTable) Validate() error {
	if len(bd.CustomerName) == 0 {
		return fmt.Errorf("name field cannot be empty")
	}
	if len(bd.ContactNo) == 0 {
		return fmt.Errorf("contact Number Cannot be empty")
	}
	if len(bd.ContactNo) != 10 {
		return fmt.Errorf("enter valid Mobile Number")
	}
	for _, char := range bd.ContactNo {
		if !unicode.IsDigit(char) {
			return fmt.Errorf("contact details must contain only digits")
		}
	}
	if len(bd.Date) == 0 {
		return fmt.Errorf("please mention date")
	}
	// if isValidDateFormat(bd.Date) {
	// 	return fmt.Errorf("enter valid format for date ie dd-mm-yyyy")
	// }
	// if len(req.Role) == 0 {
	// 	return fmt.Errorf("role field cannot be empty")
	// }
	// if req.Role != "Admin" && req.Role != "Customer" {
	// 	return fmt.Errorf("role field must be either 'Admin' or 'Customer'")
	// }
	return nil
}

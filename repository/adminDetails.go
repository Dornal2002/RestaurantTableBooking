package repository

import "project/internal/app/pkg/dto"

type AdminStorer interface {
	InsertAdminDetails(adminDetails dto.AdminDetails) (AdminPersonalDetails, error)
}

type AdminPersonalDetails struct {
	AdminID     int64
	Name        string
	ContactNo   string
	Email       string
	Password    string
	AccessToken string
}

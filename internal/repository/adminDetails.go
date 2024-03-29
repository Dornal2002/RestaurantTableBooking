package repository

import (
	"context"
	"project/internal/app/pkg/dto"
)

type AdminStorer interface {
	AdminSignup(ctx context.Context, user dto.AdminSignUpRequest) error
	AdminLogin(ctx context.Context, user dto.AdminLoginRequest) (int32, error)
	GetAdmin(ctx context.Context) ([]dto.AdminResponse, error)
}

type AdminPersonalDetails struct {
	AdminID   int64
	Name      string
	ContactNo string
	Email     string
	Password  string
}

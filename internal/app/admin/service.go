package admin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"project/internal/app/pkg/dto"
	"project/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	AdminSignup(ctx context.Context, user dto.AdminSignUpRequest) error
	AdminLogin(ctx context.Context, user dto.AdminLoginRequest) error
	GetAdmin(ctx context.Context) ([]dto.AdminResponse, error)
}

type service struct {
	AdminRepo repository.AdminStorer
}

func NewService(AdminRepo repository.AdminStorer) AdminService {
	return &service{
		AdminRepo: AdminRepo,
	}
}
func (as *service) AdminSignup(ctx context.Context, user dto.AdminSignUpRequest) error {
	valUser := validateUser(user)
	if !valUser {
		// fmt.Println("enter valid details")
		return errors.New("enter valid details")
	}
	user.Password = HashPassword(user.Password)
	err := as.AdminRepo.AdminSignup(ctx, user)
	if err != nil {
		fmt.Println("Error occured while admin signup", err.Error())
		return err
	}

	return nil

}

func (as *service) AdminLogin(ctx context.Context, user dto.AdminLoginRequest) error {

	valEmailPassword := isValidEmail(user.Email) && isValidPassword(user.Password)
	if !valEmailPassword {
		return errors.New("enter valid email and password")
	}
	// user.Password = HashPassword(user.Password)
	err := as.AdminRepo.AdminLogin(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
func (as *service) GetAdmin(ctx context.Context) ([]dto.AdminResponse, error) {

	userList, err := as.AdminRepo.GetAdmin(ctx)
	if err != nil {
		log.Println("error in GetAdmin service")
		return userList, err
	}
	return userList, nil
}
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

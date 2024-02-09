package admin

import (
	"errors"
	"fmt"
	"project/internal/app/pkg/dto"
	"project/repository"
)

type service struct {
	AdminRepo repository.AdminStorer
}

type AdminService interface {
	CreateAdminAccount(user dto.AdminDetails) (dto.AdminDetails, error)
}

func NewService(AdminRepo repository.AdminStorer) AdminService {
	return &service{
		AdminRepo: AdminRepo,
	}
}
func (as *service) CreateAdminAccount(user dto.AdminDetails) (dto.AdminDetails, error) {
	ad := dto.AdminDetails{}
	if user.Name == "" || user.ContactNo == "" || user.Email == "" || user.Password == "" || user.AccessToken == "" {
		return ad, errors.New("invalid user data")
	}
	fmt.Println(user)
	addDB, err := as.AdminRepo.InsertAdminDetails(user)

	ad = MapRepoObjectToDto(addDB) // converting db data into response
	if err != nil {
		fmt.Println(err.Error())
		return ad, err
	}

	return ad, nil

}

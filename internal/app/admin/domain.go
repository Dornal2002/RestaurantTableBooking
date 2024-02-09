package admin

import (
	"project/internal/app/pkg/dto"
	"project/repository"
)

func MapRepoObjectToDto(repoObj repository.AdminPersonalDetails) dto.AdminDetails {
	return dto.AdminDetails{
		AdminID:     repoObj.AdminID,
		Name:        repoObj.Name,
		ContactNo:   repoObj.ContactNo,
		Email:       repoObj.Email,
		Password:    repoObj.Password,
		AccessToken: repoObj.AccessToken,
	}
}

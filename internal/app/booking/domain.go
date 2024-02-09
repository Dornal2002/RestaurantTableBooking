package booking

import (
	"project/internal/app/pkg/dto"
	"project/repository"
)

func MapRepoObjectToDto(repoObj repository.BookingTableDetails) dto.BookingDetails {
	return dto.BookingDetails{
		BookingID:    int(repoObj.BookingID),
		CustomerName: repoObj.CustomerName,
		ContactNo:    repoObj.ContactNo,
		NoOfPeople:   repoObj.NoOfPeople,
		SelectDate:   repoObj.SelectDate,
		StartTime:    repoObj.StartTime,
		EndTime:      repoObj.EndTime,
	}
}

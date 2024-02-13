package booking

import (
	"project/internal/app/pkg/dto"
	"project/internal/repository"
)

func MapRepoObjectToDto(repoObj repository.BookingTableDetails) dto.BookingDetails {
	return dto.BookingDetails{
		BookingID:    int(repoObj.BookingID),
		CustomerName: repoObj.CustomerName,
		ContactNo:    repoObj.ContactNo,
		Date:         repoObj.Date,
		SlotId:       repoObj.SlotId,
		TableId:      repoObj.TableId,
	}
}

func MapRepoObjectToDto1(repoObj repository.SlotDetails) dto.SlotResponse {
	return dto.SlotResponse{
		SlotId:    repoObj.SlotId,
		StartTime: repoObj.StartTime,
		EndTime:   repoObj.EndTime,
	}
}

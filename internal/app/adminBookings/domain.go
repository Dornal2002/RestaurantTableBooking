package adminBookings

import (
	"project/internal/app/pkg/dto"
)

func MapRepoObjectToDto(repoObj dto.AdminAssignTable) dto.AdminAssignTable {
	return dto.AdminAssignTable{
		BookingID:    int(repoObj.BookingID),
		CustomerName: repoObj.CustomerName,
		ContactNo:    repoObj.ContactNo,
		Date:         repoObj.Date,
		SlotId:       repoObj.SlotId,
		TableID:      repoObj.TableID,
	}
}

func MapRepoObjectToDto1(repoObj dto.CancelTable) dto.CancelTable {
	return dto.CancelTable{
		BookingID: int(repoObj.BookingID),
	}
}

func MapRepoObjectToDto2(repoObj dto.UpdateTable) dto.UpdateTable {
	return dto.UpdateTable{
		BookingID: int(repoObj.BookingID),
	}
}

func MapRepoObjectToDto3(repoObj dto.BookingDetails) dto.BookingDetails {
	return dto.BookingDetails{
		BookingID:    int(repoObj.BookingID),
		// CustomerName: repoObj.CustomerName,
		// ContactNo:    repoObj.ContactNo,
		Id: repoObj.Id,
		Date:         repoObj.Date,
		SlotId:       repoObj.SlotId,
		TableId:      repoObj.TableId,
	}
}

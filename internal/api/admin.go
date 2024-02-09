package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/internal/app/admin"
	"project/internal/app/pkg/dto"
)

func CreateAdminAccount(AdminSrc admin.AdminService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var admin_details dto.AdminDetails

		err := json.NewDecoder(r.Body).Decode(&admin_details)

		if err != nil {
			log.Fatal("Error Occured during decoding", err)
			return
		}

		response, err := AdminSrc.CreateAdminAccount(admin_details)

		if err != nil {
			fmt.Println("Error in sending response", err)
			return
		}

		json.NewEncoder(w).Encode(response)

	}
}

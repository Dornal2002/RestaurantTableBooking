package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/internal/app/admin"
	"project/internal/app/pkg/dto"
)

func SignUpHandler(adminSvc admin.AdminService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var req dto.AdminSignUpRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error while decoding sign up data from json into struct !!")
			return
		}

		err = req.ValidateAdmin()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}

		err = adminSvc.AdminSignup(ctx, req)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Fprint(w, "Admin created successfully")
		w.WriteHeader(http.StatusAccepted)
	}
}

func LoginHandler(adminSvc admin.AdminService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		loginReq := dto.AdminLoginRequest{}

		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error in decoding")
			return
		}

		err = loginReq.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Plz, Provide Valid Credentials !!"))
			return
		}
		err = adminSvc.AdminLogin(ctx, loginReq)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Fprint(w, "login successful")
		w.WriteHeader(http.StatusAccepted)
	}
}

func GetUsersHandler(adminSvc admin.AdminService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := adminSvc.GetAdmin(ctx)
		log.Println(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}

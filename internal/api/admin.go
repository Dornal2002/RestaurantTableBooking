package api

import (
	"encoding/json"
	"log"
	"net/http"
	"project/internal/app/admin"
	"project/internal/app/pkg/dto"
)

func SignUpHandler(adminSvc admin.AdminService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sgnUpReq := dto.AdminSignUpRequest{}

		err := json.NewDecoder(r.Body).Decode(&sgnUpReq)
		if err != nil {
			log.Fatal(err)
			w.Write([]byte(err.Error()))
		}
		err = adminSvc.AdminSignup(ctx, sgnUpReq)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func LoginHandler(adminSvc admin.AdminService) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		loginReq := dto.AdminLoginRequest{}

		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			log.Fatal(err)
			w.Write([]byte(err.Error()))
		}
		err = adminSvc.AdminLogin(ctx, loginReq)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
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

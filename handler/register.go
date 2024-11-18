package handler

import (
	"database/sql"
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"go.uber.org/zap"
)

type AuthHandler struct {
	Service service.AllService
	Logger  *zap.Logger
	Config  util.Configuration
}

func NewAuthHandler(service service.AllService, logger *zap.Logger, config util.Configuration) AuthHandler {
	return AuthHandler{
		Service: service,
		Logger:  logger,
		Config:  config,
	}
}
func (rh *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	register := model.Register{
		Name:     r.FormValue("name"),
		Email:    sql.NullString{String: r.FormValue("email"), Valid: r.FormValue("email") != ""},
		Phone:    sql.NullString{String: r.FormValue("phone"), Valid: r.FormValue("phone") != ""},
		Password: r.FormValue("password"),
	}

	if !register.Email.Valid && !register.Phone.Valid {
		http.Error(w, "Either email or phone must be provided", http.StatusBadRequest)
		return
	}
	exists, err := rh.Service.AuthService.Repo.AuthRepo.CheckEmailOrPhoneExists(&register)
	if err != nil {
		rh.Logger.Error("Failed to check email or phone existence", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if exists {
		helper.SuccessResponse(w, 400, "Failed", "Email or phone already registered")
		return
	}

	err = rh.Service.AuthService.Register(&register)
	if err != nil {
		rh.Logger.Error("Failed to register user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Success", register)
}

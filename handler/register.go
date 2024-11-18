package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/helper"
	"project-app-ecommerce-ahmad-syarifuddin/model"
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"go.uber.org/zap"
)

type RegisterHandler struct {
	Service service.AllService
	Logger  *zap.Logger
	Config  util.Configuration
}

func NewRegisterHandler(service service.AllService, logger *zap.Logger, config util.Configuration) RegisterHandler {
	return RegisterHandler{
		Service: service,
		Logger:  logger,
		Config:  config,
	}
}
func (rh *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	register := model.Register{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Phone:    r.FormValue("phone"),
		Password: r.FormValue("password"),
	}

	if register.Email == "" && register.Phone == "" {
		helper.SuccessResponse(w, 400, "Failed", "Either email or phone must be provided")
		return
	}
	exists, err := rh.Service.RegisterService.Repo.RegisterRepo.CheckEmailOrPhoneExists(&register)
	if err != nil {
		rh.Logger.Error("Failed to check email or phone existence", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if exists {
		helper.SuccessResponse(w, 400, "Failed", "Email or phone already registered")
		return
	}

	err = rh.Service.RegisterService.Register(&register)
	if err != nil {
		rh.Logger.Error("Failed to register user", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "Success", register)
}

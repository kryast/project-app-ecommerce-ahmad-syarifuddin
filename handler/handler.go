package handler

import (
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	AuthHandler    AuthHandler
	ProductHandler ProductHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		AuthHandler:    NewAuthHandler(service, log, config),
		ProductHandler: NewProductHandler(service, log, config),
	}

}

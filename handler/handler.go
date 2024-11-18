package handler

import (
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	SampelHandler SampelHandler
	AuthHandler   AuthHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		SampelHandler: NewSampelService(service, log, config),
		AuthHandler:   NewAuthHandler(service, log, config),
	}

}

package handler

import (
	"net/http"
	"project-app-ecommerce-ahmad-syarifuddin/service"
	"project-app-ecommerce-ahmad-syarifuddin/util"

	"go.uber.org/zap"
)

type SampelHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewSampelService(service service.AllService, log *zap.Logger, config util.Configuration) SampelHandler {
	return SampelHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (SampelHandler *SampelHandler) Create(w http.ResponseWriter, r *http.Request) {
	// logic handler
}

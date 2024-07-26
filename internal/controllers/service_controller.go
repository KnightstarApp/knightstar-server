package controllers

import (
	"knightstar/internal/database"
	"knightstar/pkg/util"

	"net/http"
)

type ServiceController struct {
	*BaseController
}

func NewServiceController(baseController *BaseController) *ServiceController {
	return &ServiceController{baseController}
}

func (sc *ServiceController) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteJSONResponse(w, http.StatusOK, util.JSON{"message": "Hello World"})
}

func (sc *ServiceController) HealthHandler(w http.ResponseWriter, r *http.Request) {
	httpStatusCode, json := database.Health()
	util.WriteJSONResponse(w, httpStatusCode, json)
}

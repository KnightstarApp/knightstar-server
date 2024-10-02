package controllers

import (
	"knightstar/internal/database"
	"knightstar/pkg/util"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BaseController struct {
	db *gorm.DB
}

func NewBaseController() *BaseController {
	return &BaseController{db: database.New()}
}

// This method will be used to check user's authorization
func (bc BaseController) isAuthorizedForAction(w http.ResponseWriter, r *http.Request) bool {
	// Get the user id from the path
	vars := mux.Vars(r)
	userId := vars["id"]
	if userId != r.Header.Get("x-key") {
		util.WriteJSONResponse(w, http.StatusForbidden, util.JSON{"message": "Unauthorized access"})
		return false
	}
	return true
}

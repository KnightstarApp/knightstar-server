package routes

import (
	"knightstar/internal/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() http.Handler {

	r := mux.NewRouter()

	baseController := controllers.NewBaseController()
	userController := controllers.NewUserController(baseController)

	// General routes
	r.HandleFunc("/", baseController.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/health", baseController.HealthHandler).Methods("GET")

	// Subrouter for user routes
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", userController.GetAll).Methods("GET")
	userRouter.HandleFunc("", userController.Create).Methods("POST")
	userRouter.HandleFunc("/{id}", userController.Get).Methods("GET")
	userRouter.HandleFunc("/{id}", userController.Update).Methods("PUT")
	userRouter.HandleFunc("/{id}", userController.Delete).Methods("DELETE")

	return r
}

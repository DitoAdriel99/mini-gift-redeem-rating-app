package router

import (
	"go-learn/controller/auth"
	"go-learn/controller/product"
	"go-learn/middleware"
	"go-learn/repositories"
	"go-learn/service"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()
	//set dependency
	repo := repositories.NewRepo()
	serv := service.NewService(repo)
	// sdk := sdk.NewSDK()

	// storagex.CreateBucket()

	// middlewares
	tokenValidator := middleware.NewTokenValidator(*repo)

	// call controllers auth
	controllerLogin := auth.NewControllerLogin(*serv)
	controllerRegister := auth.NewControllerRegister(*serv)
	controllerStatus := auth.NewControllerStatus(*serv)

	// set logs
	// logsIntegrate := middleware.NewLogsIntegrate(*sdk)
	// router.Use(logsIntegrate.CreateLogs)
	// call controllers product
	controllerProductCreate := product.NewControllerProductCreate(*serv)
	controllerProductUpdate := product.NewControllerProductUpdate(*serv)
	controllerProductDelete := product.NewControllerProductDelete(*serv)
	controllerProductGetAll := product.NewControllerProductGetAll(*serv)
	controllerProductDetail := product.NewControllerProductDetail(*serv)
	controllerProductRating := product.NewControllerProductRating(*serv)
	controllerProductRedeem := product.NewControllerProductRedeem(*serv)

	//login
	router.HandleFunc("/login", controllerLogin.HandleLogin).Methods("POST")
	router.HandleFunc("/register", controllerRegister.HandleRegister).Methods("POST")
	router.HandleFunc("/update-status", controllerStatus.Status).Methods("PUT")

	//product for admin access
	adminRoutes := router.PathPrefix("").Subrouter()
	adminRoutes.Use(tokenValidator.ValidateTokenMiddleware("admin"))

	adminRoutes.HandleFunc("/gifts", controllerProductCreate.Create).Methods("POST")
	adminRoutes.HandleFunc("/gifts/{id}", controllerProductUpdate.Update).Methods("PUT")
	adminRoutes.HandleFunc("/gifts/{id}", controllerProductDelete.Delete).Methods("DELETE")

	// product for user and admin access
	generalRoutes := router.PathPrefix("").Subrouter()
	generalRoutes.Use(tokenValidator.ValidateTokenMiddleware("admin", "user"))

	generalRoutes.HandleFunc("/gifts", controllerProductGetAll.Get).Methods("GET")
	generalRoutes.HandleFunc("/gifts/{id}", controllerProductDetail.Detail).Methods("GET")
	generalRoutes.HandleFunc("/gifts/{id}/redeem", controllerProductRedeem.Redeem).Methods("POST")
	generalRoutes.HandleFunc("/gifts/{id}/rating", controllerProductRating.Rating).Methods("POST")

	return router
}

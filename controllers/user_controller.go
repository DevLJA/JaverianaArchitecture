package controllers

import (
	"CircuitBreaker-LauraJoya/controllers/utils"
	"CircuitBreaker-LauraJoya/services"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController interface {
	GetInfoClient(w http.ResponseWriter, r *http.Request)
	GetBasicInfoSimulated(w http.ResponseWriter, r *http.Request)
}

func NewUserController(userService services.UserService) UserController {
	return &userControllerImp{
		userService: userService,
	}
}

type userControllerImp struct {
	userService services.UserService
}

func (u userControllerImp) GetInfoClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	// Call without circuit breaker
	//user, err := u.userService.GetCompleteInfoClient(userID)
	user, err := u.userService.GetCompleteInfoClientWithCircuitBreaker(userID)

	if err != nil {
		utils.RespondWithBody(w, http.StatusServiceUnavailable, nil, err)
		return
	}

	utils.RespondWithBody(w, http.StatusOK, user, nil)
}

func (u userControllerImp) GetBasicInfoSimulated(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	email, err := u.userService.GetBasicInfoUser(userID)

	if err != nil {
		utils.RespondWithBody(w, http.StatusServiceUnavailable, nil, err)
		return
	}

	utils.RespondWithBody(w, http.StatusOK, email, nil)
}

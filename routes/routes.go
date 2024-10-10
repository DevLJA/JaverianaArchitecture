package routes

import (
	"CircuitBreaker-LauraJoya/controllers"
	"CircuitBreaker-LauraJoya/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetCustomHttpHandlers(userService services.UserService) *mux.Router {
	userController := controllers.NewUserController(userService)
	chessController := controllers.NewChessController()

	r := mux.NewRouter()
	r.HandleFunc("/status_client/{userID:[0-9]+}", userController.GetInfoClient).Methods("GET")
	r.HandleFunc("/user/{userID:[0-9]+}", userController.GetBasicInfoSimulated).Methods("GET")
	r.HandleFunc("/chess/image", chessController.GetImageFEN).Methods("POST")

	return r
}

func InitServer(router *mux.Router) {
	http.Handle("/", router)
	port := "8080"
	http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Printf("Server listening on http://localhost:%s/\n", port)
	}
}

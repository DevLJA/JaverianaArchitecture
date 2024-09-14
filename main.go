package main

import (
	customRoutes "CircuitBreaker-LauraJoya/routes"
	"CircuitBreaker-LauraJoya/services"
	"CircuitBreaker-LauraJoya/utils/circuit_braker"
	"time"
)

func main() {
	circuit_braker.NewCircuitBreaker(3, 10*time.Second)

	userService := services.NewUserService()
	router := customRoutes.SetCustomHttpHandlers(userService)
	customRoutes.InitServer(router)
}

package main

import (
	"auth/controller"
	"auth/service"
	"log"
	"net/http"
)

func main() {
	userServiceURL := "http://user-service:8085"
	jwtSecret := "your_jwt_secret"

	//userServiceURL := os.Getenv("USER_SERVICE_URL")
	//jwtSecret := os.Getenv("JWT_SECRET")

	authService := service.NewAuthService(userServiceURL, jwtSecret)
	authController := controller.NewAuthController(authService)

	http.HandleFunc("/login", authController.Login)
	http.HandleFunc("/signup", authController.Signup)

	log.Println("Running on port :8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}

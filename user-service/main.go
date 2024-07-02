package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user/controller"
	"user/repository"
	"user/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Dl!220695@tcp(user-db:3306)/usersapp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	log.Println("running on port :8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}

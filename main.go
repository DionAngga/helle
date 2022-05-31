package main

import (
	"fmt"
	"log"
	"net/http"

	"helle/auth"
	"helle/controller"
	"helle/entity/request"
	repo "helle/repository/database"
	"helle/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func main() {
	PORT := auth.PORT()
	r := mux.NewRouter()
	DNS := auth.DNS()
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&request.User{})
	userRepository := repo.New(db)
	userUsecase := usecase.New(userRepository)
	userController := controller.New(userUsecase)
	r.HandleFunc("/user/inquiry", userController.GetInquirybyaccount).Methods("POST")
	r.HandleFunc("/user/profile", userController.GetProfilebyUsername).Methods("POST")
	r.HandleFunc("/user/username_byaccount", userController.GetUsernameByAccount).Methods("POST")
	r.HandleFunc("/user/inquiry_hp_byaccount", userController.GetUserPhoneNumber).Methods("POST")

	log.Println("Database connected", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))

}

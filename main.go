package main

import (
	"fmt"
	"log"
	"net/http"

	"helle/controller"
	"helle/entity/request"
	repo "helle/repository/mysql_user"
	"helle/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

const port string = ":8080"

func initializeRoutes() {
	r := mux.NewRouter()

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/helle?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&request.User{})
	userRepository := repo.NewRepository(db)
	userUsecase := usecase.NewUser(userRepository)
	userController := controller.NewController(userUsecase)
	r.HandleFunc("/user/inquirys", userController.PostUser).Methods("POST")
	r.HandleFunc("/user/inquiry", userController.GetInquirybyaccount).Methods("POST")
	r.HandleFunc("/user/profile", userController.GetProfilebyUsername).Methods("POST")
	r.HandleFunc("/user/username_byaccount", userController.GetUsernameByAccount).Methods("POST")
	r.HandleFunc("/user/inquiry_hp_byaccount", userController.GetUserPhoneNumber).Methods("POST")

	log.Println("Database connected", port)
	log.Fatal(http.ListenAndServe(port, r))

}

func main() {
	initializeRoutes()
}

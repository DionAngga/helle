package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	accController "helle/controller/acc_controller"
	profileController "helle/controller/profile_controller"
	userController "helle/controller/user_controller"
	"helle/entity/request"

	//repositorymysql "helle/repository/database"
	tbluser "helle/repository/database/mysql/tbl_user"
	tbluseraccount "helle/repository/database/mysql/tbl_user_account"
	tbluserprofile "helle/repository/database/mysql/tbl_user_profile"
	usecase "helle/usecase/user_usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DNS")
}

func PORT() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PORT")
}

func main() {
	PORT := PORT()
	r := mux.NewRouter()
	DNS := DNS()
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&request.User{})
	userRepository := tbluser.New(db)
	userAccountRepository := tbluseraccount.New(db)
	userProfileRepository := tbluserprofile.New(db)
	userUsecase := usecase.New(userRepository, userAccountRepository, userProfileRepository)
	userController := userController.New(userUsecase)
	accController := accController.New(userUsecase)
	profileController := profileController.New(userUsecase)
	r.HandleFunc("/user/inquiry", userController.GetInquirybyaccount).Methods("POST")
	r.HandleFunc("/user/profile", profileController.GetProfilebyUsername).Methods("POST")
	r.HandleFunc("/user/username_byaccount", accController.GetUsernameByAccount).Methods("POST")
	r.HandleFunc("/user/inquiry_hp_byaccount", accController.GetUserPhoneNumber).Methods("POST")

	log.Println("Database connected", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))

}

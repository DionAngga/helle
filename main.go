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
	accUsercase "helle/usecase/acc_usecase"
	profileUsecase "helle/usecase/profile_usecase"
	userUsecase "helle/usecase/user_usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func main() {
	PORT := os.Getenv("PORT")
	r := mux.NewRouter()
	DNS := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&request.User{})
	userRepository := tbluser.New(db)
	userAccountRepository := tbluseraccount.New(db)
	userProfileRepository := tbluserprofile.New(db)
	userUsecase := userUsecase.New(userRepository, userAccountRepository, userProfileRepository)
	accUsecase := accUsercase.New(userAccountRepository)
	profileUsecase := profileUsecase.New(userProfileRepository)
	userController := userController.New(userUsecase)
	accController := accController.New(accUsecase)
	profileController := profileController.New(profileUsecase)
	r.HandleFunc("/user/inquiry", userController.GetInquirybyaccount).Methods("POST")
	r.HandleFunc("/user/profile", profileController.GetProfilebyUsername).Methods("POST")
	r.HandleFunc("/user/username_byaccount", accController.GetUsernameByAccount).Methods("POST")
	r.HandleFunc("/user/inquiry_hp_byaccount", userController.GetUserPhoneNumber).Methods("POST")

	log.Println("Database connected", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))

}

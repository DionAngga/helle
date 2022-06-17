package main

import (
	"fmt"
	"net/http"
	"os"

	accController "helle/controller/acc_controller"
	profileController "helle/controller/profile_controller"

	"github.com/sirupsen/logrus"

	//repositorymysql "helle/repository/database"

	tbluseraccount "helle/repository/database/mysql/tbl_user_account"
	tbluserprofile "helle/repository/database/mysql/tbl_user_profile"
	accUsercase "helle/usecase/acc_usecase"
	profileUsecase "helle/usecase/profile_usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	dns := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "level",
		},
	})
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	userAccountRepository := tbluseraccount.New(db)
	userProfileRepository := tbluserprofile.New(db)
	accUsecase := accUsercase.New(userAccountRepository)
	profileUsecase := profileUsecase.New(userProfileRepository, userAccountRepository)
	accController := accController.New(accUsecase)
	profileController := profileController.New(profileUsecase)
	r.HandleFunc("/user/profile_byprofile", profileController.GetProfilebyUsername).Methods("POST")
	r.HandleFunc("/user/username_byaccount", accController.GetUsernameByAccount).Methods("POST")
	r.HandleFunc("/user/inquiry_hp_byaccount", profileController.GetUserPhoneNumber).Methods("POST")
	log.Info("Database connected", port)
	log.Fatal(http.ListenAndServe(port, r))
}

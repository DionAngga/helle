package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	accController "helle/controller/acc_controller"
	profileController "helle/controller/profile_controller"

	loggers "helle/log"

	"github.com/sirupsen/logrus"

	//repositorymysql "helle/repository/database"

	tbluseraccount "helle/repository/database/mysql/tbl_user_account"
	tbluserprofile "helle/repository/database/mysql/tbl_user_profile"
	accUsercase "helle/usecase/acc_usecase"
	profileUsecase "helle/usecase/profile_usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gorilla/mux"
)

func init() {
	loggers.Init()
}

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	dns := os.Getenv("DNS")

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// if os.GetEnv("ENV") == "dev" {
	// 	log.SetLevel(logrus.DebugLevel)
	// } else {
	// 	log.SetLevel(logrus.InfoLevel)
	// }

	userAccountRepository := tbluseraccount.New(db)
	userProfileRepository := tbluserprofile.New(db)
	accUsecase := accUsercase.New(userAccountRepository)
	profileUsecase := profileUsecase.New(userProfileRepository, userAccountRepository)
	accController := accController.New(accUsecase)
	profileController := profileController.New(profileUsecase)
	r.HandleFunc("/user/profile_byprofile", profileController.GetProfilebyUsername).Methods("POST")
	r.HandleFunc("/user/username_byaccount", accController.GetUsernameByAccount).Methods("POST")
	r.HandleFunc("/user/inquiry_hp_byaccount", profileController.GetUserPhoneNumber).Methods("POST")

	if os.Getenv("ENV") == "development" {
		//fmt.Println("development mode")
		// loggers.LogDebug(`Server started on port : ` + port)
		logrus.Info("Server started on port : " + port)
	}

	fmt.Println("Hello World")
	log.Fatal(http.ListenAndServe(port, r))
}

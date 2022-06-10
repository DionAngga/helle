package migrate

import (
	"fmt"
	"helle/entity/database"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func migrate() {
	dns := os.Getenv("DNS")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&database.TblUserAccount{})
	if err != nil {
		fmt.Println("error")
	}

}

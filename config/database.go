package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GLOBAL CONFIG DB = GLOBAL
var DB *gorm.DB

func ConnectDb() error {
	dsn := "" + os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_ADDRESS") + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"
	// fmt.Println("USER:", os.Getenv("DATABASE_USER"))
	// fmt.Println("DB:", os.Getenv("DATABASE_NAME"))
	//CREATE DB CONNECT
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	//EXPORT DB FUNC
	DB = db
	return nil
}

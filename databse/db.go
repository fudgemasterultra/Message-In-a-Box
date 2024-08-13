package databse

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateMessageBox(message string, box int64) {
	DB.Create(&MessageBox{Message: message, Box: box})
}
func Connect() {
	dbUser := os.Getenv("USERNAME")
	dbPass := os.Getenv("PASSWORD")
	dbConnectionUrl := os.Getenv("CONNECTION")
	ConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=db port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbConnectionUrl, dbUser, dbPass)
	db, err := gorm.Open(postgres.Open(ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&MessageBox{})
	DB = db
}

package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

// func DBConnection() () {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := "my-secret-pw"
// 	dbName := "go_crud"
// 	dbHost := "172.19.0.3"
// 	dbPort := "3306"

// 	DB, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?tls=false")
// 	return DB, err
// }
func ConnectDatabase() {
	dbUser := "root"
	dbPass := "dbpass"
	dbName := "go_api"
	dbHost := "172.17.0.1"
	dbPort := "3306"
	database, err := gorm.Open(mysql.Open(dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?tls=false"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})

	DB =database
	
}
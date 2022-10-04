package db

import (
	"fmt"

	"github.com/quangghaa/grpc-demo/models/connection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	err    error
	host   = "localhost"
	port   = "3306"
	user   = "root"
	pass   = "root"
	dbname = "grpc_demo"
)

func Init() (*gorm.DB, error) {
	fmt.Println("Init database ...")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err = gorm.Open("mysql", username_mysql+":"+password_mysql+"@tcp("+host_mysql+":"+port_mysql+")/"+database_mysql+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&connection.Connection{},
	)
	return db, err
}

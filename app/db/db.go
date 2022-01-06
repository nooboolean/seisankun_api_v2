package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{}
}

type DatabaseManager struct{}

func (d *DatabaseManager) Connect() {
	DBMS := "mysql"
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	DBNAME := os.Getenv("DB_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	db.LogMode(true)
	db.SingularTable(true)

	DB = db

	if err != nil {
		panic(err.Error())
	}
}

func (d *DatabaseManager) Close() {
	if err := DB.Close(); err != nil {
		panic(err.Error())
	}
}

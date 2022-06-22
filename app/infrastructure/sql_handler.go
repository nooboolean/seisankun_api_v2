package infrastructure

import (
	"os"

	"github.com/nooboolean/seisankun_api_v2/interfaces/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() repositories.SqlHandler {
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := "tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")"
	DBNAME := os.Getenv("DB_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	conn, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}

	conn.Logger.LogMode(4)

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Update(sql string, value interface{}) *gorm.DB {
	return handler.Conn.Update(sql, value)
}

func (handler *SqlHandler) Updates(value interface{}) *gorm.DB {
	return handler.Conn.Updates(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func (handler *SqlHandler) Joins(query string, args ...interface{}) *gorm.DB {
	return handler.Conn.Joins(query, args...)
}

func (handler *SqlHandler) Model(value interface{}) *gorm.DB {
	return handler.Conn.Model(value)
}

func (handler *SqlHandler) Preload(column string, conditions ...interface{}) *gorm.DB {
	return handler.Conn.Preload(column, conditions...)
}

func (handler *SqlHandler) Transaction(fc func(tx *gorm.DB) error) (err error) {
	return handler.Conn.Transaction(fc)
}

func (handler *SqlHandler) Debug() *gorm.DB {
	return handler.Conn.Debug()
}

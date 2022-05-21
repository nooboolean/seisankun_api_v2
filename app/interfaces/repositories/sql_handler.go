package repositories

import "gorm.io/gorm"

type SqlHandler interface {
	Exec(string, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Update(string, interface{}) *gorm.DB
	Updates(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Joins(string, ...interface{}) *gorm.DB
	Model(interface{}) *gorm.DB
	Preload(string, ...interface{}) *gorm.DB
}

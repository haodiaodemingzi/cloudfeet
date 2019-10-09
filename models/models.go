package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/haodiaodemingzi/cloudfeet/pkgs/settings"
	"github.com/haodiaodemingzi/cloudfeet/utils"
)

var db *gorm.DB

// TODO: updatetime not auto update
type Model struct {
	ID         int             `gorm:"primary_key" json:"id"`
	CreateTime utils.LocalTime `gorm:"default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime utils.LocalTime `gorm:"default:CURRENT_TIMESTAMP" json:"update_time"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	connStr := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local`,
		settings.Config.MySQL.User,
		settings.Config.MySQL.Password,
		settings.Config.MySQL.Host,
		settings.Config.MySQL.Port,
		settings.Config.MySQL.DataBase)
	fmt.Printf("connect info : %s", connStr)
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.LogMode(settings.Config.MySQL.Debug)
	db.AutoMigrate(&PacModel{}, &ProxyModel{}, &UserModel{})
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

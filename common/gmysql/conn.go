package gmysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/haodiaodemingzi/cloudfeet/common/settings"
	"github.com/jmoiron/sqlx"
)

// DB mysql global
var DBI *sqlx.DB

// Setup init mysql
func Setup() {
	host := settings.Config.MySQL.Host
	dbName := settings.Config.MySQL.DataBase
	user := settings.Config.MySQL.User
	password := settings.Config.MySQL.Password
	port := settings.Config.MySQL.Port

	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	db, err := sqlx.Open("mysql", connstr)
	if err != nil {
		log.Println("conncet mysql error!!!")
		panic(err)
	}
	DBI = db
}

// GetDB get db handler
func GetDB() *sqlx.DB {
	fmt.Println(DBI)
	return DBI
}

package gmysql

import (
	"database/sql"
	"time"

	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/haodiaodemingzi/common/settings"
)

// DB mysql global
var DB *sql.DB

// Setup init mysql
func Setup() {
	var db *sql.DB
	host := settings.Config.MySQL.Host
	dbName := settings.Config.MySQL.DataBase
	user := settings.Config.MySQL.User
	password := settings.Config.MySQL.Password
	port := settings.Config.MySQL.Port

	db, err := manager.New(dbName, user, password, host).Set(
		manager.SetCharset("utf8"),
		manager.SetAllowCleartextPasswords(true),
		manager.SetInterpolateParams(true),
		manager.SetTimeout(5*time.Second),
		manager.SetReadTimeout(5*time.Second)).Port(port).Open(true)
	if err != nil {
		panic(err)
	}

	DB = db
}

// GetDB get db handler
func GetDB() *sql.DB {
	return DB
}

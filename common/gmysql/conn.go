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

// Insert inserts an array of data into table proxy
func InsertOne(table string, data map[string]interface{}) (int, error) {
	db := GetDB()
	sql := `insert into ` + table
	for k, _ := range data {
		sql += fmt.Sprintf("set %s=:%s", k, k)
	}
	_, err := db.NamedExec(sql, data)
	fmt.Printf("execute insert sql : %s", sql)
	fmt.Println(data)

	return 0, err
}

// Insert inserts an array of data into table proxy
func InsertMultiRows(table string, dataList []map[string]interface{}) (int, error) {
	db := GetDB()
	meta := dataList[0]
	sql := "insert into " + table

	for k, _ := range meta {
		sql += fmt.Sprintf("set %s=:%s", k, k)
	}

	tx := db.MustBegin()
	nstmt, err := db.PrepareNamed(sql)
	for _, item := range dataList {
		nstmt.MustExec(item)
	}
	err = tx.Commit()
	return 0, err
}

// Insert inserts an array of data into table proxy
func UpdateOne(table string, data map[string]interface{}) (int, error) {
	db := GetDB()
	sql := `update` + table
	for k, _ := range data {
		sql += fmt.Sprintf("set %s=:%s", k, k)
	}
	_, err := db.NamedExec(sql, data)
	fmt.Printf("execute insert sql : %s", sql)
	fmt.Println(data)

	return 0, err
}

// Insert inserts an array of data into table proxy
func UpdateMultiRows(table string, dataList []map[string]interface{}) (int, error) {
	db := GetDB()
	meta := dataList[0]
	sql := "update " + table

	for k, _ := range meta {
		sql += fmt.Sprintf("set %s=:%s", k, k)
	}

	tx := db.MustBegin()
	nstmt, err := db.PrepareNamed(sql)
	for _, item := range dataList {
		nstmt.MustExec(item)
	}
	err = tx.Commit()
	return 0, err
}

// get one row from db
func Get(table string, id int, result *interface{}) error {
	db := GetDB()
	sql := fmt.Sprintf(`select * from %s where id=? limit 1`, table)
	return db.Get(result, sql, id)
}

// get one row from db
func delete(table string, id int, result *map[string]interface{}) error {
	db := GetDB()
	sql := fmt.Sprintf(`delete  from %s where id=:id limit 1`, table)
	_, err := db.Queryx(sql, id)
	return err
}

// get one row from db
func Search(table string, args map[string]interface{}, result *[]interface{}) error {
	db := GetDB()
	sql := fmt.Sprintf(`select * from %s where`, table)
	nstmt, err := db.PrepareNamed(sql)
	nstmt.Exec(args)
	err = nstmt.Select(&result, args)
	return err
}

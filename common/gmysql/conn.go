package gmysql

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/haodiaodemingzi/cloudfeet/common/settings"
)

// DB mysql global
var DBI *sqlx.DB

type MySQL struct {
	Conn *sqlx.DB
}

// Setup init mysql
func Setup() {
	host := settings.Config.MySQL.Host
	dbName := settings.Config.MySQL.DataBase
	user := settings.Config.MySQL.User
	password := settings.Config.MySQL.Password
	port := settings.Config.MySQL.Port

	conner := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)
	db, err := sqlx.Open("mysql", conner)
	if err != nil {
		log.Println("connect mysql error!!!")
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
	sql := specSQL(table, "insert", data)
	fmt.Printf("execute insert sql : %s", sql)
	_, err := db.NamedExec(sql, data)
	fmt.Println(data)

	return 0, err
}

func specSQL(table string, op string, data ...map[string]interface{}) string {
	var sql = ""
	var keys []string

	for k, _ := range data[0] {
		keys = append(keys, k+"=:"+k)
	}
	if op == "insert" {
		sql += "insert into " + table + " set " + strings.Join(keys, ",")
	}
	if op == "update" {
		var conds []string
		for k, _ := range data[1] {
			conds = append(conds, k+"=:"+k)
		}
		sql += "update " + table + " set " + strings.Join(keys, ",") + " where " + strings.Join(conds, " and ")
	}
	return sql
}

// Insert inserts an array of data into table proxy
func InsertMultiRows(table string, dataList []map[string]interface{}) (int, error) {
	db := GetDB()
	meta := dataList[0]
	sql := "insert into " + table + " set "

	for k, _ := range meta {
		sql += fmt.Sprintf(" %s=:%s ", k, k)
	}

	tx := db.MustBegin()
	nstmt, err := db.PrepareNamed(sql)
	if nstmt != nil {
		for _, item := range dataList {
			nstmt.MustExec(item)
		}
	}
	err = tx.Commit()
	return 0, err
}

// Insert inserts an array of data into table proxy
func UpdateOne(table string, data map[string]interface{}, where map[string]interface{}) (int, error) {
	db := GetDB()
	sql := specSQL(table, "update", data, where)
	fmt.Printf("execute insert sql : %s", sql)
	_, err := db.NamedExec(sql, data)
	fmt.Println(data)

	return 0, err
}

// Insert inserts an array of data into table proxy
func UpdateMultiRows(table string, dataList []map[string]interface{}) (int, error) {
	db := GetDB()
	meta := dataList[0]
	sql := "update " + table + " set "

	for k, _ := range meta {
		sql += fmt.Sprintf("%s=:%s, ", k, k)
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
func GetOne(table string, id int, result *interface{}) error {
	db := GetDB()
	sql := fmt.Sprintf(`select * from %s where id=? limit 1`, table)
	return db.Get(result, sql, id)
}

// get one row from db
func DeleteOne(table string, id int, result *map[string]interface{}) error {
	db := GetDB()

	sql := fmt.Sprintf(`delete  from %s where id=:id limit 1`, table)
	_, err := db.Queryx(sql, id)
	return err
}

// get one row from db
func Search(table string, args map[string]interface{}, result *[]interface{}) error {
	return nil
}

// get one row from db
func Query(table string, args map[string]interface{}, result *[]interface{}) error {
	return nil
}

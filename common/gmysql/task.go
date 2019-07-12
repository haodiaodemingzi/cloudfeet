package gmysql

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"time"
)

// Task is a mapping object for task table in mysql
type Task struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	Command    string    `json:"command"`
	Args       string    `json:"args"`
	CreateTime time.Time `json:"create_time"`
	FinishTime time.Time `json:"finish_time"`
	Status     int8      `json:"status"`
}

func (t *Task) GetOne(where map[string]interface{}) (*Task, error) {
	db := GetDB()
	if nil == db {
		return nil, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect("task", where, nil)
	if nil != err {
		return nil, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()
	var ret *Task
	err = scanner.Scan(row, &ret)
	return ret, err
}

//GetMulti gets multiple records from table task by condition "where"
func (t *Task) GetMulti(where map[string]interface{}) ([]*Task, error) {
	db := GetDB()
	if nil == db {
		return nil, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildSelect("task", where, nil)
	if nil != err {
		return nil, err
	}
	row, err := db.Query(cond, vals...)
	if nil != err || nil == row {
		return nil, err
	}
	defer row.Close()
	var ret []*Task
	err = scanner.Scan(row, &ret)
	return ret, err
}

//Insert inserts an array of data into table task
func (t *Task) Insert(data []map[string]interface{}) (int64, error) {
	db := GetDB()
	if nil == db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildInsert("task", data)
	if nil != err {
		return 0, err
	}
	retult, err := db.Exec(cond, vals...)
	if nil != err || nil == retult {
		return 0, err
	}
	return retult.LastInsertId()
}

//Update updates the table task
func (t *Task) Update(where, data map[string]interface{}) (int64, error) {
	db := GetDB()
	if nil == db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildUpdate("task", where, data)
	if nil != err {
		return 0, err
	}
	retult, err := db.Exec(cond, vals...)
	if nil != err {
		return 0, err
	}
	return retult.RowsAffected()
}

// Delete deletes matched records in task
func (t *Task) Delete(where, data map[string]interface{}) (int64, error) {
	db := GetDB()
	if nil == db {
		return 0, errors.New("sql.DB object couldn't be nil")
	}
	cond, vals, err := builder.BuildDelete("task", where)
	if nil != err {
		return 0, err
	}
	retult, err := db.Exec(cond, vals...)
	if nil != err {
		return 0, err
	}
	return retult.RowsAffected()
}

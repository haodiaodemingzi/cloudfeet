package gmysql

import (
	"testing"

	"github.com/haodiaodemingzi/cloudfeet/common/settings"
	"github.com/stretchr/testify/assert"
)

func init() {
	settings.Setup()
	Setup()
}

var schema = `
CREATE TABLE person (
    first_name varchar(200),
    last_name varchar(200),
    email varchar(200)
); `

var deleteSQL = `
delete from person;
`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func TestCRUD(t *testing.T) {
	a := assert.New(t)

	var db = GetDB()

	// create test table
	_, err := db.Exec(schema)
	if err != nil {
		db.MustExec(deleteSQL)
	}

	data := make(map[string]interface{})
	data["first_name"] = "james_first"
	data["last_name"] = "james_last"
	data["email"] = "qqqqq@qq.com"
	_, err = InsertOne("person", data)
	a.NoError(err, "test insert person row success")

	updateData := map[string]interface{}{
		"first_name": "king1111",
	}
	where := map[string]interface{}{
		"first_name": "james_first",
	}
	_, err = UpdateOne("person", updateData, where)
	a.NoError(err, "test update person row success")

	db.MustExec(deleteSQL)
}

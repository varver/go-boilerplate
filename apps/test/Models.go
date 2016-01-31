package test

import (
	//	"errors"
	//	"fmt"
	"github.com/go-xorm/xorm"
	//    "strconv"
	//    "strings"
	//    "time"
)

// TestTable is a sample struct which will be used to represent a table in database.
type TestTable struct {
	Id       int64
	Name     string
	Location string `xorm:"longtext not null"` // this is how you can apply some custom stuff .
	Salary   int64
}

// Update is a sample function that will update a record stored in table created by struct TestTable.
func (c *TestTable) Update(db *xorm.Engine) error {

	_, err := db.Update(c)
	if err != nil {
		return err
	}
	return nil
}

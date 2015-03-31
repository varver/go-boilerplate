package models

import (
	//	"errors"
	//	"fmt"
	"github.com/go-xorm/xorm"
	//    "strconv"
	//    "strings"
	//    "time"
)

// CoverPage type stores the data related to a page and map it with database table called "cover_page"
type TestTable struct {
	Id       int64
	Name     string
	Location string `xorm:"longtext not null"` // actual html created is stored here
	Salary   int64
}

// Update will update a coverpage stored in database.
func (c *TestTable) Update(db *xorm.Engine) error {

	_, err := db.Update(c)
	if err != nil {
		return err
	}
	return nil
}

package middleware

import (
	"github.com/varver/go-boilerplate/apps/test"
	"github.com/varver/go-boilerplate/conf"

	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/ziutek/mymysql/godrv"
)

// DB is used to make connection to database . You can pass database type and its access information like
// path to database or database name database user and database connection password.
// You can use , mysql , postgres, sqlite , orcale , ms-sql.
func DB() martini.Handler {

	/*
	 * If you plan to use sqllite then make the following changes in config.ini
	 * DbDriver	= "sqlite3"
	 * DbSource	= "test.db"
	 *
	 * and import the sqllite driver
	 * _ "github.com/mattn/go-sqlite3"
	 *
	 */

	engine, err := xorm.NewEngine(conf.Setting.DbDriver, conf.Setting.DbSource)
	if err != nil {
		panic(err)
	}

	//engine.ShowSQL = true
	engine.Sync(
		&test.TestTable{},
	)

	return func(c martini.Context) {
		c.Map(engine)
		c.Next()
	}
}

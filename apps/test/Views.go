package test

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	"net/http"
)

//DisplayTestPage is a test controller or view that
func DisplayTestPage(resp http.ResponseWriter, req *http.Request, db *xorm.Engine, params martini.Params) string {
	resp.WriteHeader(200)
	return fmt.Sprintf("This is a response")
}

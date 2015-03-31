package controllers

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/go-xorm/xorm"
	"net/http"
)

func DisplayTestPage(resp http.ResponseWriter, req *http.Request, db *xorm.Engine, params martini.Params) string {
	resp.WriteHeader(200)
	return fmt.Sprintf("This is a response")
}

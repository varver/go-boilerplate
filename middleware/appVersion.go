package middleware

import (
	"fmt"
	"github.com/varver/go-boilerplate/conf"
	"net/http"

	"github.com/go-martini/martini"
)

var appVersion = func(resp http.ResponseWriter, req *http.Request) {
	headers := resp.Header()
	headers.Set("Server", fmt.Sprintf("%s %s", conf.Setting.ServerAppName, conf.Setting.ServerAppVersion))
	headers.Set("X-App", conf.Setting.ServerAppTagline)
}

//AppVersion returns the latest App Version of Your Project
func AppVersion() martini.Handler {
	return appVersion
}

package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/varver/go-boilerplate/conf"
	"github.com/varver/go-boilerplate/controllers"
	"github.com/varver/go-boilerplate/middleware"
	"github.com/varver/go-boilerplate/utils/logger"
	"math/rand"
	"runtime"
	"time"
)

const TimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"
const (
	LIVE = "live"
	DEV  = "dev"
)

//var store = sessions.NewCookieStore([]byte("42"))
func StartServer() {

	rand.Seed(time.Now().UTC().UnixNano())
	m := martini.Classic()
	m.Use(sessions.Sessions(conf.Setting.CookieName, sessions.NewCookieStore([]byte(conf.Setting.CookieSecret))))
	m.Use(middleware.DB())
	m.Use(middleware.AppVersion())
	m.Use(martini.Static("static"))
	static := martini.Static("static", martini.StaticOptions{Fallback: "templates/404.html"})

	// enter urls below //
	m.Get("/test", controllers.DisplayTestPage)

	// final call to initiate all the stuff
	m.RunOnAddr(":" + conf.Setting.ServerPort)
	m.NotFound(static)
	m.Run()
}

// setEnvironmentVariables set all environment cariables before doing anything else .
func setEnvironmentVariables() {
	// martini running mode
	if conf.Setting.EnvMode == LIVE {
		martini.Env = "production"
	}
	// set max go parallel processes
	runtime.GOMAXPROCS(conf.Setting.GoMaxProcs)
}

func main() {
	setEnvironmentVariables()
	StartServer()
}

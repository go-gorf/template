package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/auth"
	"github.com/go-gorf/gorf"
)

import "template/apps/hello"

// add all the apps
var apps = []gorf.GorfApp{
	&auth.AuthApp,
	&hello.HelloApp,
}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
	gorf.Settings.DbConf = &gorf.SqliteBackend{
		Name: "db.sqlite",
	}
	// app settings
	// auth.AuthSettings.NewUserState = auth.AuthState(true)
}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	gorf.Apps = append(apps)
	LoadSettings()
	gorf.InitializeDatabase()
	gorf.SetupApps()
	r := gin.Default()
	gorf.RegisterApps(r)
	return r
}

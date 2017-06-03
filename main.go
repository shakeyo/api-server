package main

import (
	_ "api-server/routers"

	"github.com/astaxie/beego"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

const (
	APP_VER = "0.1.1.0603"
)

func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		fmt.Println("Shutdown gracefully, bye...")
		// do graceful shutdown
	}

	os.Exit(0)
}

func main() {

	dev := beego.BConfig.RunMode == "dev"
	level := beego.LevelInformational
	if !dev{
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}else{
		level = beego.LevelDebug
	}

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.SetLevel(beego.AppConfig.DefaultInt("loglevel", level))
	beego.Info(beego.BConfig.AppName, APP_VER)

	//beego.ErrorController(&controllers.ErrorController{})

	graceful, _ := beego.AppConfig.Bool("graceful")
	if !graceful {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go handleSignals(sigs)
	}

	beego.Run()
}

package models

import (
//"database/sql"

"github.com/astaxie/beego"
_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

var dbEngine *xorm.Engine

func DB() *xorm.Engine {
	return dbEngine
}

func init() {

	dev := beego.BConfig.RunMode == "dev"
	url := beego.AppConfig.String("sql::url")

	var err error
	dbEngine, err = xorm.NewEngine("mssql", url)
	if err != nil{
		panic(err)
	}

	//engine.SetLogger(beego.BeeLogger)
	dbEngine.ShowSQL(true)

	var defaultLvl core.LogLevel
	if dev{
		defaultLvl = core.LOG_DEBUG
	}else{
		defaultLvl = core.LOG_INFO
	}

	level := core.LogLevel(beego.AppConfig.DefaultInt("sql:loglevel", int(defaultLvl)))
	dbEngine.Logger().SetLevel(core.LogLevel(level))
}



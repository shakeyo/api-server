package models

import (
	"errors"
	"github.com/astaxie/beego"
	//"encoding/binary"
	"strconv"
)

type SystemSettings struct {
	Name   string	`xorm:"not null default '' unique VARCHAR(32)"`
	IntVal int	`xorm:"not null default '' unique VARCHAR(191)"`
	StrVal string	`xorm:"not null default '' unique VARCHAR(512)"`
}

func init(){

}


func LoadAll() (table []SystemSettings, err error) {

	results, err := dbEngine.Query("select StatusName,StatusValue,StatusString from SystemStatusInfo")
	if err != nil {
		beego.Warn("exec sql failed:", err)
		return nil, err
	}

	if len(results) == 0{
		return nil, errors.New("no records")
	}

	var settings []SystemSettings
	for _, val := range results {

		value,_ := strconv.Atoi(string(val["StatusValue"]))
		item := SystemSettings{
			string(val["StatusName"]),
			value,
			string(val["StatusString"]),
		}

		beego.Debug("settings:", item)
		settings = append(settings, item)
	}

	return settings, nil
}
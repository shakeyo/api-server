package models

import (
	"errors"
	"api-server/data"
	"github.com/astaxie/beego"
	//"encoding/binary"
	"strconv"
)

type SystemSettings struct {
	Name   string
	IntVal int
	StrVal string
}

func init(){

}


func LoadSystemSettings() (table []SystemSettings, err error) {

	results, err := data.DB().Query("select StatusName,StatusValue,StatusString from SystemStatusInfo")
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
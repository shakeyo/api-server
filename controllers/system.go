package controllers

import (
	"api-server/models"
	//"encoding/json"

	"github.com/astaxie/beego"
	//"api-server/data"
	"github.com/astaxie/beego/cache"
)

// Operations about System
type SystemController struct {
	BaseController
}

var cacheInstance cache.Cache

func init(){
	/*cacheInstance, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil{
		panic(err)
	}*/
}

// @Title GetClientSettings
// @Description update the user
// @Success 200 {object}
// @Failure 500 :server error
// @router / [get]
func (u *SystemController) GetClientSettings() {
	data, err := models.LoadAll()
	if err != nil{
		u.RetError(errDatabase)
	}

	beego.Debug("Settings:", &data)
	u.Data["json"] = data
	u.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *SystemController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

func (u *SystemController) GetNotice() {

}
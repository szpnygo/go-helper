package controllers

import (
	"github.com/astaxie/beego"
	"neobaran.com/neo/beego/base"
)

// CheckController controllers
type CheckController struct {
	base.BController
}

// CreateLog create memo app log
func (c *CheckController) Check() {
	c.SuccessJSON(beego.AppConfig.String("version"))
}
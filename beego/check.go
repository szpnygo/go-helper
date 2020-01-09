package beego

import (
	"github.com/astaxie/beego"
)

// CheckController controllers
type CheckController struct {
	BController
}

// CreateLog create app log
func (c *CheckController) Check() {
	c.SuccessJSON(beego.AppConfig.String("version"))
}
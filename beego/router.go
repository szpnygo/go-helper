package beego

import (
	"github.com/astaxie/beego"
)

func CheckRouter() beego.LinkNamespace {
	return beego.NSRouter("/check", &CheckController{}, "*:Check")
}

func ApiFilter() {
	beego.InsertFilter("/*", beego.AfterExec, ApiLogFilter, false, false)
}

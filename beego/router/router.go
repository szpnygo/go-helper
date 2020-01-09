package router

import (
	check "tools/neo/beego/controllers"
	"tools/neo/beego/filter"

	"github.com/astaxie/beego"
)

func CheckRouter() beego.LinkNamespace {
	return beego.NSRouter("/check", &check.CheckController{}, "*:Check")
}

func ApiFilter() {
	beego.InsertFilter("/*", beego.AfterExec, filter.ApiLogFilter, false, false)
}

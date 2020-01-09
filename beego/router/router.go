package router

import (
	check "neobaran.com/neo/beego/controllers"
	"neobaran.com/neo/memo/filter"

	"github.com/astaxie/beego"
)

func CheckRouter() beego.LinkNamespace {
	return beego.NSRouter("/check", &check.CheckController{}, "*:Check")
}

func ApiFilter() {
	beego.InsertFilter("/*", beego.AfterExec, filter.ApiLogFilter, false, false)
}

package beego

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func CheckRouter() beego.LinkNamespace {
	return beego.NSRouter("/check", &CheckController{}, "*:Check")
}

func ApiFilter() {
	beego.InsertFilter("/*", beego.AfterExec, ApiLogFilter, false, false)
}

func LogApiFilter(header []string) {
	beego.InsertFilter("/*", beego.AfterExec, func(context *context.Context) {
		ApiLogFilter2(context, header)
	}, false, false)
}

func JsonFilter() {
	beego.InsertFilter("/*", beego.BeforeExec, JsonRequestFilter, false, false)
}

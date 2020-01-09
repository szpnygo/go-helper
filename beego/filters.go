package beego

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/url"
)

type ApiLog struct {
	Ip       string
	Url      string
	Header   interface{}
	Method   string
	UserId   int
	Form     url.Values
	Response interface{}
}

var ApiLogFilter = func(ctx *context.Context) {
	go apiLog(ctx)
}

func apiLog(ctx *context.Context) {
	uidID := ctx.Input.GetData("uid")
	if uidID != nil {
		uidID = uidID.(int)
	} else {
		uidID = -1
	}

	apiLog := ApiLog{
		Ip:       ctx.Input.IP(),
		Url:      ctx.Input.URL(),
		Header:   ctx.Input.Context.Request.Header,
		Method:   ctx.Input.Method(),
		UserId:   uidID.(int),
		Form:     ctx.Input.Context.Request.Form,
		Response: ctx.Input.GetData("response"),
	}

	response, err := json.MarshalIndent(apiLog, "", "    ")
	if err == nil {
		logs.Info("\n" + string(response) + "\n")
	}
}

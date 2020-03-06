package beego

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"net/url"
	"strconv"
	"strings"
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

var ApiLogFilter2 = func(ctx *context.Context, header []string) {
	go apiLog2(ctx, header)
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

func apiLog2(ctx *context.Context, headers []string) {
	uidID := ctx.Input.GetData("uid")
	if uidID != nil {
		uidID = uidID.(int)
	} else {
		uidID = -1
	}

	showHeader := map[string]string{}
	rh := ctx.Input.Context.Request.Header
	for _, h := range headers {
		if _, ok := rh[h]; ok {
			if len(rh[h]) > 0 {
				showHeader[h] = rh[h][0]
			}
		}
	}

	apiLog := ApiLog{
		Ip:       ctx.Input.IP(),
		Url:      ctx.Input.URL(),
		Method:   ctx.Input.Method(),
		Header:   showHeader,
		UserId:   uidID.(int),
		Form:     ctx.Input.Context.Request.Form,
		Response: ctx.Input.GetData("response"),
	}

	response, err := json.MarshalIndent(apiLog, "", "    ")
	if err == nil {
		logs.Info("\n" + string(response) + "\n")
	}
}

var JsonRequestFilter = func(ctx *context.Context) {

	var m interface{}
	err := json.Unmarshal(ctx.Input.RequestBody, &m)
	if err == nil && m != nil {
		var result map[string]interface{}
		result = m.(map[string]interface{})
		for k, v := range result {
			switch v.(type) {
			case string:
				ctx.Input.SetParam(k, v.(string))
			case int:
				ctx.Input.SetParam(k, strconv.Itoa(v.(int)))
			case bool:
				ctx.Input.SetParam(k, strconv.FormatBool(v.(bool)))
			case float64:
				str := fmt.Sprintf("%v", v)
				if strings.Contains(str, ".") {
					ctx.Input.SetParam(k, str)
				} else {
					ctx.Input.SetParam(k, strconv.Itoa(int(v.(float64))))
				}
			}
		}
	}
}

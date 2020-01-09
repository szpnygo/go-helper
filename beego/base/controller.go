package base

import (
	"github.com/astaxie/beego"
)

// Struct is the base struct for json response
type Struct struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// BController is the base controller
type BController struct {
	beego.Controller
}

// IsLogin ...
func (bc *BController) IsLogin() bool {
	uid := bc.Ctx.Input.GetData("uid")
	if uid != nil {
		return true
	}
	return false
}

// UserID ...
func (bc *BController) UserID() int {
	uidID := bc.Ctx.Input.GetData("uid")
	if uidID != nil {
		return uidID.(int)
	}
	return -1
}

func (bc *BController) App() string{
	return bc.Ctx.Input.Header("App")
}

func (bc *BController) DeviceID() string{
	return bc.Ctx.Input.Header("X-Device-Id")
}

// ResultSuccess just return code 0
func (bc *BController) ResultSuccess(data interface{}) Struct {
	return bc.Result(0, data, "")
}

// Result is the common response
func (bc *BController) Result(code int, data interface{}, message string) Struct {
	return Struct{Code: code, Data: data, Message: message}
}

// SuccessJSON just return code 0 with json
func (bc *BController) SuccessJSON(data interface{}) {
	bc.ResultJSON(0, data, "")
}

// ResultJSON return with json
func (bc *BController) ResultJSON(code int, data interface{}, message string) {
	bc.Data["json"] = bc.Result(code, data, message)
	bc.Ctx.Input.SetData("response", bc.Result(code, data, message))
	bc.Ctx.ResponseWriter.Header().Set("Version", beego.AppConfig.String("version"))
	bc.ServeJSON()
}
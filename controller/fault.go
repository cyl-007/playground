package controller

import (
	"apodemakeles/playground/log"
	"apodemakeles/playground/settings"
	"strconv"

	"github.com/kataras/iris/v12"
)

var faultLogger = log.RegisterScope("fault")

func Fault(ctx iris.Context) {
	body, err := ctx.GetBody()
	if err != nil {
		responseError(ctx, err)
	}
	code, err := strconv.ParseInt(string(body), 10, 32)
	if err != nil {
		responseError(ctx, err)
	}
	settings.SetFault(int(code))
	faultLogger.Infof("inject faultProcess %d successfully\r\n", code)

	responseOK(ctx)
}

package controller

import (
	"apodemakeles/playground/log"
	"apodemakeles/playground/settings"
	"apodemakeles/playground/utils"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

var echoLogger = log.RegisterScope("echo")

func Echo(ctx iris.Context) {
	if fault, has := settings.GetFault(); has {
		faultProcess(ctx, fault)
	} else {
		fineProcess(ctx)
	}
}

func fineProcess(ctx iris.Context) {
	msg := ctx.URLParam("msg")
	builder := strings.Builder{}
	builder.WriteString("version=")
	builder.WriteString(utils.GetVersion())
	builder.WriteString(",ip=")
	builder.WriteString(utils.GetIp())
	builder.WriteString(",timestamp=")
	builder.WriteString(strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
	builder.WriteString(",message=")
	builder.WriteString(msg)
	builder.WriteString("\r\n")
	printHeader := ctx.URLParam("header")
	if printHeader == "true" {
		for k, v := range ctx.Request().Header {
			builder.WriteString(k)
			builder.WriteString(": ")
			builder.WriteString(strings.Join(v, ","))
			builder.WriteString("\r\n")
		}
	}

	echoLogger.Info(builder.String())

	responseMessage(ctx, builder.String())
}

func faultProcess(ctx iris.Context, code int) {
	msg := ctx.URLParam("msg")
	builder := strings.Builder{}
	builder.WriteString("version=")
	builder.WriteString(utils.GetVersion())
	builder.WriteString(",ip=")
	builder.WriteString(utils.GetIp())
	builder.WriteString(",timestamp=")
	builder.WriteString(strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
	builder.WriteString(",message=")
	builder.WriteString(msg)
	builder.WriteString(", fault=")
	builder.WriteString(strconv.Itoa(code))
	builder.WriteString("\r\n")

	echoLogger.Info(builder.String())

	responseFault(ctx, code, builder.String())
}

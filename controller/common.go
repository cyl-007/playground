package controller

import "github.com/kataras/iris/v12"

func responseOK(ctx iris.Context) {
	ctx.StatusCode(200)
}

func responseMessage(ctx iris.Context, msg string) {
	ctx.StatusCode(200)
	ctx.Text(msg)
}

func responseObject(ctx iris.Context, obj interface{}) {
	ctx.StatusCode(200)
	ctx.ContentType("application/json")
	ctx.JSON(obj)
}

func responseFault(ctx iris.Context, code int, msg string) {
	ctx.StatusCode(code)
	ctx.Text(msg)
}

func responseError(ctx iris.Context, err error) {
	ctx.StatusCode(500)
	ctx.Text(err.Error())
}

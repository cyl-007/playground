package controller

import "github.com/kataras/iris/v12"

func Goto(ctx iris.Context) {
	ctx.Redirect("/echo", 302)
}

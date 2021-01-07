package controller

import "github.com/kataras/iris/v12"

func CreateLog(ctx iris.Context) {

	m := make(map[string]interface{})
	ctx.ReadJSON(&m)

	responseOK(ctx)

}

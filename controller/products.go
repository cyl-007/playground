package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

func CreateProducts(ctx iris.Context) {
	m := make(map[string]interface{})
	ctx.ReadJSON(&m)

	//builder := strings.Builder{}
	//builder.WriteString("==============================\n")
	//for k, v := range ctx.Request().Header {
	//	builder.WriteString(k)
	//	builder.WriteString(": ")
	//	builder.WriteString(strings.Join(v, ","))
	//	builder.WriteString("\n")
	//}
	//builder.WriteString("remote addr: " + ctx.RemoteAddr() + "\n")
	//fmt.Println(builder.String())

	fmt.Printf("query: %v \r\n", ctx.URLParams())

	responseObject(ctx, m)
}

func GetProducts(ctx iris.Context) {
	id := ctx.Params().Get("id")
	fail := ctx.URLParam("fail")
	builder := strings.Builder{}
	builder.WriteString("==============================\n")
	for k, v := range ctx.Request().Header {
		builder.WriteString(k)
		builder.WriteString(": ")
		builder.WriteString(strings.Join(v, ","))
		builder.WriteString("\n")
	}
	builder.WriteString("remote addr: " + ctx.RemoteAddr() + "\n")
	fmt.Println(builder.String())

	result := map[string]string{
		"abc": "123",
		"efg": "567",
		"id":  id,
	}

	if fail == "" {
		responseObject(ctx, result)
		return
	}
	code, _ := strconv.Atoi(fail)
	if code == 200 || code == 0 {
		responseObject(ctx, result)
	} else {
		responseFault(ctx, code, fail)
	}
}

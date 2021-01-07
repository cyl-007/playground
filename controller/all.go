package controller

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
)

func All(ctx iris.Context) {

	sb := strings.Builder{}
	sb.WriteString("path:")
	sb.WriteString(ctx.Path())
	sb.WriteString("\r\n")
	sb.WriteString(fmt.Sprintf("query: %v \r\n", ctx.URLParams()))
	sb.WriteString(fmt.Sprintf("header: %v \r\n", ctx.Request().Header))

	fmt.Println(sb.String())

	responseMessage(ctx, sb.String())
}

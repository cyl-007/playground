package main

import (
	"apodemakeles/playground/controller"
	"fmt"
	"github.com/google/gops/agent"
	"github.com/kataras/iris/v12"
	"net/http"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		fmt.Sprintf("agent.Listen err: %v", err)
	}

	//go func() {
	//	for {
	//		time.Sleep(5 * time.Second)
	//		fmt.Println("aaaaaaaaaaaaaaaa")
	//	}
	//}()

	app := iris.New()

	app.Get("/echo", controller.Echo)
	app.Put("/fault", controller.Fault)
	app.Get("/products/{id}", controller.GetProducts)
	app.Post("/products/{id}", controller.CreateProducts)
	app.Post("/logs", controller.CreateLog)
	app.Get("/goto", controller.Goto)
	app.Get("/mute", controller.Mute)
	app.Get("/*", controller.All)
	app.Post("/*", controller.All)
	app.Head("/*", controller.All)

	app.Run(iris.Server(&http.Server{Addr: ":8082"}))
}

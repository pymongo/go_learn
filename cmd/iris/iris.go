package main

import (
	"github.com/kataras/iris/v12"
)

// How to structure Go application to produce multiple binaries?
// https://stackoverflow.com/questions/50904560/how-to-structure-go-application-to-produce-multiple-binaries/50904959
func main() {
	app := iris.New()
	app.Get("/ping", func(ctx iris.Context) {
		//ctx.JSON(iris.Map{
		//	"message": "pong",
		//})
		// 每秒处理20000个请求，比actix/gin的17000要快
		ctx.Text("pong")
	})
	app.Listen(":8888")
	//app.Listen("0.0.0.0:8888")
}

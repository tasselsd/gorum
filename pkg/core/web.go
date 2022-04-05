package core

import "github.com/kataras/iris/v12"

var APP *iris.Application

func StartEngine() {
	APP = iris.Default()
	APP.Get("/demo", func(ctx iris.Context) {
		ctx.Markdown([]byte("# Title \n 1. first \n 2. second \n 3. third \n 4. four \n 5. five" + DB.Name()))
	})
	APP.Listen(":8080")
}

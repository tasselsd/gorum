package api

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/pkg/core"
)

var (
	GET    = make(map[string](func(iris.Context)))
	POST   = make(map[string](func(iris.Context)))
	PUT    = make(map[string](func(iris.Context)))
	DELETE = make(map[string](func(iris.Context)))
)

func StartEngine() {
	app := iris.Default()
	app.Logger().SetLevel(core.CFG.String("log.level"))
	app.Get("/generate_204", func(ctx iris.Context) {
		ctx.StatusCode(204)
	})
	for relativePath, route := range GET {
		app.Get(relativePath, route)
	}
	for relativePath, route := range POST {
		app.Post(relativePath, route)
	}
	for relativePath, route := range PUT {
		app.Put(relativePath, route)
	}
	for relativePath, route := range DELETE {
		app.Delete(relativePath, route)
	}
	app.Listen(fmt.Sprintf(":%d", core.CFG.Int("server.port")))
}

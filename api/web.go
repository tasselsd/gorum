package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/assets"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
)

var (
	GET    = make(map[string](func(iris.Context)))
	POST   = make(map[string](func(iris.Context)))
	PUT    = make(map[string](func(iris.Context)))
	DELETE = make(map[string](func(iris.Context)))
)

func StartEngine() {
	app := iris.Default()
	app.Use(sessionAware)
	app.WrapRouter(assetsRouter)
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

func sessionAware(ctx iris.Context) {
	token := ctx.GetCookie("token")
	var s = session.NaS
	if len(token) > 0 {
		ss, _ := session.SessionFromToken(token)
		s = ss
	}
	if s != nil && len(s.Avatar) == 0 {
		s.Avatar = core.CFG.Site.DefaultAvatar
	}
	ctx.Values().Set("session", s)
	ctx.Next()
}

func assetsRouter(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	if strings.HasPrefix(r.URL.Path, "/assets") {
		http.FileServer(http.FS(assets.FS)).ServeHTTP(w, r)
		return
	}
	router.ServeHTTP(w, r)
}

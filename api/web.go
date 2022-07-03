package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/assets"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/templates"
)

var (
	G      = make(map[string](func(iris.Context)))
	PO     = make(map[string](func(iris.Context)))
	PU     = make(map[string](func(iris.Context)))
	D      = make(map[string](func(iris.Context)))
	GET    = make(map[string](func(iris.Context)))
	POST   = make(map[string](func(iris.Context)))
	PUT    = make(map[string](func(iris.Context)))
	DELETE = make(map[string](func(iris.Context)))
	app    *iris.Application
)

func StartEngine() {
	app = iris.Default()
	app.Logger().SetLevel(core.CFG.String("log.level"))
	app.Use(applyNavStack)
	app.Use(applySiteProperties)
	app.Use(applyAuthentication)
	app.Use(applyIpRateLimiter)
	app.WrapRouter(embedAssets)
	app.Get("/generate_204", generate_204)
	app.OnErrorCode(iris.StatusNotFound, response_404)
	app.RegisterView(iris.Django(http.FS(templates.FS), ".html"))
	prepareRouters()
	prepareWs()
	app.Listen(fmt.Sprintf(":%d", core.CFG.Int("server.port")))
}

func prepareRouters() {
	for p, r := range G {
		app.Get(p, authenticationRequired, r)
	}
	for p, r := range PO {
		app.Post(p, authenticationRequired, r)
	}
	for p, r := range PU {
		app.Put(p, authenticationRequired, r)
	}
	for p, r := range D {
		app.Delete(p, authenticationRequired, r)
	}
	for p, r := range GET {
		app.Get(p, r)
	}
	for p, r := range POST {
		app.Post(p, r)
	}
	for p, r := range PUT {
		app.Put(p, r)
	}
	for p, r := range DELETE {
		app.Delete(p, r)
	}
}

func embedAssets(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	if strings.HasPrefix(r.URL.Path, "/assets") {
		http.FileServer(http.FS(assets.FS)).ServeHTTP(w, r)
		return
	}
	router.ServeHTTP(w, r)
}

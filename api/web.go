package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
	"github.com/tasselsd/gorum/assets"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
	"github.com/tasselsd/gorum/templates"
	"github.com/tomasen/realip"
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
	app.Use(navStack)
	app.Use(loadAuthentication)
	app.Use(ipRateLimiter)
	app.WrapRouter(registerAssets)
	app.Get("/generate_204", generate_204)
	app.OnErrorCode(iris.StatusNotFound, response_404)
	app.Logger().SetLevel(core.CFG.String("log.level"))
	app.RegisterView(iris.Django(http.FS(templates.FS), ".html"))
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
	app.Listen(fmt.Sprintf(":%d", core.CFG.Int("server.port")))
}

func generate_204(ctx iris.Context) {
	ctx.StatusCode(204)
}

func response_404(ctx iris.Context) {
	ctx.ViewData("statusCode", iris.StatusNotFound)
	ctx.ViewData("detail", "Not Found")
	ctx.View("failed")
}

func navStack(ctx iris.Context) {
	stack := core.NewNavStack()
	ctx.Values().Set("nav", stack)
	ctx.ViewData("nav", stack)
	ctx.Next()
}

func authenticationRequired(ctx iris.Context) {
	s := _sessionFromToken(ctx)
	if s == nil {
		ctx.Redirect(fmt.Sprintf("/signin?l=%s", ctx.Request().RequestURI), iris.StatusSeeOther)
		return
	}
	if s != nil && len(s.Avatar) == 0 {
		s.Avatar = core.CFG.Site.DefaultAvatar
	}
	ctx.Values().Set("session", s)
	ctx.Next()
}
func _sessionFromToken(ctx iris.Context) *session.Session {
	token := ctx.GetCookie("token")
	var s = session.NaS
	if len(token) > 0 {
		ss, _ := session.SessionFromToken(token)
		s = ss
	}
	return s
}

func loadAuthentication(ctx iris.Context) {
	s := _sessionFromToken(ctx)
	if s != nil && len(s.Avatar) == 0 {
		s.Avatar = core.CFG.Site.DefaultAvatar
	}
	ctx.Values().Set("session", s)
	ctx.ViewData("session", s)
	ctx.Next()
}

func registerAssets(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	if strings.HasPrefix(r.URL.Path, "/assets") {
		http.FileServer(http.FS(assets.FS)).ServeHTTP(w, r)
		return
	}
	router.ServeHTTP(w, r)
}

var ipRate = cache2go.Cache("ipRateLimiter")

func ipRateLimiter(ctx iris.Context) {

	banTime := time.Minute * 10

	ip := realip.FromRequest(ctx.Request())
	ipRate.NotFoundAdd(ip, banTime, nil)

	item, _ := ipRate.Value(ip)
	if item.AccessCount() > 60*10*5 {
		since := time.Since(item.CreatedOn())
		if since < banTime {
			write_ban_page(banTime-since, ctx)
			return
		}
		ipRate.Delete(ip)
	}
	ctx.Next()
}

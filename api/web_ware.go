package api

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/muesli/cache2go"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
	"github.com/tomasen/realip"
)

// sessionFromCtx
func sessionFromCtx(ctx iris.Context) *session.Session {
	token := ctx.GetCookie("token")
	var s = session.NaS
	if len(token) > 0 {
		ss, _ := session.SessionFromToken(token)
		s = ss
	}
	return s
}

// applyNavStack
func applyNavStack(ctx iris.Context) {
	stack := core.NewNavStack()
	ctx.Values().Set("nav", stack)
	ctx.ViewData("nav", stack)
	ctx.Next()
}

// applySiteProperties
func applySiteProperties(ctx iris.Context) {
	ctx.ViewData("site", core.CFG.Site)
	ctx.Next()
}

// authenticationRequired
func authenticationRequired(ctx iris.Context) {
	s := sessionFromCtx(ctx)
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

// applyAuthentication
func applyAuthentication(ctx iris.Context) {
	s := sessionFromCtx(ctx)
	if s != nil && len(s.Avatar) == 0 {
		s.Avatar = core.CFG.Site.DefaultAvatar
	}
	ctx.Values().Set("session", s)
	ctx.ViewData("session", s)
	ctx.Next()
}

// applyIpRateLimiter
var ipRate = cache2go.Cache("ipRateLimiter")

func applyIpRateLimiter(ctx iris.Context) {

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

// generate_204
func generate_204(ctx iris.Context) {
	ctx.StatusCode(204)
}

// response_404
func response_404(ctx iris.Context) {
	ctx.ViewData("statusCode", iris.StatusNotFound)
	ctx.ViewData("detail", "Not Found")
	ctx.View("failed")
}

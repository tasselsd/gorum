package api

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/russross/blackfriday/v2"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
)

func init() {
	GET["/"] = index
	GET["/r/{rid:string}"] = region
	GET["/d/{did:string}"] = discuss
	GET["/d/{did:string}/p/{page:int}"] = discuss
	GET["/u/{uid:string}"] = user
	GET["/c/{cid:string}"] = comment
}

func index(ctx iris.Context) {

	var discusses []core.Discuss
	ret := core.DB.Find(&discusses)
	if ret.Error != nil {
		write_e500_page(ret.Error, ctx)
		return
	}

	ctx.ViewData("session", ctx.Value("session").(*session.Session))
	ctx.ViewData("recommends", discusses)
	ctx.View("index")
}

func region(ctx iris.Context) {

}

func discuss(ctx iris.Context) {
	did := ctx.Params().GetStringDefault("did", "1")

	var d core.Discuss

	ret := core.DB.Take(&d, "sha1_prefix=?", did)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("展示讨论时遇到一个错误 [ %s ]", ret.Error), ctx)
		return
	}
	b := blackfriday.Run([]byte(d.Content))
	ctx.ViewData("discuss", &d)
	ctx.ViewData("discussHTML", string(b))
	ctx.View("discuss/discuss")
}

func user(ctx iris.Context) {
	uid := ctx.Params().GetStringDefault("uid", "nil")
	var user core.User
	ret := core.DB.Take(&user, "sha1_prefix=? or u_name=?", uid, uid)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("未找到该用户 [%s]", ret.Error), ctx)
		return
	}
	if len(user.Avatar) == 0 {
		user.Avatar = core.CFG.Site.DefaultAvatar
	}

	ctx.ViewData("session", ctx.Value("session").(*session.Session))
	ctx.ViewData("user", &user)
	ctx.View("account/profile")
}

func comment(ctx iris.Context) {

}

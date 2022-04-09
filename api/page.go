package api

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
	"github.com/tasselsd/gorum/templates"
)

func init() {
	GET["/"] = index
	GET["/r/{id:string}"] = region
	GET["/d/{id:string}"] = discuss
	GET["/d/{id:string}/p/{page:int}"] = discuss
	GET["/u/{id:string}"] = user
	GET["/c/{id:string}"] = comment
}

func index(ctx iris.Context) {
	indexData := &templates.IndexPage{
		DefaultPage: templates.DefaultPage{Session: ctx.Value("session").(*session.Session)},
		Recommends: []templates.Recommend{
			{
				DiscussName: "Gorum 正在紧急开发中 ...",
				ShortSha1:   "89705836",
			},
		}}
	templates.WriteHTML(ctx, indexData)
}

func region(ctx iris.Context) {

}

func discuss(ctx iris.Context) {

}

func user(ctx iris.Context) {
	uid := ctx.Params().GetStringDefault("id", "nil")
	var user core.User
	ret := core.DB.Take(&user, "sha1_prefix=? or u_name=?", uid, uid)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("未找到该用户 [%s]", ret.Error), ctx)
		return
	}
	if len(user.Avatar) == 0 {
		user.Avatar = core.CFG.Site.DefaultAvatar
	}
	templates.WriteHTML(ctx, &templates.ProfilePage{
		DefaultPage: templates.DefaultPage{Session: ctx.Value("session").(*session.Session)},
		User:        &user})
}

func comment(ctx iris.Context) {

}

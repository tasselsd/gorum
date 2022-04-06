package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
)

func init() {
	POST["/signup"] = signUp
	POST["/signin"] = signIn
}

func signUp(ctx iris.Context) {
	email := ctx.PostValue("email")
	password := ctx.PostValue("password")
	name := ctx.PostValue("name")

	if len(name) < 2 || len(name) > 32 {
		write_e400_page(errors.New("用户名长度限制在 [ 2-32 ] 个字符"), ctx)
		return
	}

	if len(password) < 6 || len(password) > 128 {
		write_e400_page(errors.New("密码长度限制在 [ 6-128 ] 个字符"), ctx)
		return
	}

	var user core.User
	ret := core.DB.Take(&user, "(email=? or u_name=?)", email, name)
	nameSha1 := core.NewSha1Object(name)
	user.Email = email
	user.Name = name
	user.Passwd = core.NewSha1Object(password).Sha1()
	user.CreateTime = time.Now()
	user.Sha1 = nameSha1.Sha1()
	user.ShortSha1 = nameSha1.ShortSha1()
	if ret.RowsAffected != 0 {
		if user.Valid == 1 {
			write_e400_page(errors.New("邮箱或者用户名已被注册"), ctx)
			return
		}
		core.DB.Save(&user)
		ctx.Redirect("/activation", iris.StatusSeeOther)
		return
	}

	ret = core.DB.Create(&user)
	if ret.RowsAffected != 1 {
		write_e500_page(ret.Error, ctx)
		return
	}
	statusOk(ctx)
}

func signIn(ctx iris.Context) {
	signin := ctx.PostValue("signin")
	passwd := ctx.PostValue("password")
	if len(signin) < 2 {
		write_e400_page(errors.New("账户格式错误"), ctx)
		return
	}
	if len(passwd) < 6 {
		write_e400_page(errors.New("密码格式错误"), ctx)
		return
	}
	var user core.User
	ret := core.DB.Take(&user, "(u_name=? or email=?) and passwd=?", signin, signin, core.NewSha1Object(passwd).Sha1())
	if ret.Error != nil {
		write_e400_page(fmt.Errorf("帐号或密码不正确 [ %s ]", ret.Error.Error()), ctx)
		return
	}
	s := session.NewSession(&user)
	ctx.SetCookieKV("token", s.Token(), iris.CookieExpires(24*time.Hour))
	ctx.SetCookieKV("session", s.JSON(), iris.CookieExpires(24*time.Hour))
	ctx.Redirect("/", iris.StatusSeeOther)
}

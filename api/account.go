package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
	"github.com/tasselsd/gorum/templates"
)

func init() {
	POST["/signup"] = signUp
	POST["/signin"] = signIn
	POST["/reset-password-request"] = requestResetPassword
	POST["/reset-password"] = resetPassword
	GET["/activation/{token:string}"] = doActivation
	GET["/signup"] = signUpPage
	GET["/signin"] = signInPage
	GET["/signout"] = signoutPage
	GET["/reset-password-request"] = requestResetPasswordPage
	GET["/reset-password/{token:string}"] = resetPasswordPage
	GET["/activation"] = activationPage
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
		_redirect2activation(ctx, email, user.OnceToken)
		return
	}
	user.OnceToken = session.NewTokenString()
	ret = core.DB.Create(&user)
	if ret.RowsAffected != 1 {
		write_e500_page(ret.Error, ctx)
		return
	}
	_redirect2activation(ctx, email, user.OnceToken)
}

func _redirect2activation(ctx iris.Context, email, token string) {
	err := core.SendActivation(email, token)
	if err != nil {
		ctx.Application().Logger().Warn(err)
	}
	ctx.Redirect("/activation", iris.StatusSeeOther)
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
	_writeSessionCoookie(ctx, &user)
	ctx.Redirect("/", iris.StatusSeeOther)
}

func signoutPage(ctx iris.Context) {
	if ctx.URLParamExists("y") {
		ctx.RemoveCookie("token")
		ctx.RemoveCookie("session")
		ctx.Redirect("/", iris.StatusTemporaryRedirect)
		return
	}
	templates.WriteHTML(ctx, &templates.SignoutPage{})
}

func _writeSessionCoookie(ctx iris.Context, user *core.User) {
	s := session.NewSession(user)
	ctx.SetCookieKV("token", s.Token(), iris.CookieExpires(24*time.Hour))
	ctx.SetCookieKV("session", s.JSON(), iris.CookieExpires(24*time.Hour))
}

func doActivation(ctx iris.Context) {
	var user core.User
	ret := core.DB.Take(&user, "once_token=?", ctx.Params().GetStringDefault("token", "1"))
	if ret.RowsAffected != 1 {
		write_e400_page(ret.Error, ctx)
		return
	}
	ret = core.DB.Model(&user).Update("valid", 1)
	if ret.RowsAffected != 1 {
		write_e500_page(errors.New("未激活任何账户 [ 是否是已激活状态？ ]"), ctx)
		return
	}
	_writeSessionCoookie(ctx, &user)
	templates.WriteHTML(ctx, &templates.ActivatedPage{})
}

func requestResetPassword(ctx iris.Context) {
	email := ctx.PostValue("email")
	var user core.User
	token := session.NewTokenString()
	if ret := core.DB.Model(&user).Where("email=?", email).Update("once_token", token); ret.RowsAffected != 1 {
		write_e400_page(errors.New("邮箱尚未注册"), ctx)
		return
	}

	if err := core.SendResetPassword(email, token); err != nil {
		write_e500_page(err, ctx)
		return
	}
	templates.WriteHTML(ctx, &templates.SuccessPage{Detail: "重置申请已提交，请从邮箱打开重置密码的链接，以完成密码重置"})
}

func resetPassword(ctx iris.Context) {
	token := ctx.PostValue("token")
	passwd := ctx.PostValue("password")
	var user core.User
	ret := core.DB.Take(&user, "once_token=?", token)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("令牌无效 [%s]", ret.Error), ctx)
		return
	}
	ret = core.DB.Model(&user).Update("passwd", core.NewSha1Object(passwd).Sha1())
	if ret.RowsAffected != 1 {
		write_e500_page(ret.Error, ctx)
		return
	}
	templates.WriteHTML(ctx, &templates.SuccessPage{Detail: "密码更新成功！"})
}

func signUpPage(ctx iris.Context) {
	templates.WriteHTML(ctx, &templates.SignupPage{})
}

func signInPage(ctx iris.Context) {
	templates.WriteHTML(ctx, &templates.SigninPage{})
}

func requestResetPasswordPage(ctx iris.Context) {
	templates.WriteHTML(ctx, &templates.RequestResetPasswordPage{})
}

func resetPasswordPage(ctx iris.Context) {
	token := ctx.Params().GetStringDefault("token", "1")
	var user core.User
	ret := core.DB.Take(&user, "once_token=?", token)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("令牌无效 [%s]", ret.Error), ctx)
		return
	}
	templates.WriteHTML(ctx, &templates.ResetPasswordPage{Token: token, User: &user})
}

func activationPage(ctx iris.Context) {
	templates.WriteHTML(ctx, &templates.ActivationPage{})
}

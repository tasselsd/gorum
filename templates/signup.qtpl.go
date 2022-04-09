// Code generated by qtc from "signup.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/signup.qtpl:1
package templates

//line templates/signup.qtpl:1
import "github.com/tasselsd/gorum/pkg/core"

//line templates/signup.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/signup.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/signup.qtpl:5
type SignupPage struct {
	DefaultPage
}
type SigninPage struct {
	DefaultPage
}
type ActivationPage struct {
	DefaultPage
}
type ActivatedPage struct {
	DefaultPage
}
type RequestResetPasswordPage struct {
	DefaultPage
}
type ResetPasswordPage struct {
	DefaultPage
	User  *core.User
	Token string
}
type SignoutPage struct {
	DefaultPage
}

//line templates/signup.qtpl:30
func (p *SignupPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:30
	qw422016.N().S(`<div class="sign-area"><form method="post" action="/signup"><input type="email" name="email" /><input type="text" name="name" /><input type="password" name="password" /><input type="submit" value="注册" /></form><div class="sign-links"><ul><li><a href="/signin">登录</a></li></ul></div></div>`)
//line templates/signup.qtpl:44
}

//line templates/signup.qtpl:44
func (p *SignupPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:44
	p.StreamBody(qw422016)
//line templates/signup.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:44
}

//line templates/signup.qtpl:44
func (p *SignupPage) Body() string {
//line templates/signup.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:44
	p.WriteBody(qb422016)
//line templates/signup.qtpl:44
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:44
	return qs422016
//line templates/signup.qtpl:44
}

//line templates/signup.qtpl:45
func (p *SignupPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:45
	qw422016.N().S(`注册`)
//line templates/signup.qtpl:47
}

//line templates/signup.qtpl:47
func (p *SignupPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:47
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:47
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:47
}

//line templates/signup.qtpl:47
func (p *SignupPage) Title() string {
//line templates/signup.qtpl:47
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:47
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:47
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:47
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:47
	return qs422016
//line templates/signup.qtpl:47
}

//line templates/signup.qtpl:49
func (p *SigninPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:49
	qw422016.N().S(`<div class="sign-area"><form method="post" action="/signin"><input type="text" name="signin" /><input type="password" name="password" /><input type="submit" value="登录" /></form><div class="sign-links"><ul><li><a href="/signup">注册帐号</a></li><li><a href="/reset-password-request">重置密码</a></li></ul></div></div>`)
//line templates/signup.qtpl:63
}

//line templates/signup.qtpl:63
func (p *SigninPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:63
	p.StreamBody(qw422016)
//line templates/signup.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:63
}

//line templates/signup.qtpl:63
func (p *SigninPage) Body() string {
//line templates/signup.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:63
	p.WriteBody(qb422016)
//line templates/signup.qtpl:63
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:63
	return qs422016
//line templates/signup.qtpl:63
}

//line templates/signup.qtpl:64
func (p *SigninPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:64
	qw422016.N().S(`登录`)
//line templates/signup.qtpl:66
}

//line templates/signup.qtpl:66
func (p *SigninPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:66
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:66
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:66
}

//line templates/signup.qtpl:66
func (p *SigninPage) Title() string {
//line templates/signup.qtpl:66
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:66
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:66
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:66
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:66
	return qs422016
//line templates/signup.qtpl:66
}

//line templates/signup.qtpl:68
func (p *RequestResetPasswordPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:68
	qw422016.N().S(`<div class="sign-area"><form method="post" action="/reset-password-request"><input type="email" name="email" /><input type="submit" value="提交申请" /></form><div class="sign-links"><ul><li><a href="/signup">注册帐号</a></li><li><a href="/signin">登录</a></li></ul></div></div>`)
//line templates/signup.qtpl:81
}

//line templates/signup.qtpl:81
func (p *RequestResetPasswordPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:81
	p.StreamBody(qw422016)
//line templates/signup.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:81
}

//line templates/signup.qtpl:81
func (p *RequestResetPasswordPage) Body() string {
//line templates/signup.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:81
	p.WriteBody(qb422016)
//line templates/signup.qtpl:81
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:81
	return qs422016
//line templates/signup.qtpl:81
}

//line templates/signup.qtpl:82
func (p *RequestResetPasswordPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:82
	qw422016.N().S(`申请重置密码`)
//line templates/signup.qtpl:84
}

//line templates/signup.qtpl:84
func (p *RequestResetPasswordPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:84
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:84
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:84
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:84
}

//line templates/signup.qtpl:84
func (p *RequestResetPasswordPage) Title() string {
//line templates/signup.qtpl:84
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:84
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:84
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:84
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:84
	return qs422016
//line templates/signup.qtpl:84
}

//line templates/signup.qtpl:86
func (p *ResetPasswordPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:86
	qw422016.N().S(`<div class="sign-area"><div class="current-user">`)
//line templates/signup.qtpl:88
	qw422016.E().S(p.User.Name)
//line templates/signup.qtpl:88
	qw422016.N().S(`</div><form method="post" action="/reset-password"><input type="password" name="password" /><input type="password" name="token" value="`)
//line templates/signup.qtpl:91
	qw422016.E().S(p.Token)
//line templates/signup.qtpl:91
	qw422016.N().S(`" style="display: none" /><input type="submit" value="重置密码" /></form><div class="sign-links"><ul><li><a href="/signup">注册帐号</a></li><li><a href="/signin">登录</a></li></ul></div></div>`)
//line templates/signup.qtpl:101
}

//line templates/signup.qtpl:101
func (p *ResetPasswordPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:101
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:101
	p.StreamBody(qw422016)
//line templates/signup.qtpl:101
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:101
}

//line templates/signup.qtpl:101
func (p *ResetPasswordPage) Body() string {
//line templates/signup.qtpl:101
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:101
	p.WriteBody(qb422016)
//line templates/signup.qtpl:101
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:101
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:101
	return qs422016
//line templates/signup.qtpl:101
}

//line templates/signup.qtpl:102
func (p *ResetPasswordPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:102
	qw422016.N().S(`重置密码`)
//line templates/signup.qtpl:104
}

//line templates/signup.qtpl:104
func (p *ResetPasswordPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:104
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:104
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:104
}

//line templates/signup.qtpl:104
func (p *ResetPasswordPage) Title() string {
//line templates/signup.qtpl:104
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:104
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:104
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:104
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:104
	return qs422016
//line templates/signup.qtpl:104
}

//line templates/signup.qtpl:106
func (p *ActivationPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:106
	qw422016.N().S(`<div class="activation-tips">你的帐号已登记注册！但此刻还未激活！<br /><br />注册时填写邮箱的收件箱似乎已经收到了一封激活邮件，请点击链接激活它</div><div class="error-nav"><ul><li><a href="javascript:history.go(-1);">返回上一步</a></li></ul></div>`)
//line templates/signup.qtpl:115
}

//line templates/signup.qtpl:115
func (p *ActivationPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:115
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:115
	p.StreamBody(qw422016)
//line templates/signup.qtpl:115
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:115
}

//line templates/signup.qtpl:115
func (p *ActivationPage) Body() string {
//line templates/signup.qtpl:115
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:115
	p.WriteBody(qb422016)
//line templates/signup.qtpl:115
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:115
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:115
	return qs422016
//line templates/signup.qtpl:115
}

//line templates/signup.qtpl:116
func (p *ActivationPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:116
	qw422016.N().S(`激活帐号`)
//line templates/signup.qtpl:118
}

//line templates/signup.qtpl:118
func (p *ActivationPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:118
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:118
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:118
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:118
}

//line templates/signup.qtpl:118
func (p *ActivationPage) Title() string {
//line templates/signup.qtpl:118
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:118
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:118
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:118
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:118
	return qs422016
//line templates/signup.qtpl:118
}

//line templates/signup.qtpl:120
func (p *ActivatedPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:120
	qw422016.N().S(`<div class="activated-tips">激活成功，正在<a href="/">跳转首页</a>...</div><script>setTimeout(()=>window.location.href="/", 1000)</script>`)
//line templates/signup.qtpl:127
}

//line templates/signup.qtpl:127
func (p *ActivatedPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:127
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:127
	p.StreamBody(qw422016)
//line templates/signup.qtpl:127
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:127
}

//line templates/signup.qtpl:127
func (p *ActivatedPage) Body() string {
//line templates/signup.qtpl:127
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:127
	p.WriteBody(qb422016)
//line templates/signup.qtpl:127
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:127
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:127
	return qs422016
//line templates/signup.qtpl:127
}

//line templates/signup.qtpl:128
func (p *ActivatedPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:128
	qw422016.N().S(`激活成功`)
//line templates/signup.qtpl:130
}

//line templates/signup.qtpl:130
func (p *ActivatedPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:130
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:130
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:130
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:130
}

//line templates/signup.qtpl:130
func (p *ActivatedPage) Title() string {
//line templates/signup.qtpl:130
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:130
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:130
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:130
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:130
	return qs422016
//line templates/signup.qtpl:130
}

//line templates/signup.qtpl:132
func (p *SignoutPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:132
	qw422016.N().S(`<div class="signout-tips">确定退出当前会话吗？<a href="/signout?y">是的，立即退出！</a></div><div class="error-nav"><ul><li><a href="javascript:history.go(-1);">返回上一步</a></li></ul></div>`)
//line templates/signup.qtpl:141
}

//line templates/signup.qtpl:141
func (p *SignoutPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:141
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:141
	p.StreamBody(qw422016)
//line templates/signup.qtpl:141
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:141
}

//line templates/signup.qtpl:141
func (p *SignoutPage) Body() string {
//line templates/signup.qtpl:141
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:141
	p.WriteBody(qb422016)
//line templates/signup.qtpl:141
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:141
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:141
	return qs422016
//line templates/signup.qtpl:141
}

//line templates/signup.qtpl:142
func (p *SignoutPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/signup.qtpl:142
	qw422016.N().S(`退出登录`)
//line templates/signup.qtpl:144
}

//line templates/signup.qtpl:144
func (p *SignoutPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/signup.qtpl:144
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/signup.qtpl:144
	p.StreamTitle(qw422016)
//line templates/signup.qtpl:144
	qt422016.ReleaseWriter(qw422016)
//line templates/signup.qtpl:144
}

//line templates/signup.qtpl:144
func (p *SignoutPage) Title() string {
//line templates/signup.qtpl:144
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/signup.qtpl:144
	p.WriteTitle(qb422016)
//line templates/signup.qtpl:144
	qs422016 := string(qb422016.B)
//line templates/signup.qtpl:144
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/signup.qtpl:144
	return qs422016
//line templates/signup.qtpl:144
}

package api

import (
	"errors"
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
	GET["/d/{did:string}/c/{cid:string}"] = comment
	GET["/u/{uid:string}"] = user
}

type Recommend struct {
	core.Discuss
	Initiator *core.User
}

type Comment struct {
	core.Comment
	Initiator   *core.User
	CommentHTML string
}

func index(ctx iris.Context) {

	var discusses []core.Discuss
	ret := core.DB.Find(&discusses)
	if ret.Error != nil {
		write_e500_page(ret.Error, ctx)
		return
	}

	uc := core.NewUserCache()
	var recommends []Recommend
	for _, d := range discusses {
		u := uc.Get(d.InitiatorUid)
		recommends = append(recommends, Recommend{
			Discuss:   d,
			Initiator: u,
		})
	}
	ctx.ViewData("session", ctx.Value("session").(*session.Session))
	ctx.ViewData("recommends", recommends)
	ctx.View("index")
}

func region(ctx iris.Context) {

}

func discuss(ctx iris.Context) {
	did := ctx.Params().GetStringDefault("did", "nil")
	page := ctx.Params().GetInt64Default("p", 1)

	if page < 0 {
		write_e400_page(errors.New("页码错误"), ctx)
		return
	}

	var d core.Discuss

	ret := core.DB.Take(&d, "sha1_prefix=?", did)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("展示讨论时遇到一个错误 [ %s ]", ret.Error), ctx)
		return
	}

	var comments []core.Comment

	ret = core.DB.Find(&comments).Where("discuss_did=? limit ?,? order by create_time desc", d.ID, (page-1)*20, 20)
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("服务器出现了一个错误 [%s]", ret.Error.Error()), ctx)
		return
	}

	var commentsN []Comment

	uc := core.NewUserCache()
	for _, c := range comments {
		u := uc.Get(c.InitiatorUid)
		b := blackfriday.Run([]byte(c.Content))
		commentsN = append(commentsN, Comment{
			Comment:     c,
			Initiator:   u,
			CommentHTML: string(b),
		})
	}
	b := blackfriday.Run([]byte(d.Content))
	ctx.ViewData("discuss", &d)
	ctx.ViewData("discussHTML", string(b))
	ctx.ViewData("comments", commentsN)
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

	var discusses []core.Discuss

	ret = core.DB.Find(&discusses).Where("initiator_uid=? order by create_time desc", user.ID)
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("服务器发生了一个错误[%s]", ret.Error.Error()), ctx)
		return
	}

	ctx.ViewData("session", ctx.Value("session").(*session.Session))
	ctx.ViewData("user", &user)
	ctx.ViewData("discusses", discusses)
	ctx.View("account/profile")
}

func comment(ctx iris.Context) {
	cid := ctx.Params().GetStringDefault("cid", "nil")
	did := ctx.Params().GetStringDefault("did", "nil")
	var comment core.Comment
	ret := core.DB.Take(&comment, "sha1_prefix=?", cid)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("未找到该评论 [%s]", ret.Error), ctx)
		return
	}
	var afteMe int64
	ret = core.DB.Model(&core.Comment{}).
		Where("discuss_did=? and create_time>?", comment.DiscussDid, comment.CreateTime).
		Count(&afteMe)
	if ret.RowsAffected != 1 {
		write_e500_page(fmt.Errorf("服务器发生了一个错误 [%s]", ret.Error), ctx)
		return
	}

	page := afteMe/20 + 1
	ctx.Redirect(fmt.Sprintf("/d/%s/p/%d", did, page))
}

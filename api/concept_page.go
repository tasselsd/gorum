package api

import (
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/microcosm-cc/bluemonday"
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

type RegionDiscuss struct {
	core.Discuss
	Initiator *core.User
}

type ProfileComment struct {
	core.Comment
	CommentHTML string
	DiscussName string
}

type Comment struct {
	core.Comment
	Initiator   *core.User
	CommentHTML string
}

func index(ctx iris.Context) {

	var discusses []core.Discuss
	ret := core.DB.Limit(50).Order("create_time desc").Find(&discusses)
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
	rid := ctx.Params().GetStringDefault("rid", "1")
	var region core.Region
	ret := core.DB.Take(&region, "sha1_prefix=? or r_name=? or id=?", rid, rid, rid)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("不存在的辖区 [%s]", ret.Error), ctx)
		return
	}

	var discusses []core.Discuss
	ret = core.DB.Find(&discusses, "division_rid=?", region.ID)
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("发现一个错误 [%s]", ret.Error.Error()), ctx)
		return
	}

	uc := core.NewUserCache()
	var ds []RegionDiscuss
	for _, d := range discusses {
		ds = append(ds, RegionDiscuss{
			Discuss:   d,
			Initiator: uc.Get(d.InitiatorUid),
		})
	}

	ctx.ViewData("discusses", ds)
	ctx.ViewData("region", region)
	ctx.View("discuss/region")
}

func discuss(ctx iris.Context) {
	did := ctx.Params().GetStringDefault("did", "nil")
	page := ctx.Params().GetInt64Default("p", 1)

	if page < 0 {
		write_e400_page(errors.New("页码错误"), ctx)
		return
	}

	var d core.Discuss

	ret := core.DB.Take(&d, "sha1_prefix=? or id=?", did, did)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("展示讨论时遇到一个错误 [ %s ]", ret.Error), ctx)
		return
	}

	var comments []core.Comment

	ret = core.DB.Order("create_time desc").Limit(20).Offset(int((page-1)*20)).Find(&comments, "discuss_did=?", d.ID)
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("服务器出现了一个错误 [%s]", ret.Error.Error()), ctx)
		return
	}

	var commentsN []Comment

	uc := core.NewUserCache()
	for _, c := range comments {
		u := uc.Get(c.InitiatorUid)
		b := blackfriday.Run([]byte(c.Content), blackfriday.WithExtensions(blackfriday.HardLineBreak))
		commentsN = append(commentsN, Comment{
			Comment:     c,
			Initiator:   u,
			CommentHTML: string(bluemonday.UGCPolicy().SanitizeBytes(b)),
		})
	}
	b := blackfriday.Run([]byte(d.Content), blackfriday.WithExtensions(blackfriday.HardLineBreak))

	ctx.Value("nav").(*core.NavStack).Push(d.Division, fmt.Sprintf("/r/%d", d.DivisionRid))
	ctx.ViewData("discuss", &d)
	ctx.ViewData("discussHTML", string(bluemonday.UGCPolicy().SanitizeBytes(b)))
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

	ret = core.DB.Order("create_time desc").Find(&discusses, "initiator_uid=?", user.ID)
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("发生一个错误 [%s]", ret.Error.Error()), ctx)
		return
	}

	var comments []core.Comment
	ret = core.DB.Find(&comments, "initiator_uid=?", user.ID).Order("create_time desc")
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("发生一个错误 [%s]", ret.Error.Error()), ctx)
		return
	}
	var commentsN []ProfileComment

	for _, c := range comments {
		b := blackfriday.Run([]byte(c.Content), blackfriday.WithExtensions(blackfriday.HardLineBreak))
		commentsN = append(commentsN, ProfileComment{
			Comment:     c,
			CommentHTML: string(bluemonday.UGCPolicy().SanitizeBytes(b)),
			DiscussName: c.Discuss,
		})
	}
	ctx.ViewData("session", ctx.Value("session").(*session.Session))
	ctx.ViewData("user", &user)
	ctx.ViewData("discusses", discusses)
	ctx.ViewData("comments", commentsN)
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

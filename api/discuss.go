package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/tasselsd/gorum/pkg/core"
	"github.com/tasselsd/gorum/pkg/session"
)

func init() {
	G["/r/{rid:string}/discuss-form"] = postDiscussPage
	G["/region-selector"] = regionSelectorPage
	POST["/r/{rid:string}/discuss"] = postDiscuss
	PO["/d/{did:string}/comment"] = postComment
}

func regionSelectorPage(ctx iris.Context) {
	var regions []core.Region

	ret := core.DB.Find(&regions)
	if ret.Error != nil {
		write_e500_page(fmt.Errorf("遇到一个错误 [%s]", ret.Error.Error()), ctx)
		return
	}
	ctx.ViewData("regions", regions)
	ctx.View("discuss/region-selector")
}

func postDiscussPage(ctx iris.Context) {
	rid := ctx.Params().GetStringDefault("rid", "1")
	if len(rid) == 0 {
		write_e400_page(errors.New("请指定知识辖区"), ctx)
		return
	}
	r, err := _regionByRid(rid)
	if err != nil {
		write_e400_page(err, ctx)
		return
	}
	ctx.Value("nav").(*core.NavStack).Push("辖区选择", "/region-selector")
	ctx.ViewData("division", r)
	ctx.View("discuss/post-discuss")
}

func _regionByRid(rid string) (*core.Region, error) {
	var region core.Region

	ret := core.DB.Take(&region, "sha1_prefix=? or r_name=?", rid, rid)
	if ret.RowsAffected != 1 {
		return nil, fmt.Errorf("未找到该知识辖区 [%s]", ret.Error)
	}
	return &region, nil
}

func postDiscuss(ctx iris.Context) {
	rid := ctx.Params().GetStringDefault("rid", "1")
	title := ctx.PostValue("title")
	content := ctx.PostValue("content")
	if len(title) < 5 || len(title) > 128 {
		write_e400_page(errors.New("标题长度被限制在 [ 4-128 ] 个字符"), ctx)
		return
	}

	if len(content) < 10 || len(content) > 1024*1024 {
		write_e400_page(errors.New("内容被限制在 [ 10-1M ] 个字符"), ctx)
		return
	}
	b := blackfriday.Run([]byte(content), blackfriday.WithExtensions(blackfriday.HardLineBreak))
	finalByets := bluemonday.UGCPolicy().SanitizeBytes(b)
	if len(finalByets) < 10 {
		write_e400_page(errors.New("输入些可见字符吧！"), ctx)
		return
	}

	s := ctx.Values().Get("session").(*session.Session)

	r, err := _regionByRid(rid)
	if err != nil {
		write_e400_page(err, ctx)
		return
	}

	sha1 := core.NewSha1Object(content)

	var d core.Discuss
	d.Content = content
	d.Name = title
	d.CreateTime = time.Now()
	d.Division = r.Name
	d.DivisionRid = r.ID
	d.Initiator = s.Name
	d.InitiatorUid = s.ID
	d.Sha1 = sha1.Sha1()
	d.ShortSha1 = sha1.ShortSha1()
	ret := core.DB.Create(&d)
	if ret.RowsAffected != 1 {
		write_e500_page(fmt.Errorf("没有任何内容被存储 [ %s ]", ret.Error), ctx)
		return
	}
	ctx.Redirect("/d/"+d.ShortSha1, iris.StatusSeeOther)
}

func postComment(ctx iris.Context) {
	did := ctx.Params().GetStringDefault("did", "nil")
	content := ctx.PostValue("comment")
	if len(content) < 5 || len(content) > 1024 {
		write_e400_page(errors.New("内容长度限制在 [ 5-1M ] 个字符"), ctx)
		return
	}

	var discuss core.Discuss
	ret := core.DB.Take(&discuss, "sha1_prefix=?", did)
	if ret.RowsAffected != 1 {
		write_e400_page(fmt.Errorf("不存在的讨论 [%s]", ret.Error), ctx)
		return
	}

	s := ctx.Value("session").(*session.Session)
	sha1 := core.NewSha1Object(content)
	comment := core.Comment{
		Content:      content,
		InitiatorUid: s.ID,
		Initiator:    s.Name,
		DiscussDid:   discuss.ID,
		CreateTime:   time.Now(),
		Sha1:         sha1.Sha1(),
		ShortSha1:    sha1.ShortSha1(),
	}
	ret = core.DB.Create(&comment)
	if ret.RowsAffected != 1 {
		write_e500_page(fmt.Errorf("没有任何内容被存储 [%s]", ret.Error), ctx)
		return
	}
	ctx.Redirect(fmt.Sprintf("/d/%s/c/%s", did, comment.ShortSha1), iris.StatusSeeOther)
}

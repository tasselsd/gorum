package api

import (
	"github.com/kataras/iris/v12"
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
	indexData := &templates.IndexPage{Recommends: []templates.Recommend{
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

}

func comment(ctx iris.Context) {

}

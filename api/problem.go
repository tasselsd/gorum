package api

import (
	"github.com/kataras/iris/v12"
)

func write_e400_page(err error, ctx iris.Context) {
	ctx.ViewData("statusCode", iris.StatusBadRequest)
	ctx.ViewData("detail", err.Error())
	ctx.View("failed")
}

func write_e500_page(err error, ctx iris.Context) {
	ctx.ViewData("statusCode", iris.StatusInternalServerError)
	ctx.ViewData("detail", err.Error())
	ctx.View("failed")
}

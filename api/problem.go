package api

import (
	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/templates"
)

func write_e400_page(err error, ctx iris.Context) {
	templates.WriteHTML(ctx, &templates.ErrorPage{
		StatusCode: iris.StatusBadRequest,
		Detail:     err.Error(),
	})
}
func write_e500_page(err error, ctx iris.Context) {
	templates.WriteHTML(ctx, &templates.ErrorPage{
		StatusCode: iris.StatusInternalServerError,
		Detail:     err.Error(),
	})
}

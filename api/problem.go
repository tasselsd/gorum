package api

import (
	"github.com/kataras/iris/v12"
	"github.com/tasselsd/gorum/templates"
)

func p400_wrap(err error) iris.Problem {
	return iris.NewProblem().DetailErr(err).Status(iris.StatusBadRequest)
}

func p400(err string) iris.Problem {
	return iris.NewProblem().Detail(err).Status(iris.StatusBadRequest)
}

func p500_wrap(err error) iris.Problem {
	return iris.NewProblem().DetailErr(err).Status(iris.StatusInternalServerError)
}

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

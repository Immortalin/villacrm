package routes

import (
	"github.com/kataras/iris"
	"github.com/speedwheel/villacrm/models"
	"github.com/speedwheel/villacrm/repositories"
	"github.com/speedwheel/villacrm/services"
)

var (
	PathClientList = "client/list"
)

// GetClientsHandler handles the GET: /clients
func GetClientsHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View(PathClientList + ".html")
}

func PostClientHandler(ctx iris.Context) {
	if !ctx.IsAjax() {
		ctx.StatusCode(iris.StatusNotFound)
	}
	var msg string
	var (
		name     = ctx.FormValue("name")
		email    = ctx.FormValue("email")
		villas   = ctx.FormValue("villas")
		dates    = ctx.FormValue("dates")
		status   = ctx.FormValue("status")
		currency = ctx.FormValue("currency")
		price    = ctx.FormValue("price")
		source   = ctx.FormValue("source")
	)

	_, err := services.CreateUser(villas, dates, price, status, models.Client{
		Name:     name,
		Email:    email,
		Currency: currency,
		Source:   source,
	})

	if err != nil {
		ctx.Application().Logger().Error(err.Error())
		msg = err.Error()
	}

	if err != nil {
		ctx.Application().Logger().Error(err.Error())
	}

	ctx.JSON(map[string]interface{}{"error": msg})
}

func TableClientsHandler(ctx iris.Context) {
	if !ctx.IsAjax() {
		ctx.StatusCode(iris.StatusNotFound)
	}
	urlQuery := ctx.Request().URL.Query()
	draw := urlQuery["draw"][0]
	Data, CountFiltered, Count := repositories.GetClientTable(urlQuery)
	ctx.JSON(map[string]interface{}{"draw": draw, "recordsTotal": Count, "recordsFiltered": CountFiltered, "data": Data})

}

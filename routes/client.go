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

func GetClientBookedHandler(ctx iris.Context) {
	if !ctx.IsAjax() {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	villa := ctx.FormValue("villa")
	booked := repositories.SelectVillaBookedDates(villa)
	ctx.JSON(map[string]interface{}{"booked": booked})
}

func PostClientHandler(ctx iris.Context) {
	if !ctx.IsAjax() {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	var msg string
	var (
		name     = ctx.FormValue("name")
		email    = ctx.FormValue("email")
		villa    = ctx.FormValue("villa")
		dates    = ctx.FormValue("dates")
		status   = ctx.FormValue("status")
		currency = ctx.FormValue("currency")
		price    = ctx.FormValue("price")
		referral = ctx.FormValue("referral")
	)

	_, err := services.CreateUser(dates, price, status, models.Client{
		Name:     name,
		Email:    email,
		Villa:    villa,
		Currency: currency,
		Referral: referral,
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
		return
	}
	urlQuery := ctx.Request().URL.Query()
	draw := urlQuery["draw"][0]
	Data, CountFiltered, Count := repositories.GetClientTable(urlQuery)
	ctx.JSON(map[string]interface{}{"draw": draw, "recordsTotal": Count, "recordsFiltered": CountFiltered, "data": Data})

}

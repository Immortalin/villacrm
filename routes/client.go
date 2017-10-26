package routes

import (
	"github.com/kataras/iris"
	"github.com/speedwheel/villacrm/models"
	"github.com/speedwheel/villacrm/repositories"
	"gopkg.in/mgo.v2/bson"
)

var (
	PathClientList = "client/list"
)

// GetClientsHandler handles the GET: /clients
func GetClientsHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View(PathClientList + ".html")
}

func AddGetClientsHandler(ctx iris.Context) {
	if !ctx.IsAjax() {
		ctx.StatusCode(iris.StatusNotFound)
	}
	client := models.Client{}
	err := ctx.ReadForm(&client)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Error(err.Error())
	}
	client.ID = bson.NewObjectId()
	r := repositories.Insert(client)
	if r == false {
		ctx.StatusCode(iris.StatusInternalServerError)
	}
	ctx.JSON(map[string]bool{"status": r})
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

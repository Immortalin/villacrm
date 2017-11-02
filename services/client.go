package services

import (
	"strconv"
	"strings"

	"github.com/kataras/golog"
	"github.com/speedwheel/villacrm/models"
	"github.com/speedwheel/villacrm/repositories"
)

func CreateUser(villas string, dates string, price string, status string, client models.Client) (models.Client, error) {
	var error error
	d1, d2, days := models.CalculateBookingDates(dates)
	villasSlice := strings.Split(villas, " ")
	statusInt := 0
	if status == "on" {
		statusInt = 1
	}
	priceFloat64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		golog.Error(err.Error())
	}

	client.In = d1
	client.Out = d2
	client.Days = days
	client.Villas = villasSlice
	client.Status = statusInt
	client.Price = priceFloat64

	return repositories.CreateUser(client)

	return client, error
}

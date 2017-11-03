package services

import (
	"errors"
	"strconv"

	"github.com/kataras/golog"
	"github.com/speedwheel/villacrm/models"
	"github.com/speedwheel/villacrm/repositories"
)

func CreateUser(dates string, price string, status string, client models.Client) (models.Client, error) {

	if client.Name == "" || client.Email == "" || client.Villa == "" || status == "" || client.Currency == "" || price == "" || client.Referral == "" {
		return client, errors.New("Please fill in all the fields")
	}
	if dates != "" {
		d1, d2, days := models.CalculateBookingDates(dates)
		client.In = d1
		client.Out = d2
		client.Days = days
	}
	statusInt := 0
	if status == "on" {
		statusInt = 1
	}
	priceFloat64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		golog.Error(err.Error())
	}

	client.Status = statusInt
	client.Price = priceFloat64

	return repositories.CreateUser(client)
}

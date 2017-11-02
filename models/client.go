package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Client this is the client structure
type Client struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name" form:"name"`
	Email    string        `json:"email" bson:"email" form:"email"`
	Villas   []string      `json:"villas" bson:"villas" form:"villas"`
	In       time.Time     `json:"in" bson:"in" form:"-"`
	Out      time.Time     `json:"out" bson:"out" form:"-"`
	Days     int           `json:"days" bson:"days" form:"-"`
	Status   int           `json:"status" bson:"status" form:"status"`
	Currency string        `json:"currency" bson:"currency" form:"currency"`
	Price    float64       `json:"price" bson:"price" form:"price"`
	Source   string        `json:"source" bson:"source" form:"source"`
}

func CalculateBookingDates(dates string) (time.Time, time.Time, int) {
	d := strings.Split(dates, " - ")

	d1, _ := time.Parse("2006-01-02", d[0])
	d2, _ := time.Parse("2006-01-02", d[1])
	diff := d2.Sub(d1)
	days := int(diff.Hours() / 24)

	return d1, d2, days
}

package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	No      string        `json:"no" bson:"no" form:"no"`
	Name    string        `json:"name" bson:"name" form:"name"`
	Villa   string        `json:"villa" bson:"villa" form:"villa"`
	In      time.Time     `json:"in" bson:"in" form:"in"`
	Out     time.Time     `json:"out" bson:"out" form:"out"`
	Days    int           `json:"days" bson:"days" form:"days"`
	Persons int           `json:"persons" bson:"persons" form:"persons"`
}

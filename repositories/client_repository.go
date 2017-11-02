package repositories

import (
	"errors"
	"log"
	"net/url"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/golog"
	"github.com/speedwheel/villacrm/datasource"
	"github.com/speedwheel/villacrm/models"
)

var (
	ClientCollection = "clients"
)

func CheckIfUserExists(email string) bool {
	found := true
	client := models.Client{}
	Db := datasource.MgoDb{}
	Db.Init()
	c := Db.C(ClientCollection)
	if err := c.Find(bson.M{"email": email}).One(&client); err != nil {
		found = false
	}
	Db.Close()

	return found
}

// CreateUser inserts a new user in the db
func CreateUser(client models.Client) (models.Client, error) {

	if CheckIfUserExists(client.Email) {
		return models.Client{}, errors.New("User already exists")
	}
	var error error
	Db := datasource.MgoDb{}
	client.ID = bson.NewObjectId()
	Db.Init()
	c := Db.C(ClientCollection)
	if err := c.Insert(&client); err != nil {
		golog.Error(err.Error())
		error = err
	}
	Db.Close()
	return client, error
}

func GetClientTable(urlQuery url.Values) ([]models.Client, int, int) {
	clients := []models.Client{}
	query := bson.M{}
	limit := 0
	limit, _ = strconv.Atoi(urlQuery["length"][0])
	skips, _ := strconv.Atoi(urlQuery["start"][0])
	searchValue := urlQuery["search[value]"][0]
	if searchValue != "" {
		query["name"] = bson.M{"$regex": "^" + searchValue}
	}

	pm := bson.M{
		"$match": query,
	}

	pl := bson.M{
		"$limit": limit,
	}

	ps := bson.M{
		"$skip": skips,
	}
	po := bson.M{

		"$sort": bson.D{
			bson.DocElem{Name: "name", Value: 1},
		},
	}
	Db := datasource.MgoDb{}
	Db.Init()
	c := Db.C(ClientCollection)

	pipe := c.Pipe([]bson.M{pm, po, ps, pl})
	if err := pipe.All(&clients); err != nil {
		log.Printf(err.Error())
	}
	CountFiltered, err := c.Find(query).Count()
	if err != nil {
		panic(err)
	}
	Count, err := c.Find(nil).Count()
	if err != nil {
		panic(err)
	}
	Db.Close()
	return clients, CountFiltered, Count
}

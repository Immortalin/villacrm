package datasource

import (
	"github.com/kataras/golog"
	"gopkg.in/mgo.v2"
)

var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
	dbName      = "villacrm"
)

type MgoDb struct {
	Session *mgo.Session
	Db      *mgo.Database
	Col     *mgo.Collection
}

func init() {
	if mainSession == nil {
		var err error
		mainSession, err = mgo.Dial("localhost:27017")

		if err != nil {
			golog.Error(err)
		}

		mainSession.SetMode(mgo.Monotonic, true)
		mainDb = mainSession.DB(dbName)
	}

}

func (this *MgoDb) Init() *mgo.Session {

	this.Session = mainSession.Copy()
	this.Db = this.Session.DB(dbName)

	/*searchAuth := mgo.Index{
		Key: []string{"slug", "like", "pro", "check"},
	  }
	this.C("businesses").EnsureIndex(searchAuth)*/

	return this.Session
}

func (this *MgoDb) C(collection string) *mgo.Collection {
	this.Col = this.Session.DB(dbName).C(collection)
	return this.Col
}

func (this *MgoDb) Close() bool {
	defer this.Session.Close()
	return true
}

func (this *MgoDb) DropoDb() {
	err := this.Session.DB(dbName).DropDatabase()
	if err != nil {
		golog.Error(err)
	}
}

func (this *MgoDb) RemoveAll(collection string) bool {
	this.Session.DB(dbName).C(collection).RemoveAll(nil)

	this.Col = this.Session.DB(dbName).C(collection)
	return true
}

func (this *MgoDb) Index(collection string, keys []string) bool {

	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := this.Db.C(collection).EnsureIndex(index)
	if err != nil {
		golog.Error(err)
		return false
	}

	return true
}

func (this *MgoDb) IsDup(err error) bool {

	if mgo.IsDup(err) {
		return true
	}

	return false
}

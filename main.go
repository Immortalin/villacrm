package main

import (
	"github.com/speedwheel/villacrm/bootstrap"
	"github.com/speedwheel/villacrm/middleware/identity"
	"github.com/speedwheel/villacrm/routes"
)

var app = bootstrap.New("VillaCRM", "edi.ultras@gmai.com",
	identity.Configure,
	routes.Configure,
)

func init() {
	app.Bootstrap()
}

func main() {
	app.Listen(":8080")
}

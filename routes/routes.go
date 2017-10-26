package routes

import (
	"github.com/speedwheel/villacrm/bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", GetIndexHandler)

	client := b.Party("/client")
	client.Get("/list", GetClientsHandler)
	client.Post("/add", AddGetClientsHandler)
	client.Get("/table", TableClientsHandler)
}

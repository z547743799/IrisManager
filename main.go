package main

import (
	"gitlab.com/z547743799/irismanager/bootstrap"
	"gitlab.com/z547743799/irismanager/webapp/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "極")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}

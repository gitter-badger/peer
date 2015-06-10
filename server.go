package main

import "github.com/go-martini/martini"

func main() {
	app := martini.Classic()

	app.Get("/", func() string {
		return "Hello World"
	})

	app.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	app.Run()
}

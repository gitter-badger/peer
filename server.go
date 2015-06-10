package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/gzip"

func main() {
	app := martini.Classic()
	app.Use(gzip.All())

	app.Get("/", func() string {
		return "Hello World"
	})

	app.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	app.Run()
}

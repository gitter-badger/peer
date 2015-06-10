package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/gzip"
import "github.com/martini-contrib/render"

func main() {
	app := martini.Classic()
	app.Use(gzip.All())
	app.Use(render.Renderer())

	app.Get("/", func() string {
		return "Hello World"
	})

	app.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	app.Get("/json", func(render render.Render) {
		render.JSON(200, map[string]interface{}{"hello": "World"})
	})

	app.Run()
}

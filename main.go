package main

import (
	"github.com/PhotoTresor/peer/app/models"
	"github.com/jinzhu/gorm"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/martini-contrib/binding"
	"os"
	"io"
	"log"
)

var (
	db gorm.DB
)

func main() {
	initDB()

	app := martini.Classic()
	app.Use(gzip.All())
	app.Use(render.Renderer())

	app.Get("/photos", func(render render.Render) {
		photos := []models.Photo{}
		db.Find(&photos)

		render.JSON(200, photos)
	})

	app.Post("/photos", binding.MultipartForm(models.PhotoUpload{}), func(render render.Render, upload models.PhotoUpload) {
		file, err := upload.File.Open()
		defer file.Close()

		if err != nil {
			log.Panic(err)
			render.Error(500)
			return
		}

		dst, err := os.Create("./photos/" + upload.File.Filename)
		defer dst.Close()

		if err != nil {
			log.Fatal(err)
			render.Error(500)
			return
		}

		io.Copy(dst, file)

		render.Status(201)
	})

	app.Get("/photos/:id", func(params martini.Params, render render.Render) {
		photo := models.Photo{}

		db.First(&photo, params["id"])

		render.JSON(200, photo)
	})

	app.Run()
}

func initDB() {
	var err error

	db, err = gorm.Open("sqlite3", "./phototresor.db")

	if err != nil {
		panic(err)
		return
	}

	db.DB().Ping()
	db.LogMode(true)

	db.CreateTable(&models.Photo{})
}

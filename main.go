package main

import (
	"github.com/PhotoPeer/peer/app/models"
	"github.com/jinzhu/gorm"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/martini-contrib/binding"
	"os"
	"io"
	"log"
	"io/ioutil"
)

const PHOTOS_PATH string = "./photos/"
const DATABASE string = "photopeer.db"

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
		}

		os.Mkdir(PHOTOS_PATH, os.FileMode(uint(0775)))
		dst, err := os.Create(PHOTOS_PATH + upload.File.Filename)
		defer dst.Close()

		if err != nil {
			log.Fatal(err)
			render.Error(500)
		}

		io.Copy(dst, file)

		imageFile, _ := os.Open(PHOTOS_PATH + upload.File.Filename)

		photo := new(models.Photo)
		imageFileStat, _ := imageFile.Stat()
		photo.FileName = imageFileStat.Name()
		photo.FileSize = imageFileStat.Size()

		db.NewRecord(photo)
		db.Create(&photo)

		render.JSON(201, photo)
	})

	app.Get("/photos/:id", func(params martini.Params, render render.Render) {
		photo := models.Photo{}

		db.First(&photo, params["id"])

		render.JSON(200, photo)
	})

	app.Get("/photos/:id/download", func(params martini.Params, render render.Render) {
		photo := models.Photo{}

		db.First(&photo, params["id"])

		photoData, err := ioutil.ReadFile(PHOTOS_PATH + photo.FileName)
		if err != nil {
			render.Error(500)
		}

		render.Data(200, photoData)
	})

	app.Run()
}

func initDB() {
	var err error

	db, err = gorm.Open("sqlite3", DATABASE)

	if err != nil {
		panic(err)
		return
	}

	db.DB().Ping()
	db.LogMode(true)

	db.CreateTable(&models.Photo{})
}

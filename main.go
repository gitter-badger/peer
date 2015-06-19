package main

import (
	"github.com/PhotoPeer/peer/app/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/martini-contrib/binding"
	"github.com/jinzhu/gorm"
	"github.com/PhotoPeer/peer/app/controllers"
	"github.com/PhotoPeer/peer/env"
)

func main() {
	DB := initDB(env.DATABASE)

	app := martini.Classic()
	app.Use(gzip.All())
	app.Use(render.Renderer())
	app.Map(DB)

	app.Get("/photos", controllers.ListPhotos)
	app.Post("/photos", binding.MultipartForm(models.PhotoUpload{}), controllers.CreatePhoto)
	app.Get("/photos/:id", controllers.GetPhoto)
	app.Get("/photos/:id/download", controllers.DownloadPhoto)
	app.Delete("/photos/:id", controllers.DeletePhoto)

	app.Run()
}

func initDB(DATABASE string) gorm.DB {
	var err error
	var DB gorm.DB

	DB, err = gorm.Open("sqlite3", DATABASE)

	if err != nil {
		panic(err)
	}

	DB.DB().Ping()
	DB.LogMode(true)

	DB.CreateTable(&models.Photo{})

	return DB
}

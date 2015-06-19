package controllers
import (
	"github.com/PhotoPeer/peer/app/models"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"log"
	"os"
	"io"
	"github.com/PhotoPeer/peer/env"
	"github.com/go-martini/martini"
	"io/ioutil"
)


func ListPhotos(render render.Render, DB gorm.DB) {
	photos := []models.Photo{}
	DB.Find(&photos)

	render.JSON(200, photos)
}

func CreatePhoto(render render.Render, upload models.PhotoUpload, DB gorm.DB) {
	file, err := upload.File.Open()
	defer file.Close()

	if err != nil {
		log.Panic(err)
		render.Error(500)
	}

	os.Mkdir(env.PHOTOS_PATH, os.FileMode(uint(0775)))
	dst, err := os.Create(env.PHOTOS_PATH + upload.File.Filename)
	defer dst.Close()

	if err != nil {
		log.Fatal(err)
		render.Error(500)
	}

	io.Copy(dst, file)

	imageFile, _ := os.Open(env.PHOTOS_PATH + upload.File.Filename)

	photo := new(models.Photo)
	imageFileStat, _ := imageFile.Stat()
	photo.FileName = imageFileStat.Name()
	photo.FileSize = imageFileStat.Size()

	DB.NewRecord(photo)
	DB.Create(&photo)

	render.JSON(201, photo)
}

func GetPhoto(params martini.Params, render render.Render, DB gorm.DB) {
	photo := models.Photo{}

	DB.First(&photo, params["id"])

	render.JSON(200, photo)
}

func DownloadPhoto(params martini.Params, render render.Render, DB gorm.DB) {
	photo := models.Photo{}

	DB.First(&photo, params["id"])

	photoData, err := ioutil.ReadFile(env.PHOTOS_PATH + photo.FileName)
	if err != nil {
		render.Error(500)
	}

	render.Data(200, photoData)
}

func DeletePhoto(params martini.Params, render render.Render, DB gorm.DB) {
	photo := models.Photo{}

	DB.First(&photo, params["id"])

	if (photo.ID == 0) {
		render.Status(404)
	}

	err := os.Remove(env.PHOTOS_PATH + photo.FileName)
	if (err != nil) {
		render.Error(500)
	}

	DB.Delete(&photo)
	render.Status(200)
}

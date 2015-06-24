package controllers
import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"github.com/PhotoPeer/peer/app/models"
	"github.com/PhotoPeer/peer/env"
	"github.com/PhotoPeer/peer/app/models/Thumbnail"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
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

	if(photo.Empty()) {
		render.Error(404)
		return
	}

	render.JSON(200, photo)
}

const height = 220
func PhotoThumbnail(params martini.Params, render render.Render, DB gorm.DB, response http.ResponseWriter) {
	photo := models.Photo{}

	DB.First(&photo, params["id"])

	if(photo.Empty()) {
		render.Error(404)
		return
	}

	thumbnail := Thumbnail.New(photo, height)

	response.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(response, thumbnail.Image, &jpeg.Options{95})
}

func DeletePhoto(params martini.Params, render render.Render, DB gorm.DB) {
	photo := models.Photo{}

	DB.First(&photo, params["id"])

	if (photo.Empty()) {
		render.Error(404)
		return
	}

	err := os.Remove(env.PHOTOS_PATH + photo.FileName)
	if (err != nil) {
		render.Error(500)
		return
	}

	DB.Delete(&photo)
	render.Status(200)
}

package app
import (
	"image"
	"fmt"
	"crypto/md5"
	"os"
	"github.com/PhotoPeer/peer/env"
	"image/jpeg"
	"github.com/PhotoPeer/peer/app/models"
)

func GetThumbnail(photo models.Photo, height uint) image.Image {
	hash := getHash(photo, height)

	file, err := os.Open(env.THUMBNAILS_PATH + hash)
	defer file.Close()
	if err != nil {
		return nil
	}

	image, err := jpeg.Decode(file)
	if err != nil {
		return nil
	}

	return image
}

func PutThumbnail(photo models.Photo, height uint, image image.Image) {
	if (GetThumbnail(photo, height) == nil) {

		hash := getHash(photo, height)

		os.Mkdir(env.THUMBNAILS_PATH, os.FileMode(uint(0775)))
		dst, _ := os.Create(env.THUMBNAILS_PATH + hash)
		defer dst.Close()

		jpeg.Encode(dst, image, &jpeg.Options{95})
	}
}

func getHash(photo models.Photo, height uint) string {
	data := []byte(fmt.Sprintf("%s%d", photo.FileName, height))
	return fmt.Sprintf("%x", md5.Sum(data))
}

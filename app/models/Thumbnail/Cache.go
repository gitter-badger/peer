package Thumbnail
import (
	"os"
	"github.com/PhotoPeer/peer/env"
	"image/jpeg"
	"fmt"
	"crypto/md5"
	"image"
)

func getThumbnailImage(thumbnail Thumbnail) image.Image {
	hash := getHash(thumbnail)

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

func put(thumbnail Thumbnail) {
	if (getThumbnailImage(thumbnail) == nil) {

		hash := getHash(thumbnail)

		os.Mkdir(env.THUMBNAILS_PATH, os.FileMode(uint(0775)))
		dst, _ := os.Create(env.THUMBNAILS_PATH + hash)
		defer dst.Close()

		jpeg.Encode(dst, thumbnail.Image, &jpeg.Options{95})
	}
}

func getHash(thumbnail Thumbnail) string {
	data := []byte(fmt.Sprintf("%s%d", thumbnail.Photo.FileName, thumbnail.Height))
	return fmt.Sprintf("%x", md5.Sum(data))
}

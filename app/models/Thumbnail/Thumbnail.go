package Thumbnail
import (
	"github.com/PhotoPeer/peer/app/models"
	"image"
	"os"
	"github.com/PhotoPeer/peer/env"
	"image/jpeg"
	"github.com/nfnt/resize"
)

type Thumbnail struct {
	Photo  models.Photo
	Height uint
	Image  image.Image
}

func New(photo models.Photo, height uint) *Thumbnail {
	thumbnail := new(Thumbnail)

	thumbnail.Photo = photo
	thumbnail.Height = height
	thumbnail.Image = getThumbnailImage(*thumbnail)

	if thumbnail.Image == nil {
		thumbnail.Image = createThumbnailImage(*thumbnail)
		put(*thumbnail)
	}

	return thumbnail
}

func createThumbnailImage(thumbnail Thumbnail) image.Image {
	file, _ := os.Open(env.PHOTOS_PATH + thumbnail.Photo.FileName)
	defer file.Close()

	photoJpeg, _ := jpeg.Decode(file)

	return resize.Resize(0, thumbnail.Height, photoJpeg, resize.Bilinear)
}

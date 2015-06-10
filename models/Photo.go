package models
import (
	"github.com/jinzhu/gorm"
	"time"
	"mime/multipart"
)

type Photo struct {
	gorm.Model
	FileName     string
	FileSize     int
	FileMimeType string
	FileHash     string
	Width        int
	Height       int
	CapturedAt   time.Time
}

type PhotoUpload struct {
	File *multipart.FileHeader `form:"photo"`
}

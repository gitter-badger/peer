package models
import (
	"time"
	"mime/multipart"
)

type Photo struct {
	ID           uint `gorm:"primary_key" json:"id"`
	FileName     string `json:"file_name"`
	FileSize     int64 `json:"file_size"`
	FileMimeType string `json:"file_mime_type"`
	FileHash     string `json:"file_hash"`
	Width        int `json:"width"`
	Height       int `json:"height"`
	CapturedAt   time.Time `json:"captured_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type PhotoUpload struct {
	File *multipart.FileHeader `form:"file"`
}

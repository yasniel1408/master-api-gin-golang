package albums_models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

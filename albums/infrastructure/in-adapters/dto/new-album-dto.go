package albums_infrastructure_dto

import "errors"

type AlbumDTO struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// validate DTO
func (album *AlbumDTO) Validate() (bool, error) {
	if album.Title == "" || album.Artist == "" || album.Year == 0 {
		return false, errors.New("Invalid album, all fields are required")
	}
	return true, nil
}

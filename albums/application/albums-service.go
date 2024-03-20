package albums_application

import (
	"errors"

	albums_models "github.com/yasniel1408/master-api-gin-golang/albums/domain/models"
)

var albums = []albums_models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
	{ID: "4", Title: "Time Out", Artist: "Dave Brubeck", Year: 1959},
	{ID: "5", Title: "Somethin' Else", Artist: "Cannonball Adderley", Year: 1958},
	{ID: "6", Title: "The Sidewinder", Artist: "Lee Morgan", Year: 1963},
}

func GetAlbumsService() []albums_models.Album {
	return albums
}

func SaveAlbumsService(album albums_models.Album) []albums_models.Album {
	albums = append(albums, album)
	return albums
}

func GetAlbumService(albumID string) (albums_models.Album, error) {
	for _, album := range albums {
		if album.ID == albumID {
			return album, nil
		}
	}
	return albums_models.Album{}, errors.New("Album not found")
}

func DeleteAlbumsService(albumID string) ([]albums_models.Album, error) {
	for i, album := range albums {
		if album.ID != albumID {
			albums = append(albums[:i], albums[i+1:]...)
			return albums, nil
		}
	}
	return albums, errors.New("Album not found")
}

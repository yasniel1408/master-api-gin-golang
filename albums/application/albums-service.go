package albums_application

import (
	albums_models "github.com/yasniel1408/master-api-gin-golang/albums/domain/models"
	albums_infrastructure_dto "github.com/yasniel1408/master-api-gin-golang/albums/infrastructure/in-adapters/dto"
	albums_infrastructure_db "github.com/yasniel1408/master-api-gin-golang/albums/infrastructure/out-adapters"
)

type AlbumService struct {
	albums_infrastructure_db.AlbumDB
}

func (albums *AlbumService) GetAlbumsService() []albums_models.Album {
	return albums.ListAlbums()
}

func (albums *AlbumService) SaveAlbumsService(newAlbum albums_infrastructure_dto.AlbumDTO) albums_models.Album {

	// Convert newAlbum to albums_models.Album
	album := albums_models.Album{
		Title:  newAlbum.Title,
		Artist: newAlbum.Artist,
		Year:   newAlbum.Year,
	}
	var response = albums.SaveAlbum(album)
	return response
}

func (albums *AlbumService) GetAlbumService(albumID string) (albums_models.Album, error) {
	album, error := albums.GetAlbum(albumID)
	return album, error
}

func (albums *AlbumService) DeleteAlbumsService(albumID string) (albums_models.Album, error) {
	album, error := albums.DeleteAlbum(albumID)
	return album, error
}

package albums_infrastructure_db

import (
	"errors"
	"github.com/yasniel1408/master-api-gin-golang/shared"
	"gorm.io/gorm"

	albums_models "github.com/yasniel1408/master-api-gin-golang/albums/domain/models"
	albums_ports "github.com/yasniel1408/master-api-gin-golang/albums/domain/ports"
)

type AlbumDB struct {
	albums_ports.AlbumDBInterfacePort
}

func (e *AlbumDB) ListAlbums() []albums_models.Album {
	var albums []albums_models.Album
	shared.ConnectGorm().Find(&albums)
	return albums
}

func (e *AlbumDB) SaveAlbum(album albums_models.Album) albums_models.Album {
	shared.ConnectGorm().Create(&album)
	return album
}

func (e *AlbumDB) GetAlbum(albumID string) (albums_models.Album, error) {
	var album albums_models.Album
	result := shared.ConnectGorm().First(&album, "id = ?", albumID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return albums_models.Album{}, errors.New("album not found")
		}
		return albums_models.Album{}, result.Error
	}
	return album, nil
}

func (e *AlbumDB) DeleteAlbum(albumID string) (albums_models.Album, error) {
	var album albums_models.Album
	result := shared.ConnectGorm().Delete(&album, "id = ?", albumID)
	if result.RowsAffected == 0 {
		return albums_models.Album{}, errors.New("Álbum no encontrado")
	}
	// Aquí deberías devolver la lista de álbumes actualizada,
	// dependiendo de cómo estés gestionando tus datos.
	return album, nil
}

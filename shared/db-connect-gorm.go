package shared

import (
	"fmt"
	albums_models "github.com/yasniel1408/master-api-gin-golang/albums/domain/models"
	auth_models "github.com/yasniel1408/master-api-gin-golang/auth/domain/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// only in development mode
	err = db.AutoMigrate(&albums_models.Album{}, &auth_models.User{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

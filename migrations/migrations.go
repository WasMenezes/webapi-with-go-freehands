package migrations

import (
	"github.com/WasMenezes/webapi-with-go-freehands/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
	db.AutoMigrate(models.User{})
}

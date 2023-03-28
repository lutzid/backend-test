package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Db object will be used in all packages
var GlobalDB *gorm.DB

// Creates a sqlite db
func InitDatabase() (err error) {
	GlobalDB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		return
	}

	return
}

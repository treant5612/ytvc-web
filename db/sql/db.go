package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitSqlite(dbPath string) (err error) {
	db, err = gorm.Open("sqlite3", dbPath)
	return err
}

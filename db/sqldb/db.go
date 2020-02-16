package sqldb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/treant5612/ytvc-web/model"
)

var db *gorm.DB

func InitSqlite(dbPath string) (err error) {
	db, err = gorm.Open("sqlite3", dbPath)
	db.AutoMigrate(&model.RequestLog{})
	db.AutoMigrate(&model.Comment{})
	return err
}

func SaveRequestLog(reqLog *model.RequestLog) (err error) {
	return db.Save(reqLog).Error
}

package model

import (
	"time"
)

type Comment struct {
	ID       uint   `gorm:"primary_key"`
	Content  string `gorm:"size:1024` //内容
	NickName string //昵称
	Addr     string

	Time      time.Time
	DeletedAt *time.Time `gorm:"index:id_deleted"`
}

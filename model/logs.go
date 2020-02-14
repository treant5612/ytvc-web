package model

import "github.com/jinzhu/gorm"

type Log struct {
	gorm.Model
	RemoteAddr string
	RequestUrl string
	RequestMethod string

}
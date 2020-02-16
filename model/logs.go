package model

import (
	"time"
)

type RequestLog struct {
	ID            uint `gorm:"primary_key"`
	RemoteAddr    string
	RequestUrl    string
	RequestMethod string

	RequestTime     time.Time
	RequestDuration time.Duration

	ResponseStatus int
	ResponseSize   int
}

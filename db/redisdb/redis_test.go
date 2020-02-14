package redisdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/treant5612/ytvc-web/model"
	"testing"
	"time"
)

func init() {
	Init(&redis.Options{Addr: ":6379"})
}
func TestVideoDetail(t *testing.T) {
	video := &model.Video{
		Info: &model.VideoInfo{
			ID:            "testid",
			Title:         "",
			Description:   "",
			DatePublished: time.Time{},
			Uploader:      "",
			Duration:      0,
		},
		Files:    nil,
		Captions: nil,
	}
	SetVideoDetail(video)
	v, err := GetVideoDetail("texstid")

	rdb.Set("a","a",0)

	fmt.Println(v, err)

}

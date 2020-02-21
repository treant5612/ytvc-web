package redisdb

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/treant5612/ytvc-web/model"
	"log"
	"time"
)

var rdb *redis.Client

var (
	expire = time.Hour / 2
)

func SetVideoDetail(video *model.Video) {
	bytes, err := json.Marshal(video)
	if err != nil {
		log.Printf("redis set video detail failed:%v", err)
		return
	}
	rdb.Set("video."+video.Info.ID, bytes, expire)

}

func GetVideoDetail(id string) (video *model.Video, err error) {
	video = new(model.Video)
	bytes, err := rdb.Get("video." + id).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, video)
	return
}

func Init(options *redis.Options) {
	rdb = redis.NewClient(options)
	rdb.Options()
}

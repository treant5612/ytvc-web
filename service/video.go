package service

import (
	"errors"
	"fmt"
	"github.com/rylio/ytdl"
	"github.com/treant5612/ytvc-web/db/redisdb"
	"github.com/treant5612/ytvc-web/model"
	"github.com/treant5612/ytvc-web/utils"
	"log"
	"net/http"
	"sync"
)

func Video(videoId string) (video *model.Video, err error) {
	if v, err := redisdb.GetVideoDetail(videoId); err == nil {
		return v, nil
	}
	if videoId == "" {
		return nil, errors.New("get video id failed")
	}
	video, err = VideoInfo(videoId, "youtube")
	if err != nil {
		log.Printf("video err:%v", err)
		err = fmt.Errorf("get video info failed")
		return
	}
	redisdb.SetVideoDetail(video)
	return video, nil
}

func VideoInfo(videoId string, kind string) (video *model.Video, err error) {
	switch kind {
	case "youtube":
		return youtubeVideoInfo(videoId)
	}
	return nil, errors.New("cannot match url link")
}

/*
get video
*/
func youtubeVideoInfo(id string) (video *model.Video, err error) {
	video = new(model.Video)
	videoInfo := new(model.VideoInfo)
	var videoFiles []*model.FileInfo

	v, err := ytdl.GetVideoInfoFromID(id)
	if err != nil {
		return nil, err
	}
	video.Captions = v.Captions

	//基础信息
	if err = utils.Copy(videoInfo, v); err != nil {
		return nil, err
	}
	//缩略图
	videoInfo.ThumbnailUrl = v.GetThumbnailURL(ytdl.ThumbnailQualityDefault).String()
	//视频文件信息
	wg := &sync.WaitGroup{}
	client := http.DefaultClient
	//client.Timeout = time.Second
	for _, format := range v.Formats {
		f := new(model.FileInfo)
		if err = utils.Copy(f, format.Itag); err != nil {
			log.Printf("copy format to fileInfo failed:%v", err)
			continue
		}
		videoFiles = append(videoFiles, f)
		url, err := v.GetDownloadURL(format)
		wg.Add(1)
		//获取视频文件url
		//使用http.Head从url获取文件长度
		go func() {
			defer wg.Done()
			if err != nil {
				log.Printf("get download url failed :%v\n", err)
				return
			}
			f.Url = url.String()
			resp, err := client.Head(f.Url)
			if err != nil {
				log.Printf("http head video url failed:%v\n", err)
				return
			}
			if resp.StatusCode == http.StatusOK {
				f.Size = resp.ContentLength
			}
		}()
	}
	wg.Wait()
	video.Info = videoInfo
	video.Files = videoFiles
	return video, nil

}

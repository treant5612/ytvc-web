package service

import (
	"errors"
	"fmt"
	"github.com/rylio/ytdl"
	"github.com/treant5612/pornhub-dl"
	"github.com/treant5612/ytvc-web/db/redisdb"
	"github.com/treant5612/ytvc-web/model"
	"github.com/treant5612/ytvc-web/utils"
	"log"
	"net/url"
	"sync"
)

func Video(videoUrl string) (video *model.Video, err error) {
	u, err := url.Parse(videoUrl)
	if err != nil {
		return nil, err
	}

	videoId, kind := utils.ExtractVideoInfo(u)

	if v, err := redisdb.GetVideoDetail(videoId); err == nil {
		return v, nil
	}
	if videoId == "" {
		return nil, errors.New("get video id failed")
	}
	video, err = VideoInfo(videoId, kind)
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
	case "pornhub":
		return pornhubVideoInfo(videoId)
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
	//client.Timeout = time.Second
	for _, format := range v.Formats {
		f := new(model.FileInfo)
		if err = utils.Copy(f, format.Itag); err != nil {
			log.Printf("copy format to fileInfo failed:%v", err)
			continue
		}
		videoFiles = append(videoFiles, f)
		url, err := v.GetDownloadURL(format)
		if err != nil {
			log.Println(err)
			continue
		}
		wg.Add(1)
		//获取视频文件url
		//使用http.Head从url获取文件长度
		go func() {
			defer wg.Done()
			f.Url = url.String()
			size, err := utils.GetFileSize(f.Url)
			if err != nil {
				f.Size = size
			}
		}()
	}
	wg.Wait()
	video.Info = videoInfo
	video.Files = videoFiles
	return video, nil

}

func pornhubVideoInfo(id string) (video *model.Video, err error) {
	v, err := pornhub_dl.GetVideoInfoByKey(id)
	if err != nil {
		return nil, err
	}
	video = new(model.Video)
	videoInfo := new(model.VideoInfo)
	utils.Copy(videoInfo, v)
	video.Info = videoInfo
	wg := &sync.WaitGroup{}
	for _, f := range v.Files {
		file := new(model.FileInfo)
		file.Extension = f.Extension
		file.Number = f.Number
		file.Resolution = f.Resolution
		file.Url = f.Url
		wg.Add(1)
		go func() {
			file.Size, _ = utils.GetFileSize(file.Url)
			wg.Done()
		}()
		video.Files = append(video.Files, file)
	}
	wg.Wait()
	return video, nil
}

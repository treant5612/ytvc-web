package service

import (
	"errors"
	"github.com/rylio/ytdl"
	"github.com/treant5612/ytvc-web/model"
	"github.com/treant5612/ytvc-web/utils"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func VideoInfo(videoUrl string) (video *model.Video, err error) {
	vUrl, err := url.Parse(videoUrl)
	if err != nil {
		return nil, errors.New("parse url failed")
	}
	domain := utils.Domain(vUrl.Host)
	switch domain {
	case "youtu", "youtube":
		return youtubeVideoInfo(vUrl)
	}
	return nil, errors.New("cannot match url link")
}

func youtubeVideoInfo(url *url.URL) (video *model.Video, err error) {
	video = &model.Video{}
	v, err := ytdl.GetVideoInfoFromURL(url)
	//基础信息
	if err = utils.Copy(&video.Info, v); err != nil {
		return nil, err
	}
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
		video.Files = append(video.Files, f)
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
	return video, nil

}

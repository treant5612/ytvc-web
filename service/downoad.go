package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	ErrDownloadFailed = errors.New("download failed")
)

func DownloadInfo(videoId string, no int, kind string) (fileName, url string, err error) {
	video, err := VideoInfo(videoId, kind)
	//video,err :=redisdb.GetVideoDetail(videoId)
	if err != nil {
		return
	}

	for _, f := range video.Files {
		if f.Number == no {
			fileName = fmt.Sprintf("%s.%s", video.Info.Title, f.Extension)
			return fileName, f.Url, nil
		}
	}
	return "", "", errors.New("wrong file number")
}

type ReqOptions func(r *http.Request)

func Download(url string, options ...ReqOptions) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, op := range options {
		op(req)
	}
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return resp, nil
}

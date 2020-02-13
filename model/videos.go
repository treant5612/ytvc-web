package model

import (
	"google.golang.org/api/youtube/v3"
	"time"
)

type Video struct {
	Info     *VideoInfo          `json:"info"`
	Files    []*FileInfo        `json:"files"`
	Captions []*youtube.Caption `json:"captions"`
}

type VideoInfo struct {
	ID            string        `json:"id"`            // ID
	Title         string        `json:"title"`         //标题
	Description   string        `json:"description"`   //描述
	DatePublished time.Time     `json:"datePublished"` //发表日期
	Uploader      string        `json:"uploader"`      //上传者
	Duration      time.Duration `json:"duration"`      //时长
	ThumbnailUrl  string
}

type FileInfo struct {
	Number        int    `json:"number"`
	Extension     string `json:"extension"`  //格式
	Resolution    string `json:"resolution"` //解析度
	VideoEncoding string `json:"videoEncoding"`
	AudioEncoding string `json:"audioEncoding"`
	AudioBitrate  int    `json:"audioBitrate"`
	FPS           int    `json:"fps"`  // FPS are frames per second
	Url           string `json:"url"`  //视频下载地址
	Size          int64  `json:"size"` //视频大小
}



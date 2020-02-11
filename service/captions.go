package service

import (
	"github.com/treant5612/ytvc-web/manage/youtubeapi"
	"google.golang.org/api/youtube/v3"
)

func Captions(videoId string) (captions []*youtube.Caption, err error) {
	call := youtubeapi.ServiceFSC.Captions.List("snippet", videoId)
	resp, err := call.Do()
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

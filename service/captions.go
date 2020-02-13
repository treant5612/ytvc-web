package service

import (
	"fmt"
	"github.com/asticode/go-astisub"
	"github.com/treant5612/ytvc-web/manager/youtubeapi"
	"google.golang.org/api/youtube/v3"
	"path"
)

func Captions(videoId string) (captions []*youtube.Caption, err error) {
	call := youtubeapi.ServiceFSC.Captions.List("snippet", videoId)
	resp, err := call.Do()
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

/*
	Download caption and return it's path in local filesystem.
*/
func DownloadCaption(captionId string, tlang string) (path string, err error) {
	call := youtubeapi.ServiceFSC.Captions.Download(captionId)
	if tlang != "" {
		call = call.Tlang(tlang)
	}
	resp, err := call.Tfmt("srt").Download()
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return DownloadToLocalFile(resp.Body, captionId+tlang+".srt")
}

/*
	download two captions and merge them.
	return the filepath in local fs.
*/
func DownloadAndMergeCaption(mainId, mainTlang, secondaryId, secondaryTlang string) (path string, err error) {
	ch := make(chan error)
	var mainPath, secondaryPath string
	go func() {
		var err1 error
		mainPath, err1 = DownloadCaption(mainId, mainTlang)
		ch <- err1
	}()
	go func() {
		var err2 error
		secondaryPath, err2 = DownloadCaption(secondaryId, secondaryTlang)
		ch <- err2
	}()

	for i := 0; i < 2; i++ {
		err = <-ch
		if err != nil {
			return "", fmt.Errorf("download caption failed:%w", err)
		}
	}
	return merge(mainPath, secondaryPath)
}

func merge(mainPath, secondaryPath string) (filepath string, err error) {
	s1, err := astisub.OpenFile(mainPath)
	if err != nil {
		return
	}
	s2, err := astisub.OpenFile(secondaryPath)
	if err != nil {
		return
	}
	s1.Merge(s2)
	//Write 方法需要标注格式
	fileName := fmt.Sprintf("%s_%s.srt", path.Base(mainPath), path.Base(secondaryPath))
	filepath = path.Join(downloadPath, fileName)
	err = s1.Write(filepath)
	return
}

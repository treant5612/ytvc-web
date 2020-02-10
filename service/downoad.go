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

func Download(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil,err
	}
	if resp.StatusCode!=http.StatusOK{
		return nil,fmt.Errorf("%w",ErrDownloadFailed)
	}
	return resp,nil
}

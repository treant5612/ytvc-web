package utils

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

func Copy(dst interface{}, src interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	dstVal := reflect.ValueOf(dst).Elem()
	for i := 0; i < srcVal.NumField(); i++ {
		fieldName := srcVal.Type().Field(i).Name
		if v := dstVal.FieldByName(fieldName); !reflect.ValueOf(v).IsZero() {
			v.Set(srcVal.FieldByName(fieldName))
		}
	}
	return nil
}

func Domain(url string) (domain string) {
	regStr := `([a-z0-9][-a-z0-9]{0,62})\.([a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$`
	reg := regexp.MustCompile(regStr)
	matches := reg.FindStringSubmatch(url)
	if len(matches) < 1 {
		return ""
	}
	return matches[1]
}

func ExtractVideoInfo(u *url.URL) (id string, kind string) {

	switch true {
	case strings.Contains(u.Host, "pornhub"):
		id = u.Query().Get("viewkey")
		return id, "pornhub"
	case strings.Contains(u.Host, "youtu"):
		return ExtractVideoID(u), "youtube"
	}
	return "", ""
}

// GetVideoInfoFromShortURL fetches video info from a short youtube url
func ExtractVideoID(u *url.URL) string {
	switch u.Host {
	case "www.youtube.com", "youtube.com", "m.youtube.com":
		if u.Path == "/watch" {
			return u.Query().Get("v")
		}
		if strings.HasPrefix(u.Path, "/embed/") {
			return u.Path[7:]
		}
	case "youtu.be":
		if len(u.Path) > 1 {
			return u.Path[1:]
		}
	}
	return ""
}

func LanguageDisplay(lang string) string {
	tag, err := language.Parse(lang)
	if err != nil {
		return lang
	}
	return display.Chinese.Tags().Name(tag)
}

func GetFileSize(url string) (size int64, err error) {
	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, nil
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("request failed:%v", resp.StatusCode)
	}
	return resp.ContentLength, nil

}

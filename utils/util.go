package utils

import (
	"fmt"
	"reflect"
	"regexp"
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
	if len(matches)<1{
		return ""
	}
	return matches[1]
}

package utils

import (
	"fmt"
	"regexp"
	"testing"
)

type A struct {
	F1 string
	F2 string
	F3 string
}
type B struct {
	F0 string
	F1 string
	F2 string
}

func TestCopy(t *testing.T) {
	a := &A{
		"123", "234", "345",
	}
	b := &B{}
	err := Copy(b, a)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\na:%+v\nb:%+v", a, b)
}

func TestDomain(t *testing.T) {
	var url = "www.youtube.com"

	domainPattern := `([a-z0-9][-a-z0-9]{0,62})\.` +
		`(com\.cn|com\.hk|` +
		`cn|com|net|edu|gov|biz|org|info|pro|name|xxx|xyz|be|` +
		`me|top|cc|tv|tt)`
_=domainPattern
	reg := regexp.MustCompile(`([a-z0-9][-a-z0-9]{0,62})\.([a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$`)
	//reg := regexp.MustCompile(domainPattern)
	fmt.Println(reg.MatchString(url))
	matches := reg.FindStringSubmatch(url)
	fmt.Println(matches)
}

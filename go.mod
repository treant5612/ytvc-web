module github.com/treant5612/ytvc-web

go 1.13

require (
	github.com/asticode/go-astisub v0.2.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/rylio/ytdl v0.6.2
	github.com/treant5612/y2bcaptions v0.0.0-20200214082019-09d70222f437
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.2
	google.golang.org/api v0.17.0
)

replace github.com/rylio/ytdl v0.6.2 => github.com/treant5612/ytdl v0.6.3-0.20200214091400-8424bbb14de5

package service

import (
	"testing"
)

func TestYoutubeVideoInfo(t *testing.T) {
	tests := []struct {
		url       string
		assertion bool
	}{
		{
			url:       "https://www.youtube.com/watch?v=b75ZDL1n4Vs",
			assertion: true,
		},
		{
			url:       "https://www.youtube.com/watch?v",
			assertion: false,
		},
	}

	for _, test := range tests {
		video, err := VideoInfo(test.url)
		if !test.assertion && err == nil {
			t.Fail()
			continue
		}
		if test.assertion && err != nil {
			t.Fail()
			continue
		}
		if test.assertion {
			switch {
			case
				video.ID == "",
				len(video.Files) == 0,
				video.Files[0].Number == 0:
				t.Fail()
				continue
			}

		}
	}
}

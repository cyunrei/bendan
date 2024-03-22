package enhance

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_bilibili_enhance(t *testing.T) {
	tests := []struct {
		name     string
		url      *url.URL
		wantUrls []*url.URL
	}{
		{
			name: "TestCase1",
			url:  parseURL("https://bilibili.com/video/BV1hs411Q7zf"),
			wantUrls: []*url.URL{
				parseURL("https://bilibilibb.com/video/BV1hs411Q7zf"),
			},
		}, {
			name: "TestCase2",
			url:  parseURL("https://b23.tv/Of6ssFS"),
			wantUrls: []*url.URL{
				parseURL("https://b23bb.tv/Of6ssFS"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			bilibiliInstance := &bilibili{}
			if gotUrls := bilibiliInstance.enhance(tt.url); !reflect.DeepEqual(gotUrls, tt.wantUrls) {
				t1.Errorf("enhance() = %v, want %v", gotUrls, tt.wantUrls)
			}
		})
	}
}

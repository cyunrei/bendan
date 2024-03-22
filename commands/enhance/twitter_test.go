package enhance

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_twitter_enhance(t *testing.T) {
	tests := []struct {
		name     string
		url      *url.URL
		wantUrls []*url.URL
	}{
		{
			name: "TestCase1",
			url:  parseURL("https://twitter.com/daidai_kasame/status/1771107049847636385"),
			wantUrls: []*url.URL{
				parseURL("https://fxtwitter.com/daidai_kasame/status/1771107049847636385"),
				parseURL("https://vxtwitter.com/daidai_kasame/status/1771107049847636385"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			twitterInstance := &twitter{}
			if gotUrls := twitterInstance.enhance(tt.url); !reflect.DeepEqual(gotUrls, tt.wantUrls) {
				t1.Errorf("enhance() = %v, want %v", gotUrls, tt.wantUrls)
			}
		})
	}
}

func parseURL(urlStr string) *url.URL {
	parsedURL, _ := url.Parse(urlStr)
	return parsedURL
}

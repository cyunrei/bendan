package enhance

import (
	"net/url"
	"regexp"
)

var enhanceBilibiliRules = map[*regexp.Regexp]string{
	regexp.MustCompile(`^(https?://bilibili.com/)(.+)`): "https://bilibilibb.com/$2",
	regexp.MustCompile(`^(https?://b23.tv/)(.+)`):       "https://b23bb.tv/$2",
}

type bilibili struct {
}

func (t *bilibili) enhance(url *url.URL) (urls []*url.URL) {
	enhancedURL := url.String()
	for re, repl := range enhanceBilibiliRules {
		enhancedURL = re.ReplaceAllString(enhancedURL, repl)
	}
	enhancedParsedURL, _ := url.Parse(enhancedURL)
	urls = append(urls, enhancedParsedURL)
	return
}

package enhance

import (
	"net/url"
	"regexp"
)

var enhanceTwitterDomains = []string{
	"fxtwitter.com",
	"vxtwitter.com",
}

var reTwitter = regexp.MustCompile(`^https?://twitter.com/(.+?/status/\d+)`)

type twitter struct{}

func (t *twitter) enhance(url *url.URL) (urls []*url.URL) {
	if reTwitter.MatchString(url.String()) {
		for _, domain := range enhanceTwitterDomains {
			newURL := *url
			newURL.Host = domain
			urls = append(urls, &newURL)
		}
	}
	return
}

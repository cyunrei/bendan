package enhance

import "net/url"

type enhancer interface {
	enhance(*url.URL) []*url.URL
}

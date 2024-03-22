package enhance

import (
	"github.com/sxyazi/bendan/utils"
	"strings"
)

var Enhancers []enhancer

func init() {
	Enhancers = append(Enhancers, &bilibili{})
	Enhancers = append(Enhancers, &twitter{})
}

func Do(text string) (et []string) {
	urls := utils.ExtractUrls(text)
	if len(urls) != 1 {
		return []string{text}
	}
	for _, enhancer := range Enhancers {
		for _, u := range enhancer.enhance(urls[0]) {
			if u.String() != urls[0].String() {
				et = append(et, strings.ReplaceAll(text, urls[0].String(), u.String()))
			}
		}
	}
	return
}

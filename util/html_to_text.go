package util

import (
	"golang.org/x/net/html"
	"strings"
)

func ExtractTextFromHtml(s string, limit int) string {
	var desc = ""
	var descSlice []rune

	domDocTest := html.NewTokenizer(strings.NewReader(s))
	previousStartTokenTest := domDocTest.Token()

loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest // End of the document,  done
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			TxtContent := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			var runesTxt = []rune(TxtContent)
			if len(runesTxt) > 0 {
				descSlice = append(descSlice, []rune(TxtContent)...)
			}
		}
	}
	if len(descSlice) > limit {
		desc = string(descSlice[:limit])
	} else {
		desc = string(descSlice)
	}
	return desc
}

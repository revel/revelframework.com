package site

import (

	"strings"



)

// converts .html and .md etc into clean path
func StripExt(url string) string {

	if strings.HasSuffix(url, ".html") {
		url = url[0:len(url)-5]

	} else if strings.HasSuffix(url, ".md") {
		url = url[0:len(url)-3]
	}
	return url
}

// make lower and strip
func CleanStr(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
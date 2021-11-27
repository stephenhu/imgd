package main

import (
	"context"
	"strings"

	"github.com/chromedp/chromedp"

)

const (
	DOUBLE_SLASH					= "//"
	HTTP									= "http://"
	HTTPS									= "https://"
)

const (
	HTML_ANCHOR									= "a"
	HTML_IMG                  	= "img"
	HTML_SCRIPT                	= "script"
	HTML_TITLE                 	= "title"
	HTML_INPUT                  = "input"
)

const (
	HTML_ATTR_HREF							= "href"
	HTML_ATTR_SRC              	= "src"
	HTML_ATTR_DATA_LAZYLOAD    	= "data-lazyload"
)

const (
	CHROME_HEADLESS							= "http://127.0.0.1:9222"
)


func cleanUrl(s string) string {

	if strings.HasPrefix(s, DOUBLE_SLASH) {
		return strings.Replace(s, DOUBLE_SLASH, HTTPS, 1)
	} else {

		if !strings.Contains(s, HTTP) && !strings.Contains(s, HTTPS) {
			return HTTPS + s
		} else {
			return s
		}

	}

} // cleanUrl


func crawl(s string) {

	ctx, cancel := context.WithTimeout(context.Background(), 500)
	defer cancel()

	u := "imgd"

	chromedp.Run(ctx,
		chromedp.Navigate(s),
		chromedp.WaitVisible(TB_USER, chromedp.ByID),
		chromedp.SetValue(TB_USER, u, chromedp.ByID),
		chromedp.Submit(TB_SUBMIT))


} // crawl

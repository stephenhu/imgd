package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-redis/redis"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	JD_URL_SEARCH 		= "http://search.jd.com/Search?keyword=%s&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=%s&page=%d&s=1&click=0"
	JD_URL_ITEM  			= "item.jd.com"
	JD_URL_CCC_X  		= "ccc-x.jd.com"
	JD_URL_CD      		= "cd.jd.com"
	JD_JS_DESC				= "cd.jd.com/description"
)

const (
	DELIMITER         = ":"
	HASHTAG_COMMENT		= "#comment"
)


func job(imageUrl string, originUrl string, title string) string {
	return fmt.Sprintf("%s%s%s%s%s", imageUrl, DELIMITER, originUrl,
		DELIMITER, title)
} // job


func getTitle(doc *goquery.Document) string {

	title := ""

	doc.Find(HTML_TITLE).Each(func(index int, item *goquery.Selection) {

		t := item.Text()

		title = t

	})

	return title

} // getTitle


func storeImageList(link string, originUrl string) {

	l := cleanUrl(link)

	res, err := http.Get(l)

	if err != nil {
		appLogd("storeImageList", err.Error())
	} else {

		defer res.Body.Close()

		var j map[string] interface{}

		err := json.NewDecoder(res.Body).Decode(&j)

		if err != nil {
			appLogd("storeImageList", err.Error())
		} else {

			doc := html.NewTokenizer(strings.NewReader(j["content"].(string)))

			for {

				e := doc.Next()

				if e == html.ErrorToken {
					break
				} else {

					name, _ := doc.TagName()

					if string(name) == HTML_IMG {

						k, v, _ := doc.TagAttr()

						if string(k) == HTML_ATTR_DATA_LAZYLOAD {

							cleanUrl := cleanUrl(string(v))

							if cleanUrl != "" {

								_, err := client.ZAddNX(HASH_IMAGES, &redis.Z{Score: 0, Member: cleanUrl}).Result()

								if err != nil {
									appLogd("storeImageList", err.Error())
								}
	
							}

						}

					}

				}

			}

		}

	}

} // storeImageList


func (j Jd) ParsePics(link string) {

	re := regexp.MustCompile(
		`cd.jd.com/description/channel\?skuId=[0-9]+&mainSkuId=[0-9]+&cdn=[0-9]+`)

	res, err := http.Get(link)

	if err != nil {
		appLogd("ParsePics", err.Error())
	} else {

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			appLogd("ParsePics", err.Error())
		} else {

			doc.Find(HTML_SCRIPT).Each(func(index int, item *goquery.Selection) {

				t := item.Text()

				if strings.Contains(t, JD_JS_DESC) {

					match := re.FindString(t)

					storeImageList(match, link)

				}

			})

		}

	}

} // ParsePics


func (j Jd) Search(q string, page int) {

	res, err := http.Get(fmt.Sprintf(JD_URL_SEARCH, q, q, page))

	if err != nil {
		fmt.Println(err)
	} else {

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			fmt.Println(err)
		} else {

			doc.Find(HTML_ANCHOR).Each(func(index int, item *goquery.Selection) {

				l, _ := item.Attr(HTML_ATTR_HREF)

				if strings.Contains(l, JD_URL_ITEM) || strings.Contains(l, JD_URL_CCC_X) {
					
					s := cleanUrl(strings.TrimSuffix(l, HASHTAG_COMMENT))

					r, err := client.SAdd(HASH_PRODUCTS, s).Result()

					if err != nil {
						fmt.Println(err.Error())
					}

					if r > 0 {
						
						err = client.LPush(QUEUE_PRODUCTS, s).Err()

						if err != nil {
							fmt.Println(err.Error())
						}

					}

				}

			})

		}

	}

} // Search

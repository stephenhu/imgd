package main

import (
	//"bufio"
	"context"
	"fmt"
  "io/ioutil"
	"log"
	//"net/http"
	//"strings"
	"time"

	"github.com/chromedp/chromedp"
	//"github.com/PuerkitoBio/goquery"
)

const (
	TB_SEARCH		= "https://s.taobao.com/search?q=%s&imgfile=&js=1&stats_click=search_radio_all%3A1&initiative_id=staobaoz_20191013&ie=utf8"
	TB_SEARCH2  = "https://s.taobao.com/search?q=%s&imgfile=&commend=all&ssid=s5-e&search_type=item&sourceId=tb.index&spm=a21bo.2017.201856-taobao-item.1&ie=utf8&initiative_id=tbindexz_20170306&bcoffset=-12&ntoffset=-12&p4ppushleft=1%2C48&s=264"
	TB_SEARCH3	= "https://s.taobao.com/search?q=%s"
)

const (
	TB_JS_VAR		= "g_page_config"
	TB_USER     = "#TPL_username_1"
	TB_PASS     = "#TPL_password_1"
	TB_SUBMIT   = "#J_Form"
	TB_SUBMIT1  = "#J_SubmitStatic"
	TB_LOGIN    = "#J_Quick2Static"
	TB_LOGIN1   = "#J_Static2Quick"
	TB_SCALE    = "nc_1__scale_text"
)

type Auction struct {
	ID          string          `json:"nid"`
	DetailUrl   string          `json:"detail_url"`
}

type ItemData struct {
	Auctions		[]Auction				`json:"auctions"`
}

type ItemList struct {
  Data      	ItemData				`json:"data"`
}

type PageConfig struct {
  Items      	ItemList				`json:"itemList"`
}


func (t Tb) ParsePics(link string) {

} // ParsePics


func (t Tb) Search(q string, page int) {

	//crawl(fmt.Sprintf(TB_SEARCH3, q))

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()

	buf := []byte{}

	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf(TB_SEARCH3, q)),
		//chromedp.WaitVisible(TB_USER, chromedp.ByID),
		//chromedp.Value(TB_USER, &g, chromedp.ByID),
		chromedp.Click(TB_LOGIN, chromedp.ByID),
		chromedp.WaitVisible(TB_USER, chromedp.ByID),
		//chromedp.Sleep(5*time.Second),
		chromedp.SetValue(TB_USER, "", chromedp.ByID),
		chromedp.SetValue(TB_PASS, "", chromedp.ByID),
		chromedp.WaitVisible(TB_SUBMIT1, chromedp.ByID),
		chromedp.Click(TB_SUBMIT1, chromedp.ByID),
		chromedp.Click(TB_SUBMIT, chromedp.ByID),
		chromedp.Sleep(5*time.Second),		
		//chromedp.Value(TB_USER, &g, chromedp.ByID),		
		chromedp.CaptureScreenshot(&buf))
		//chromedp.WaitVisible("#main", chromedp.ByID),

		
	if err != nil {
		fmt.Println(err)
	} else {

		ioutil.WriteFile("tb.png", buf, 0644)
		
	}

	/*
	res, err := http.Get(fmt.Sprintf(TB_SEARCH3, q))

	if err != nil {
		fmt.Println(err)
	} else {

		defer res.Body.Close()

		buf, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		} else {
			ioutil.WriteFile("out.html", buf, 0644)
		}

	}
	*/

} // Search

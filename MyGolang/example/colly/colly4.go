package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"log"
	"regexp"
	"strings"
	"time"
)

func main() {
	//使用dom方式抓取
	testcollydom()
	log.Println("2秒後使用xpath抓取...........")
	//停止2秒後，使用xpath抓取
	time.Sleep(2*time.Second)
	fmt.Println()
	testcollyxpath()
}

func testcollydom(){
	//創建新的採集器
	c := colly.NewCollector(
		//這次在colly.NewCollector裏面加了一項colly.Async(true)，表示抓取時異步的
		colly.Async(true),
		//模擬瀏覽器
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	//限制採集規則
	//在Colly裏面非常方便控制併發度，只抓取符合某個(些)規則的URLS，有一句c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 5})，表示限制只抓取域名是douban(域名後綴和二級域名不限制)的地址，當然還支持正則匹配某些符合的 URLS，具體的可以看官方文檔。
	c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*",Parallelism:5})
	/*
	另外Limit方法中也限制了併發是5。爲什麼要控制併發度呢？因爲抓取的瓶頸往往來自對方網站的抓取頻率的限制，如果在一段時間內達到某個抓取頻率很容易被封，所以我們要控制抓取的頻率。另外爲了不給對方網站帶來額外的壓力和資源消耗，也應該控制你的抓取機制。
	*/

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(0).Text()))
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.alexa.com/")
	c.Wait()

}


func testcollyxpath(){

	c := colly.NewCollector(

		colly.Async(true),

		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),

		colly.MaxDepth(2),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*",Parallelism:5})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err)
		}
		nodes := htmlquery.Find(doc, `//ol[@class="grid_view"]/li//div[@class="hd"]`)
		for _, node := range nodes {
			url := htmlquery.FindOne(node, "./a/@href")
			title := htmlquery.FindOne(node, `.//span[@class="title"]/text()`)
			log.Println(strings.Split(htmlquery.InnerText(url), "/")[4],
				htmlquery.InnerText(title))
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// 查找行首以 ?start=0&filter= 的字符串（非貪婪模式）
		reg := regexp.MustCompile(`(?U)^\?start=(\d+)&filter=`)
		regMatch := reg.FindAllString(link, -1)
		//如果找的到的話
		if(len(regMatch) > 0){

			link = "https://www.alexa.com/"+regMatch[0]
			//訪問該鏈接
			e.Request.Visit(link)
		}

		// Visit link found on page
	})

	//結束
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	//採集開始
	c.Visit("https://www.alexa.com/")
	c.Wait()

}
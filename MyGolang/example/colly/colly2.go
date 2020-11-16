package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Article struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	URL      string `json:"url,omitempty"`
	Created  string `json:"created,omitempty"`
	Reads    string `json:"reads,omitempty"`
	Comments string `json:"comments,omitempty"`
	Feeds    string `json:"feeds,omitempty"`
}

// 資料持久化
func csvSave(fName string, data []Article) error {
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Title", "URL", "Created", "Reads", "Comments", "Feeds"})
	for _, v := range data {
		writer.Write([]string{strconv.Itoa(v.ID), v.Title, v.URL, v.Created, v.Reads, v.Comments, v.Feeds})
	}
	return nil
}

func main() {
	articles := make([]Article, 0, 200)
	// 1.準備收集器例項
	c := colly.NewCollector(
		// 開啟本機debug
		// colly.Debugger(&debug.LogDebugger{}),
		colly.AllowedDomains("learnku.com"),
		// 防止頁面重複下載
		// colly.CacheDir("./learnku_cache"),
	)

	// 2.分析頁面資料
	c.OnHTML("div.blog-article-list > .event", func(e *colly.HTMLElement) {
		article := Article{
			Title: e.ChildText("div.content > div.summary"),
			URL:   e.ChildAttr("div.content a.title", "href"),
			Feeds: e.ChildText("div.item-meta > a:first-child"),
		}
		// 查詢同一集合不同子項
		e.ForEach("div.content > div.meta > div.date>a", func(i int, el *colly.HTMLElement) {
			switch i {
			case 1:
				article.Created = el.Attr("data-tooltip")
			case 2:
				// 用空白切割字串
				article.Reads = strings.Fields(el.Text)[1]
			case 3:
				article.Comments = strings.Fields(el.Text)[1]
			}
		})
		// 正則匹配替換,字串轉整型
		article.ID, _ = strconv.Atoi(regexp.MustCompile(`\d+`).FindAllString(article.URL, -1)[0])
		articles = append(articles, article)
	})

	// 下一頁
	c.OnHTML("a[href].page-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	// 啟動
	c.Visit("https://learnku.com/blog/pardon")

	// 輸出
	csvSave("pardon.csv", articles)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(articles)

	// 顯示收集器的列印資訊
	log.Println(c)
}

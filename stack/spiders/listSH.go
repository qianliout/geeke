package spiders

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"outback/stack/items"
	"outback/stack/pipline"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

type NameCode struct {
	create pipline.Create
}

func NewNameCode(cre pipline.Create) *NameCode {

	return &NameCode{create: cre}
}

func (s *NameCode) ListSh() {
	// NewCollector(options ...func(*Collector)) *Collector
	// 声明初始化NewCollector对象时可以指定Agent，连接递归深度，URL过滤以及domain限制等
	c := colly.NewCollector(
		// colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"),
		// colly.AllowedDomains("sina.com.cn"),
		colly.MaxDepth(-1),
	)

	// 发出请求时附的回调
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		// r.Headers.Set("Host", "baidu.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
		r.Headers.Set("Referer", "http://www.sse.com.cn/")
		r.Headers.Set("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
		log.Print("Visiting", r.URL)
	})

	// 发现并访问其他年度的利润表
	c.OnHTML(`#con02-1`, func(e *colly.HTMLElement) {
		e.ForEach(" td a ", func(i int, el *colly.HTMLElement) {
			url := el.Attr("href")
			if strings.Contains(url, "vFD_ProfitStatement") {
				fmt.Println(url)

				// if err := e.Request.Visit(url); err != nil {
				// 	log.Error().Err(err).Msgf("Visit:%s", url)
				// }
			}
		})
	})
	// 发现并访问其他公司
	c.OnHTML(`#con02-1`, func(e *colly.HTMLElement) {
		e.ForEach(" td a ", func(i int, el *colly.HTMLElement) {
			url := el.Attr("href")
			if strings.Contains(url, "vFD_ProfitStatement") {
				fmt.Println(url)

				// if err := e.Request.Visit(url); err != nil {
				// 	log.Error().Err(err).Msgf("Visit:%s", url)
				// }
			}
		})
	})

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", r.StatusCode)
		// 设置context
		fmt.Println(r.Ctx.Get("url"))
	})

	c.OnResponse(func(resp *colly.Response) {
		fmt.Println("response received", resp.StatusCode)

		res := new(items.NubSh)
		err := json.Unmarshal(resp.Body, res)
		if err != nil {
			log.Error().Err(err).Msg("error")
			return
		}
		for i := range res.Result {
			if err := s.create.CreateCode(context.Background(), items.NameCode{Name: res.Result[i].Name, Code: res.Result[i].Code}); err != nil {
				log.Error().Err(err)
			}
			log.Info().Msgf("写入股票代码成功：%d", len(res.Result))
		}
	})

	// 对visit的线程数做限制，visit可以同时运行多个
	c.Limit(&colly.LimitRule{
		Parallelism: 1,
		Delay:       10 * time.Second,
	})

	c.OnError(func(response *colly.Response, err error) {
		log.Error().Err(err).Msg(response.Ctx.Get("url"))
	})
	for i := 1; i <= 66; i++ {
		url := fmt.Sprintf("http://query.sse.com.cn/security/stock/getStockListData2.do?&isPagination=true&stockCode=&csrcCode=&areaName=&stockType=1&pageHelp.cacheSize=1&pageHelp.beginPage=%d&pageHelp.pageSize=25&pageHelp.pageNo=%d&pageHelp.endPage=651&_=1633510195103", i, i)
		c.Visit(url)
	}
}

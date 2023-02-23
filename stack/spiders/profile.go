package spiders

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"outback/stack/items"
	"outback/stack/pipline"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"

	"github.com/gocolly/colly"
)

// type Document struct {
// 	*Selection
// 	Url      *url.URL
// 	rootNode *html.Node
// }
// type Selection struct {
// 	Nodes    []*html.Node
// 	document *Document
// 	prevSel  *Selection
// }

type StarkSpider struct {
	create pipline.Create
}

func NewStarkSpider(cre pipline.Create) *StarkSpider {

	return &StarkSpider{create: cre}
}

func (s *StarkSpider) Start() {

	// 声明初始化NewCollector对象时可以指定Agent，连接递归深度，URL过滤以及domain限制等
	c := colly.NewCollector(
		// colly.AllowedDomains("news.baidu.com"),
		colly.UserAgent("Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50"),
		colly.MaxDepth(-1),
	)

	// 发出请求时附的回调
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		// r.Headers.Set("Host", "baidu.com")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", "")
		r.Headers.Set("Referer", "http://vip.stock.finance.sina.com.cn/")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")
		r.Headers.Set("authority", "money.finance.sina.com.cn")
		r.Headers.Set("authority", "money.finance.sina.com.cn")
		r.Headers.Set("accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
		r.Headers.Set("Cookie", "name=sinaAds; post=massage; page=23333; NowDate=Sat Oct 30 2021 11:42:10 GMT+0800 (ä¸­å›½æ ‡å‡†æ—¶é—´); UOR=www.google.com,finance.sina.com.cn,; SINAGLOBAL=101.206.250.69_1606034230.695058; U_TRS1=00000000.5f3ccf7.5fba2341.a0a53296; kke_CnLv1_PPT_v2=know; UM_distinctid=179ea4edb87329-0f1204ed8e9d55-1f3b6254-13c680-179ea4edb88403; __gads=ID=eb2fd0309922d2cf-2262930549c90069:T=1623133708:RT=1623133708:S=ALNI_MZAaX1lyVKcp3US2kTz_5qbQ6cJ_g; SR_SEL=1_511; Apache=182.150.57.253_1635498661.533198; ULV=1635498802120:12:12:4:182.150.57.253_1635498661.533198:1635498661440; _s_upa=3; U_TRS2=000000fd.9066544.617bbf44.a46a46b3; FIN_ALL_VISITED=sh603155%2Csz002932%2Csz300459%2Csz002171%2Csz002756%2Csz002240; rotatecount=2; FINA_V_S_2=sh603155,sz002932,sz300459,sz002171,sz002756,sz002240; display=hidden; sinaH5EtagStatus=y")
	})

	c.OnResponse(func(resp *colly.Response) {
		url := resp.Request.URL.String()
		if resp.StatusCode != http.StatusOK {
			log.Info().Msgf("Status is not OK:%s", url)
			return
		}

		if strings.Contains(url, "hq.sinajs.cn/list") {
			s.ParseStarkPrice(bytes.NewReader(resp.Body))
		} else {
			// goquery直接读取resp.Body的内容
			htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))
			if err != nil {
				log.Error().Err(err).Msg("error")
				return
			}

			if strings.Contains(url, "vFD_ProfitStatement") {
				htmlDoc.Find(`#ProfitStatementNewTable0`).Each(s.ParseProfile)
			}

			if strings.Contains(url, "vFD_CashFlow") {
				htmlDoc.Find(`#ProfitStatementNewTable0`).Each(s.ParseCash)
			}

			if strings.Contains(url, "vFD_BalanceSheet") {
				htmlDoc.Find(`#BalanceSheetNewTable0`).Each(s.ParseBalance)
			}
		}

	})

	// 对visit的线程数做限制，visit可以同时运行多个
	if err := c.Limit(&colly.LimitRule{
		Parallelism: 1,
		Delay:       15 * time.Second,
	}); err != nil {
		log.Error().Err(err).Msg("Limit")
	}

	c.OnError(func(response *colly.Response, err error) {
		url := response.Request.URL.String()
		status := response.StatusCode
		if status == 456 {
			log.Info().Msg("IP已被封禁了")
			return
		}

		log.Error().Err(err).Msgf("get :%s", url)

	})

	codes, err := s.create.SearchNameCode(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("find name and code")
		return
	}
	years := []string{"2021", "2020"}
	for i := range codes {
		for j := range years {
			// if codes[i].Profile == 0 {
			// 	proUrl := fmt.Sprintf("https://money.finance.sina.com.cn/corp/go.php/vFD_ProfitStatement/stockid/%s/ctrl/%s/displaytype/4.phtml", codes[i].Code, years[j])
			// 	if err := c.Visit(proUrl); err != nil {
			// 		log.Error().Err(err).Msgf("Visit：%s", proUrl)
			// 	}
			// }

			// if codes[i].CashFlow == 0 {
			// 	cashUrl := fmt.Sprintf("https://money.finance.sina.com.cn/corp/go.php/vFD_CashFlow/stockid/%s/ctrl/%s/displaytype/4.phtml", codes[i].Code, years[j])
			// 	if err := c.Visit(cashUrl); err != nil {
			// 		log.Error().Err(err).Msgf("Visit: %s", cashUrl)
			// 	}
			// }
			//
			if codes[i].Balance == 0 {
				balanceUrl := fmt.Sprintf("https://money.finance.sina.com.cn/corp/go.php/vFD_BalanceSheet/stockid/%s/ctrl/%s/displaytype/4.phtml", codes[i].Code, years[j])
				if err := c.Visit(balanceUrl); err != nil {
					log.Error().Err(err).Msgf("Visit:%s", balanceUrl)
				}
			}
		}

		// ti := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
		// if codes[i].CrawlDate < ti.Unix() && len(codes[i].Code) > 0 && string([]rune(codes[i].Code)[:1]) != "3" {
		// 	if codes[i].SHSZ != "" {
		// 		balanceUrl := fmt.Sprintf("https://hq.sinajs.cn/list=%s%s", codes[i].SHSZ, codes[i].Code)
		// 		if err := c.Visit(balanceUrl); err != nil {
		// 			log.Error().Err(err).Msgf("Visit:%s", balanceUrl)
		// 		}
		// 	} else {
		// 		balanceUrl := fmt.Sprintf("https://hq.sinajs.cn/list=sh%s", codes[i].Code)
		// 		if err := c.Visit(balanceUrl); err != nil {
		// 			log.Error().Err(err).Msgf("Visit:%s", balanceUrl)
		// 		}
		//
		// 		balanceUrl = fmt.Sprintf("https://hq.sinajs.cn/list=sz%s", codes[i].Code)
		// 		if err := c.Visit(balanceUrl); err != nil {
		// 			log.Error().Err(err).Msgf("Visit:%s", balanceUrl)
		// 		}
		// 	}
		// }
	}

}

func (s *StarkSpider) ParseProfile(i int, selection *goquery.Selection) {

	res := make([]string, 0)
	name, code := "", ""

	selection.Find(" tr td").Each(
		func(i int, selection *goquery.Selection) {
			t := selection.Text()
			res = append(res, t)
		},
	)

	selection.Find("tr th ").Each(
		func(i int, selection *goquery.Selection) {
			na, co := parseNameCode(selection)
			name = na
			code = co
		})

	incomes, err := parseProfile(name, code, res)
	if err != nil {
		log.Error().Err(err).Msg("parseProfile")
		return
	}
	for i := range incomes {
		if err := s.create.CreateProfile(context.Background(), incomes[i]); err != nil {
			log.Error().Err(err).Msg("CreateProfile")
		}
		updater := map[string]interface{}{"profile": time.Now().Unix()}
		if err := s.create.UpdateNameCode(context.Background(), code, updater); err != nil {
			log.Error().Err(err).Msg("UpdateNameCode")
		}
	}

	log.Info().Msgf("写入利润表成功,Name %s:Code：%s", name, code)
}

func (s *StarkSpider) ParseCash(i int, selection *goquery.Selection) {
	res := make([]string, 0)
	name, code := "", ""

	selection.Find(" tr td ").Each(
		func(i int, selection *goquery.Selection) {
			t := selection.Text()
			res = append(res, t)
		},
	)

	selection.Find("tr th ").Each(
		func(i int, selection *goquery.Selection) {
			na, co := parseNameCode(selection)
			name = na
			code = co
		})

	cashs, err := parseCashFlow(name, code, res)
	if err != nil {
		log.Error().Err(err).Msg("parseProfile")
		return
	}
	for i := range cashs {
		if err := s.create.CreateCodeCashFlow(context.Background(), cashs[i]); err != nil {
			log.Error().Err(err).Msg("CreateProfile")
		}

		updater := map[string]interface{}{"cash_flow": time.Now().Unix()}
		if err := s.create.UpdateNameCode(context.Background(), code, updater); err != nil {
			log.Error().Err(err).Msg("UpdateNameCode")
		}
	}
	log.Info().Msgf("写入现金表成功,Name:%s Code：%s", name, code)

}

func (s *StarkSpider) ParseBalance(i int, selection *goquery.Selection) {

	res := make([]string, 0)
	name, code := "", ""

	selection.Find(" tr td").Each(
		func(i int, selection *goquery.Selection) {
			t := selection.Text()
			res = append(res, t)
		},
	)

	selection.Find("tr th ").Each(
		func(i int, selection *goquery.Selection) {
			na, co := parseNameCode(selection)
			name = na
			code = co
		})

	cashs, err := parseBalance(name, code, res)
	if err != nil {
		log.Error().Err(err).Msg("parseProfile")
		return
	}
	for i := range cashs {
		if err := s.create.CreateBalance(context.Background(), cashs[i]); err != nil {
			log.Error().Err(err).Msg("CreateProfile")
		}
		updater := map[string]interface{}{"balance": time.Now().Unix()}
		if err := s.create.UpdateNameCode(context.Background(), code, updater); err != nil {
			log.Error().Err(err).Msg("UpdateNameCode")
		}
	}
	log.Info().Msgf("写入资产负债表,Name: %s Code：%s", name, code)

}

func parseProfile(name, code string, res []string) ([]items.Profile, error) {
	per := Period(res)
	date := ReportDate(res)
	if per != len(date) {
		log.Error().Err(fmt.Errorf("日期列不符:%s,%s", name, code))
		return nil, fmt.Errorf("日期列不符")
	}
	ans := make([]items.Profile, per)
	for i := 0; i < len(date); i++ {
		ans[i].ReportingPeriod = date[i]
		ans[i].Code = code
		ans[i].Name = name
	}
	i := 0
	for i < len(res) {
		switch res[i] {
		case "一、营业总收入":
			for j := 0; j < per; j++ {
				// parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				ss := strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1)
				if ss != "" {
					if parseInt, err := strconv.ParseFloat(ss, 64); err == nil {
						ans[j].OperateAllIncome = int64(parseInt)
					}
				}
			}
			i = i + per
		case "营业收入":
			for j := 0; j < per; j++ {
				ss := strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1)
				if ss != "" {
					if parseInt, err := strconv.ParseFloat(ss, 64); err == nil {
						ans[j].OperateIncome = int64(parseInt)
					}
				}
			}
			i = i + per
		case "二、营业总成本":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].OperateAllCost = int64(parseInt)
				}
			}
			i = i + per
		case "营业成本":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].OperateCost = int64(parseInt)
				}
			}
			i = i + per
		case "营业税金及附加":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].Tax = int64(parseInt)
				}
			}
			i = i + per
		case "销售费用":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].SalesExpense = int64(parseInt)
				}
			}
			i = i + per
		case "管理费用":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].ManageExpense = int64(parseInt)
				}
			}
			i = i + per
		case "财务费用":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].FinancialExpense = int64(parseInt)
				}
			}
			i = i + per
		case "研发费用":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].RDExpense = int64(parseInt)
				}
			}
			i = i + per
		case "稀释每股收益(元/股)":
			for j := 0; j < per; j++ {
				parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				if err != nil {
					log.Error().Err(err).Msg("strconv.ParseInt")
				}
				ans[j].DilutedEarn = parseInt
			}
		}
		i = i + 1
	}

	return ans, nil
}

func Period(res []string) int {
	start := 0

	for i := 0; i < len(res); i++ {
		if res[i] == "报表日期" {
			start = i
		}
		if res[i] == "一、营业总收入" || res[i] == "一、经营活动产生的现金流量" || res[i] == "流动资产" || res[i] == "资产" || res[i] == "一、营业收入" {
			return i - start - 1
		}
	}
	return len(res)
}

func ReportDate(res []string) []string {
	ans := make([]string, 0)
	for i := 0; i < len(res); i++ {
		if res[i] == "报表日期" {
			for j := i + 1; j < len(res); j++ {
				if res[j] == "一、营业总收入" || res[j] == "一、经营活动产生的现金流量" || res[j] == "流动资产" || res[j] == "资产" || res[j] == "一、营业收入" {
					return ans
				}
				ans = append(ans, res[j])
			}
		}
	}
	return ans
}

func parseCashFlow(name, code string, res []string) ([]items.CashFlow, error) {
	per := Period(res)
	date := ReportDate(res)
	if per != len(date) {
		log.Error().Err(fmt.Errorf("日期列不符"))
		return nil, fmt.Errorf("日期列不符")
	}
	ans := make([]items.CashFlow, per)
	for i := 0; i < len(date); i++ {
		ans[i].ReportingPeriod = date[i]
		ans[i].Code = code
		ans[i].Name = name
	}
	i := 0
	for i < len(res) {
		switch res[i] {
		case "销售商品、提供劳务收到的现金":
			for j := 0; j < per; j++ {
				parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				if err != nil {
					log.Error().Err(err).Msg("strconv.ParseInt")
				}
				ans[j].SalesCash = parseInt
			}
			i = i + per
		case "经营活动现金流入小计":
			for j := 0; j < per; j++ {
				parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				if err != nil {
					log.Error().Err(err).Msg("strconv.ParseInt")
				}
				ans[j].SumInFow = parseInt
			}
			i = i + per
		case "购买商品、接受劳务支付的现金":
			for j := 0; j < per; j++ {
				parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				if err != nil {
					log.Error().Err(err).Msg("strconv.ParseInt")
				}
				ans[j].BuyCash = parseInt
			}
			i = i + per
		case "经营活动现金流出小计":
			for j := 0; j < per; j++ {
				parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				if err != nil {
					log.Error().Err(err).Msg("strconv.ParseInt")
				}
				ans[j].SumOutFow = parseInt
			}
			i = i + per
		case "经营活动产生的现金流量净额":
			for j := 0; j < per; j++ {
				parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64)
				if err != nil {
					log.Error().Err(err).Msg("strconv.ParseInt")
				}
				ans[j].NetCashFlow = parseInt
			}
			i = i + per
		}
		i = i + 1
	}

	return ans, nil
}

func parseBalance(name, code string, res []string) ([]items.Balance, error) {
	per := Period(res)
	date := ReportDate(res)
	if per != len(date) {
		log.Error().Err(fmt.Errorf("日期列不符"))
		return nil, fmt.Errorf("日期列不符")
	}
	ans := make([]items.Balance, per)
	for i := 0; i < len(date); i++ {
		ans[i].ReportingPeriod = date[i]
		ans[i].Code = code
		ans[i].Name = name
	}
	i := 0
	for i < len(res) {
		switch res[i] {
		case "货币资金":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].MoneyFunds = int64(parseInt)
				}
			}
			i = i + per
		case "交易性金融资产":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].TransFinance = int64(parseInt)
				}
			}
			i = i + per
		case "存货":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].Stock = int64(parseInt)
				}
			}
			i = i + per
		case "短期借款":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].ShortLoan = int64(parseInt)
				}
			}
			i = i + per
		case "长期借款":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].LongLoan = int64(parseInt)
				}
			}
			i = i + per

		case "实收资本(或股本)", "股本", "实收资本":
			for j := 0; j < per; j++ {
				if parseInt, err := strconv.ParseFloat(strings.Replace(strings.Replace(res[i+j+1], ",", "", -1), "-", "", -1), 64); err == nil {
					ans[j].Capital = int64(parseInt)
				}
			}
		}
		i = i + 1
	}

	return ans, nil
}

var (
	re = regexp.MustCompile("[0-9]+")
)

// ParseStarkPrice 解析当前股价
func (s *StarkSpider) ParseStarkPrice(reader io.Reader) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Info().Msg("解析结果出错")
		return
	}
	// 解析code

	regexp.MustCompile("[0-9]+")

	split := strings.Split(string(body), ",")
	ti := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	if len(split) >= 4 {
		codes := re.FindAllString(split[0], -1)

		price, err := strconv.ParseFloat(split[3], 64)
		if err == nil && len(codes) > 0 {
			update := map[string]interface{}{
				"stock_price": price,
				"crawl_date":  ti.Unix(),
			}
			if strings.Contains(split[0], "sh") {
				update["shsz"] = "sh"
			} else if strings.Contains(split[0], "sz") {
				update["shsz"] = "sz"
			}

			if err := s.create.UpdateNameCode(context.Background(), codes[0], update); err != nil {
				log.Info().Err(err).Msgf("存储：%s 的股价出错", codes[0])
			} else {
				log.Info().Msgf("存储：%s 的股价成功:%s:%d", codes[0], ti.String(), price)
			}
		}
	}
}

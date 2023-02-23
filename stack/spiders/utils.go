package spiders

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseNameCode(selection *goquery.Selection) (name, code string) {
	t := selection.Text()
	rcode := regexp.MustCompile(`\(.*?\)`)
	co := rcode.FindString(t)
	code = strings.Replace(strings.Replace(co, "(", "", -1), ")", "", -1)
	rname := regexp.MustCompile(`(.*?)\(`)
	na := rname.FindString(t)
	name = strings.Replace(strings.Replace(na, "(", "", -1), ")", "", -1)
	return name, code
}

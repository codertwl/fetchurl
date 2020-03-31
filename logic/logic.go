package logic

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
	"strings"
)

func isValidUrl(url string) bool {

	reg, _ := regexp.Compile(`[\w#/.:]+`)
	if len(url) < 3 {
		return false
	} else if !reg.MatchString(url) {
		return false
	} else if !strings.Contains(url, ".") {
		return false
	} else if (!strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://")) || strings.HasSuffix(url, ".exe") || strings.HasSuffix(url, ":void(0);") {
		return false
	}

	return true
}

func DoFetch(url string, leftDeep int, mapUrl map[string]bool, fp *os.File) {

	if leftDeep <= 0 {
		return
	}

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("url to doc failed err:", err.Error())
		return
	}

	f := func(i int, elem *goquery.Selection) {

		href, IsExist := elem.Attr("href")
		if IsExist == false {
			return
		}

		href = strings.Replace(href, " ", "", -1)
		href = strings.Replace(href, "\n", "", -1)
		if !isValidUrl(href) {
			return
		}

		if _, ok := mapUrl[href]; ok {
			return
		}

		mapUrl[href] = true
		fp.WriteString(href + "\n")
		DoFetch(href, leftDeep-1, mapUrl, fp)
	}

	doc.Find("a").Each(f)
}

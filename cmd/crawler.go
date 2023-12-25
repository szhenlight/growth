package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	var website string

	flag.StringVar(&website, "website", "https://www.baidu.com", "")
	flag.Parse()
	// 发起HTTP GET请求
	response, err := http.Get(website)
	if err != nil {
		fmt.Println("请求发生错误:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应发生错误:", err)
		return
	}

	// 解析HTML内容
	title := parseTitle(string(body))
	fmt.Println("标题:", title)
}

// 解析标题函数
func parseTitle(html string) string {
	// 使用正则表达式匹配标题标签中的内容
	titleRegex := "<title>(.*?)</title>"
	match := regexp.MustCompile(titleRegex).FindStringSubmatch(html)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

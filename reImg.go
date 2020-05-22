package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var (
	reImg = `wx[1-9].sinaimg.cn.*?(png|jpg|gif)`
)

func htmlGet(url string, host string) string {
	// 实例化一个http请求
	client := &http.Client{}
	// 构造一个请求的地址,方法，url，数据(get用nil)
	req, _ := http.NewRequest("GET", url, nil)
	// 添加header
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Add("host", host)
	// 提交构造的请求
	resp, _ := client.Do(req)
	// 关闭请求
	defer resp.Body.Close()
	// 读取返回的数据
	sult, _ := ioutil.ReadAll(resp.Body)
	// 将读书的字节转成string
	html := string(sult)
	// 返回结果
	return html

}
func reImge(reImg string, html string) []string {
	//构造一个正则表达式的匹配实例
	re := regexp.MustCompile(reImg)
	// 用构造的表达式去查询Html内容
	urls := re.FindAllString(html, -1)
	// 返回结果集
	return urls

}
func downImg(url string) {
	//配置下载的文件名，可以走一个函数
	urlsp := strings.Split(url, "/")
	filename := urlsp[len(urlsp)-1]
	// 开始请求图片地址
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	// 写入文件，如果报错则提示
	err2 := ioutil.WriteFile(filename, imgBytes, 0666)
	if err2 != nil {
		fmt.Println("下载失败！")
	}

}
func main() {
	html := htmlGet("http://jandan.net/pic/MjAyMDA1MjItMjEx#comments", "jandan.net")

	urls := reImge(reImg, html)
	for _, img := range urls {
		img_url := "https://" + img
		downImg(img_url)
	}
	//	fmt.Println(urls)
	// 需要优化的地方，文件名的获取。
	// 遇到异常就跳过
	// 请求超时设置
	// 异步请求
}

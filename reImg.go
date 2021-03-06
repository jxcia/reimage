package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	reImg      = `wx[1-9].sinaimg.cn.*?(png|jpg|gif)`
	chSem      = make(chan int, 5)
	downloadWG sync.WaitGroup
)

func GetRandomInt(start, end int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := r.Intn(end - start)
	randNum := strconv.Itoa(ret)
	return randNum
}

func urlGet(url string, host string) []byte {
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
	html, _ := ioutil.ReadAll(resp.Body)
	// 返回结果为[]byte,方便不同的数据做另外的处理
	return html

}

func reImge(reImg string, html []byte) []string {
	//对传入的html做string 处理
	page := string(html)
	//构造一个正则表达式的匹配实例
	re := regexp.MustCompile(reImg)
	// 用构造的表达式去查询Html内容
	urls := re.FindAllString(page, -1)
	// 返回结果集
	return urls

}
func getName(url string) string {
	name := strings.Split(url, "/")
	randNum := GetRandomInt(100, 200)
	filename := "./img/" + randNum + name[len(name)-1]
	return filename
}

/*
func downImgAnsyc(url string, host string) {
	go func() {
		downloadWG.Add(1)
		chSem <- 234234
		go downImg(url, host)
		<-chSem
		downloadWG.Done()
	}()

}
*/

func downImg(url string, host string) {
	// 开始请求图片地址
	imgBytes := urlGet(url, host)

	// 写入文件，如果报错则提示
	filename := getName(url)
	err2 := ioutil.WriteFile(filename, imgBytes, 0666)
	if err2 != nil {
		fmt.Println(url, "下载失败！")

	}
	fmt.Println(url, "下载完成！")

}
func main() {
	html := urlGet("http://jandan.net/ooxx/MjAyMDA1MjQtMTU0#comments", "jandan.net")

	urls := reImge(reImg, html)
	for _, img := range urls {
		img_url := "https://" + img
		downImg(img_url, "jandan.net")
	}
	//	downloadWG.Wait()
	//	fmt.Println(urls)
	// 遇到异常就跳过
	// 请求超时设置
	// 异步请求
}

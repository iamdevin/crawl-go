package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reQQEmail = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱正则
	reLink    = ``                                            //注意符号 链接正则
)


//获取指定url的HTMl
func GetPageHtmlStr(url string) string {
	//1.去网站拿数据
	resp, err := http.Get(url)
	HandleErr(err, "http.Get URL")
	defer resp.Body.Close()
	//2.读取数据
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, "readBody err")
	//字节转字符串
	pageString := string(pageBytes)
	return pageString

}

//获取指定网页上的电子邮件地址
func GetMail(url string, reStr string) {
	pageString := GetPageHtmlStr(url)
	re := regexp.MustCompile(reStr)
	results := re.FindAllStringSubmatch(pageString, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}

}

// 获取链接地址
func GetLink(url string, reStr string) {
	pageString := GetPageHtmlStr(url)
	re := regexp.MustCompile(reStr)
	results := re.FindAllStringSubmatch(pageString, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}

}


// 处理异常
func HandleErr(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main() {
	GetMail("https://tieba.baidu.com/p/1753244412?red_tag=3243539944", reQQEmail)
}

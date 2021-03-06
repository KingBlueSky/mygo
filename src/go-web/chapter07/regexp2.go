package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.nbaidu.com")
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error:", err)
		return
	}

	src := string(body)

	// 将html标签全部转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	// 去除style
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\>\\</style\\>")
	src = re.ReplaceAllString(src, "")

	// 去除script
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\>\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// 去除所有尖括号内的html代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

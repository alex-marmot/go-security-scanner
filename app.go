package main

import (
	"fmt"
	"go-crawler/spider"
)

func main() {
	res := spider.Crawler("https://www.shiyanlou.com")
	fmt.Println(res)
}

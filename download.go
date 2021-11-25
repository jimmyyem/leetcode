package main

/**
github.com/levigross/grequests 用于上传和下载的go package
@see https://github.com/levigross/grequests
*/

import (
	"fmt"
	"github.com/levigross/grequests"
	"leetcode/lib"
	"log"
	"os"
	"path"
	"path/filepath"
)

const DOWNLOAD_DIR = "download_files"

func main() {
	pic := "https://i.ugc.corp3g.cn/group61/M00/B4/4D/CmQUOV6XGCeEXL5PAAAAAHYD6bo317752545.jpg"
	download(pic)

	fmt.Println(path.Ext(pic))
}

// 下载图片
func download(url string) string {
	if _, err := os.Stat(DOWNLOAD_DIR); err != nil {
		os.Mkdir(DOWNLOAD_DIR, 0755)
	}

	fileName := lib.GetRandStringMask(32)
	ext := filepath.Ext(url)
	path := fmt.Sprintf("%s/%s%s", DOWNLOAD_DIR, fileName, ext)

	options := &grequests.RequestOptions{
		Params: map[string]string{
			"hello": "sss",
		},
	}
	resp, err := grequests.Get(url, options)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	if err := resp.DownloadToFile(path); err != nil {
		log.Println("Unable to download file: ", err)
	}

	return path
}

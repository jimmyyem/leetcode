package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	//url := "http://1257120875.vod2.myqcloud.com/0ef121cdvodtransgzp1257120875/3055695e5285890780828799271/v.f230.m3u8"
	url := "http://devimages.apple.com/iphone/samples/bipbop/bipbopall.m3u8"

	// 读取m3u8文件并且解析出ts文件列表
	body := getMeta(url)

	wg := &sync.WaitGroup{}
	wg.Add(len(body))

	// 根据ts文件列表获取ts文件内容
	for i := 0; i < len(body); i++ {
		go getTsContent(body[i], i, wg)
	}

	wg.Wait()

	concatVideo()

	fmt.Println("finish", body)
}

// 获取m3u8格式内容
func getMeta(url string) (res []string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	bodyReader := bufio.NewReader(resp.Body)

	baseUrl := getUrlBase(url)
	//fmt.Println(baseUrl)

	for {
		lineByte, _, err := bodyReader.ReadLine()
		fmt.Printf("%s\n", lineByte)
		if err == io.EOF {
			break
		}
		line := strings.Trim(string(lineByte), "\r\n\t")

		// 判断是否是合法的ts文件地址
		if strings.HasPrefix(line, "#") {
			continue
		}

		// 拼凑完整地址
		if strings.Contains(line, "/") {
			line = line[1:]
		}
		line = baseUrl + "/" + line
		res = append(res, line)
	}
	return
}

// 获取每个setment内容，并保存文件
func getTsContent(url string, idx int, group *sync.WaitGroup) {
	defer group.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fileName := fmt.Sprintf("file_tmp/%d.ts", idx)
	fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer fp.Close()

	fp.Write(content)
}

// 获取URL中除去最后一段的部分
func getUrlBase(address string) string {
	segment := strings.Split(address, "/")
	return strings.Join(segment[:len(segment)-1], "/")
}

// 合并ts文件成mp4格式
func concatVideo() {
	//cmd := exec.Command("/bin/echo *ts > whole.mp4")
	//res, err := cmd.Output()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(res))
	fmt.Println("/usr/local/Cellar/ffmpeg/4.4/bin/ffmpeg -allowed_extensions ALL -i http://xxoo.com/index.m3u8 -c copy index.mp4")
	return
}

func testCmd() {
	var ip, whoami []byte
	var err error
	var cmd *exec.Cmd

	// 执行单个shell命令时, 直接运行即可
	cmd = exec.Command("whoami")
	if whoami, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 默认输出有一个换行
	fmt.Println(string(whoami))
	// 指定参数后过滤换行符
	fmt.Println(strings.Trim(string(whoami), "\n"))

	fmt.Println("====")

	// mac平台获取ip地址
	// 执行连续的shell命令时, 需要注意指定执行路径和参数, 否则运行出错
	cmd = exec.Command("/bin/sh", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(ip))
	fmt.Println(strings.Trim(string(ip), "\n"))
}

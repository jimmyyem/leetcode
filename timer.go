package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("%v\n", time.Time{})
	var beijingZone = time.FixedZone("Asia/Shanghai", 8*60*60)
	fmt.Println(beijingZone.String(), time.UTC, time.Local)

	ticker := time.NewTicker(time.Second)
	count := 0

	go func() {
		for {
			t := <-ticker.C
			count++

			fmt.Printf("当前时间：%v, count=%d\n", t, count)
			if count == 10 {
				ticker.Stop()
				runtime.Goexit()
			}
		}
	}()

	for {
		time.Sleep(time.Second * 1)
	}
}

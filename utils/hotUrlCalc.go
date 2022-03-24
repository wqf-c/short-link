package utils

import (
	"sync"
	"time"
)

//url访问次数放到内存

var (
	hotLink map[string]int
	mtx     sync.Mutex
)

func init() {
	hotLink = make(map[string]int)
	//定时清理url访问次数
	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for { //循环
			<-ticker.C
			mtx.Lock()
			hotLink = make(map[string]int)
			mtx.Unlock()
		}
	}()
}

func VisitUrl(url string) (freq int) {
	mtx.Lock()
	defer mtx.Unlock()
	hotLink[url]++
	freq = hotLink[url]
	return
}

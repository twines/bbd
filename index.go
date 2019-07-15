package main

import (
	"./bilibili"
	"math"
	"os"
	"sync"
)

func main() {

	/*
		f, e := os.Create("./bilibili/bi.txt")
		defer f.Close()
		if e != nil {
			os.Exit(0)
		}
		w := csv.NewWriter(f)

		w.WriteAll([][]string{
			{"A"},
			{"B"},
		})
		w.Flush()
	*/
	//keywords := []string{"豫剧", "京剧",  }
	keywords := []string{"豫剧", "京剧", "秦腔",
		"曲剧", "晋剧", "二人转", "太平调", "川剧",
	}

	v := make(chan bool, len(keywords))
	//defer close(v)
	//for _, value := range keywords {
	//	keyword := value
	//	go bilibili.Bilibili(1, keyword, v)
	//
	//}
	//for i := 0; i < cap(v); i++ {
	//	<-v
	//}

	//控制后续启动之后的并发量
	var step = 3
	var index = 0
	var lock sync.RWMutex
	next := int(math.Min(float64(index+step), float64(len(keywords))))
	sub := keywords[index:next]
	for _, value := range sub {
		value := value
		go stepper(v, value)
	}
	for {
		select {
		case <-v:
			lock.Lock()
			index = index + 1

			lock.Unlock()
			if index < len(keywords) {
				go stepper(v, keywords[index])
			} else {
				os.Exit(0)
			}
		}
	}
}
func stepper(v chan bool, keyword string) {
	go bilibili.Bilibili(1, keyword, v)

}

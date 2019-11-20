package hotload

import (
	"log"
	"time"
)

/*
 * Golang下针对特定文件监控 以及热更新
 */

type HotLoad struct {
}

func New() *HotLoad {
	return &HotLoad{}
}

// 更新配置方法
type LoadFunc func(filePath string)

func (h *HotLoad) Load(filePath string, dstFunc LoadFunc) {
	md51 := ""
	for {
		select {
		case <-time.After(time.Millisecond * 200):
			md5, e := md5File(filePath)
			if e != nil {
				log.Fatalln(e)
			}
			if md51 != md5 {
				// 如果文集hash发生改变
				// 重新读取文件并调用dstFunc方法 进行热更新
				dstFunc(filePath)
				md51 = md5
			}
			continue
		}
	}
}

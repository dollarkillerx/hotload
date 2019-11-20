package hotload

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestHotLoad_Load(t *testing.T) {
	load := New()
	path := "test.conf"
	load.Load(path, update)
}

func update(filePath string) {
	log.Println("更新了配置")
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		log.Fatalln(e)
	}

	log.Println(string(bytes))
}

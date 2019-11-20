# hotload
golang 下针对某个文件进行监控 (可应用与配置文件热加载)
### 如果使用?
`go get github.com/dollarkillerx/hotload`

``` go
func TestHotLoad_Load(t *testing.T) {
	load := New()           // 初始化hotload
	path := "test.conf"
	load.Load(path, update) // 配置监控
}

// 定义更新配置方法
func update(filePath string) {
	log.Println("更新了配置")
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		log.Fatalln(e)
	}

	log.Println(string(bytes))
}
```
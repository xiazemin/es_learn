package main

import (
	"fmt"

	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var (
	filterText  = `hells习近平s xi我习是需要过滤的近平内容，内容为：**文@@件，需要习近平过滤。。。`
	filterText1 = `hells习近平s xi我习是需要过滤的近平内容，内容为：**文@@件，需要习近平过滤。。。`
)

func getFM(words []string) *filter.DirtyManager {
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		DataSource: words,
	})
	if err != nil {
		panic(err)
	}
	return filter.NewDirtyManager(memStore) //在这里构建的前缀树，如果有更新的话可以通过chan动态更新
}

func getResult(filterManage *filter.DirtyManager, filterText, filterText1 string) {
	result, err := filterManage.Filter().FilterResult(filterText, '*', '@')
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result1, err1 := filterManage.Filter().FilterResult(filterText1, '*', '@')
	if err1 != nil {
		panic(err)
	}
	fmt.Println(result1)
}
func main() {
	filterManage := getFM([]string{"文件", "习近平"})
	getResult(filterManage, filterText, filterText1)
}

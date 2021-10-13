package main

import (
	"fmt"
)

func main() {
	userMap := make(map[string]string)
	userMap["Tom"] = "My is Tom"
	userMap["Shure"] = "My is Shure"
	// 遍历key，value
	for k, v := range userMap {
		fmt.Println(k, v)
	}

	// 遍历 key
	for k := range userMap {
		fmt.Println(k)
	}

	// 删除元素键值
	delete(userMap, "Tom")
	fmt.Println(userMap)
}

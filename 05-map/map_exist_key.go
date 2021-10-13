package main

import (
	"fmt"
)

func main() {
	userMap := make(map[string]string)
	userMap["Tom"] = "My is Tom"
	userMap["Shure"] = "My is Shure"
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := userMap["Tom"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

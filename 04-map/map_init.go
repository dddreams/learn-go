package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["Tom"] = 90
	scoreMap["Shure"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["Shure"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "Shure",
		"password": "123456",
	}
	fmt.Println(userInfo)
}

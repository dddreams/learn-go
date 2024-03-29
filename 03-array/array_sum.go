package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * 求数组所有元素之和
 */

func sumArr(a [10]int) int {
	var sum int = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	// 若想做一个真正的随机数，要种子
	// seed()种子默认是1
	//rand.Seed(1)
	rand.Seed(time.Now().Unix())

	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}
	sum := sumArr(arr)
	fmt.Printf("sum=%d\n", sum)
}

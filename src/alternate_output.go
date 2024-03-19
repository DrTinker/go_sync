package src

import (
	"strconv"
	"sync"
)

/*
	题目来源：百度一面
	两个goroutine，一个负责打印数字，另一个负责打印字符，最后的效果：12AB34CD56EF78GH910IJ
	实现方法
	func AlternateOutput(n, step int)
	n 为交替输出轮次
	step 为每轮输出元素个数
*/

// 参考实现
// 双chan有缓冲
func AlternateOutput1(n, step int) string {
	wg := sync.WaitGroup{}
	num, letter := make(chan bool, 1), make(chan bool)
	res := ""

	wg.Add(2)
	go func() {
		for i := 0; i < n; i++ {
			<-num
			// 引用对象，在方法内发生了扩容，指向了新的底层数组
			res = outputNum(i, step, res)
			// fmt.Printf("i, res: %d %s\n", i, res)
			letter <- true
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < n; i++ {
			<-letter
			res = outputLetter(i, step, res)
			// fmt.Printf("j, res: %d %s\n", i, res)
			num <- true
		}
		wg.Done()
	}()
	// 先输出num
	num <- true
	wg.Wait()

	return res
}

// 双chan无缓冲
func AlternateOutput2(n, step int) string {
	wg := sync.WaitGroup{}
	num, letter := make(chan bool), make(chan bool)
	res := ""

	wg.Add(2)
	go func() {
		for i := 0; i < n; i++ {
			<-num
			// 引用对象，在方法内发生了扩容，指向了新的底层数组
			res = outputNum(i, step, res)
			// fmt.Printf("i, res: %d %s\n", i, res)
			letter <- true
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < n; i++ {
			<-letter
			res = outputLetter(i, step, res)
			// fmt.Printf("j, res: %d %s\n", i, res)
			// 字母输出完就结束了
			if i != n-1 {
				num <- true
			}
		}
		wg.Done()
	}()
	// 先输出num
	num <- true
	wg.Wait()

	return res
}

// 单无缓冲chan
func AlternateOutput3(n, step int) string {
	wg := sync.WaitGroup{}
	ch := make(chan bool)
	res := ""

	wg.Add(2)
	go func() {
		for i := 0; i < n*2; i++ {
			ch <- true
			// 是对方的term就跳过
			if i%2 == 1 {
				continue
			}
			res = outputNum(i/2, step, res)
			// fmt.Printf("i, res: %d %s\n", i, res)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < n*2; i++ {
			<-ch
			// 是对方的term就跳过
			if i%2 == 0 {
				continue
			}
			res = outputLetter(i/2, step, res)
			// fmt.Printf("j, res: %d %s\n", i, res)
		}
		wg.Done()
	}()
	wg.Wait()

	return res
}

func outputNum(i, step int, res string) string {
	// 组装str
	start := 1 + ((i * step) % 10)
	str := ""
	for j := 0; j < step; j++ {
		str += strconv.Itoa(start)
		start = (start + 1) % 10
	}
	res += str

	return res
}

func outputLetter(i, step int, res string) string {
	// 组装str
	start := 'A' + ((i * step) % 26)
	str := []byte{}
	for j := 0; j < step; j++ {
		str = append(str, byte(start))
		start = (start + 1) % 26
	}
	res += string(str)

	return res
}

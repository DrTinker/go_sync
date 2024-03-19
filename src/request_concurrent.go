package src

import (
	"time"
)

/*
	题目来源：腾讯一面
	并发请求多个接口，每个接口获取部分数据，最后正确合并数据，并且控制超时时间
	实现：从monkFunc中获取数组并拼接
	实现方法 RequestConcurrent(n int, expire time.Duration) ([]int, error)
	n 数组上限，从0开始，意味着最终拼接的数组为[0, 1, ... , n-1]
	expire 为超时时间
*/

// 参考实现
// 主要考察context的使用
const part = 10
const defaultExpire = time.Second

func RequestConcurrent(n int, expire time.Duration) ([]int, error) {
	// ctx, cancel := context.WithTimeout(context.TODO(), expire)
	// res := make([]int, n)
	// wg := sync.WaitGroup{}

	// for i := 0; i < n; i += part {
	// 	go func(ctx context.Context) {
	// 		tmp := monkFunc(i, i+part)
	// 		for j := 0; j < part; j++ {
	// 			res[i+j] = tmp[j]
	// 		}
	// 	}(ctx)
	// }
	return nil, nil
}

// 返回从[start, end)的数组
func monkFunc(start, end int) []int {
	return nil
}

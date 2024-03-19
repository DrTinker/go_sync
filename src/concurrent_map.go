package src

import (
	"errors"
	"sync"
	"time"
)

/*
	题目来源：@小徐先生1212 bilibili
	实现一个map，要求：
	1. 并发安全
	2. 只存在插入和查询 O(1)
	3. 查询时，若key存在则返回value，若key不存在则阻塞直到value被放入
	4. 若阻塞超过最大等待时长 maxWaitDuration，则返回超时错误
*/

type ConcurrentMap interface {
	Get(key int, maxWaitDuration time.Duration) (int, error)
	Put(key, val int)
}

// 参考实现
type MyConcurrentMap struct {
	// 保证并发安全的锁
	mu sync.Mutex
	// 基础的map
	baseMap map[int]int
	// 阻塞读取的channel
	waitCh map[int]chan struct{}
}

func NewConcurrentMap() ConcurrentMap {
	return &MyConcurrentMap{
		baseMap: make(map[int]int),
		waitCh:  make(map[int]chan struct{}),
	}
}

func (c *MyConcurrentMap) Get(key int, maxWaitDuration time.Duration) (int, error) {
	// 加锁操作
	c.mu.Lock()
	// 正常读取
	if v, ok := c.baseMap[key]; ok {
		c.mu.Unlock()
		return v, nil
	}
	// 元素不存在需要阻塞
	// 当chan不存在时需要先创建
	if _, ok := c.waitCh[key]; !ok {
		c.waitCh[key] = make(chan struct{})
	}
	// 这里要先解锁，不能抱着锁阻塞
	c.mu.Unlock()
	// 等待chan
	select {
	// k-v被写入
	case <-c.waitCh[key]:
		// 加锁读取
		c.mu.Lock()
		res := c.baseMap[key]
		c.mu.Unlock()
		return res, nil
	// 超时
	case <-time.After(maxWaitDuration):
		return 0, errors.New("get time out")
	}
}

func (c *MyConcurrentMap) Put(key, val int) {
	// 加锁操作
	c.mu.Lock()
	defer c.mu.Unlock()
	// 写入map
	c.baseMap[key] = val
	// 看是否有协程在因为读取本key而阻塞
	ch, ok := c.waitCh[key]
	if !ok {
		return
	}
	// 写入chan等待Get读取
	ch <- struct{}{}
}

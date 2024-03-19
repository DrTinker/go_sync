package src

import "time"

/*
	题目来源：@小徐先生1212 bilibili
	实现一个带有过期自动释放功能的单机锁，要求：
	1. 过期自动释放
	2. 解锁时身份校验
*/

type ExpireLock interface {
	Lock()
	LockWithExpire(expire time.Duration)
	UnLock() error
}

type MyExpireLock struct{}

func (e *MyExpireLock) Lock() {

}
func (e *MyExpireLock) LockWithExpire(expire time.Duration) {

}
func (e *MyExpireLock) UnLock() error {
	return nil
}

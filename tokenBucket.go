package main

import (
	"sync"
	"time"
)

func main() {

}

type tokenBucket struct {
	cap       int           //容量
	limitRate int           //速率 一分钟可以生产放进去多少令牌
	bucket    chan struct{} //放令牌的桶  chan
	lock      *sync.Mutex   //保证线程安全
	stop      bool          //是否停止桶
}

func (t *tokenBucket) NewtokenBucket(cap int, limitRate int) *tokenBucket {
	if cap <= 0 {
		panic("桶大小不能为0")
	}
	return &tokenBucket{
		cap:       cap,
		limitRate: limitRate,
		bucket:    make(chan struct{}, cap),
		lock:      new(sync.Mutex),
	}
}

func (t *tokenBucket) Produce() {
	for {
		t.lock.Lock()
		if t.stop {

			t.lock.Unlock()
			return
		}
		t.bucket <- struct{}{}

		d := time.Minute / time.Duration(t.limitRate) //1个需要几分钟

		t.lock.Unlock()
		time.Sleep(d)
	}
}

func (t *tokenBucket) Consume() {
	<-t.bucket
}

func (t *tokenBucket) Stop() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.stop = true

}

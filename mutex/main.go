package main

import (
	"fmt"
	"sync"
)

func step1() {
	var count = 0
	var wg sync.WaitGroup
	// 需要等待的goroutine数
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

func step2() {
	// 加入互斥锁
	// 注意 这里不需要初始化 Mutex零值表示没有goroutine等待的为加锁状态
	var mu sync.Mutex
	var count = 0
	var wg sync.WaitGroup
	// 需要等待的goroutine数
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// 加锁
				mu.Lock()
				count++
				// 释放锁
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

type Counter struct {
	mu    sync.Mutex
	count int64
}

func step3() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.mu.Lock()
				counter.count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}

func step4()  {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

func (c *Counter) Incr()  {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() int64 {
	// 这里有读操作 因此也需要加锁
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	step4()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	mu   sync.Mutex
	cond *sync.Cond
	data []int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.mu)
	return q
}

// Push 生产者
func (q *Queue) Push(v int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
	// 条件发生变化，队列从空变为非空
	q.cond.Signal()
}

// Pop 消费者，如果队列为空会阻塞等待
func (q *Queue) Pop() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.data) == 0 {
		// 条件不满足，等待
		q.cond.Wait()
	}

	v := q.data[0]
	q.data = q.data[1:] // 删除第一个值，保留后面的值
	return v
}

func main() {
	q := NewQueue()
	// 消费者
	for i := 0; i < 2; i++ {
		go func(id int) {
			for {
				v := q.Pop()
				fmt.Printf("consumer %d got %d\n", id, v)
			}
		}(i)
	}

	// 生产者
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("produce", i)
			q.Push(i)
		}
	}()

	time.Sleep(8 * time.Second)
}

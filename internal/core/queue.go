package core

import (
	"sync"
)

// Queue - очередь на отправку сообщений
type Queue struct {
	ch chan *Email // Канал для отправки сообщений
}

// Push помещает сообщение в очередь
func (q *Queue) Push(m *Email) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		q.ch <- m
		wg.Done()
	}()

	wg.Wait()
	return nil
}

// NewQueue создает новую очередь сообщений размером n
func NewQueue(n int) *Queue {

	q := &Queue{
		ch: make(chan *Email, n),
	}
	return q
}

// PopWait извлекает сообщение из очереди на отправку
func (q *Queue) PopWait() *Email {
	select {
	case val := <-q.ch:
		return val
	default:
		return nil
	}
}

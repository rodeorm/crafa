package queue

import "money/internal/core"

// Queue - очередь на отправку сообщений
type MessageQueue struct {
	ch chan *core.Message // Канал для отправки сообщений
}

// Push помещает сообщение в очередь
func (q *MessageQueue) Push(m *core.Message) {
	q.ch <- m
}

// Len возвращает количество сообщений в очереди
func (q *MessageQueue) Len() int {
	return len(q.ch)
}

// NewQueue создает новую очередь сообщений размером n
func NewQueue(n int) *MessageQueue {
	q := &MessageQueue{
		ch: make(chan *core.Message, n),
	}
	return q
}

// PopWait извлекает сообщение из очереди на отправку
func (q *MessageQueue) PopWait() *core.Message {
	select {
	case val := <-q.ch:
		return val
	default:
		return nil
	}
}

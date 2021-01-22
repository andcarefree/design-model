package singleton

import "sync"

type Message struct {
	Count int
}

type messagePool struct {
	pool *sync.Pool
}

var msgPool *messagePool
var once = &sync.Once{}

func Instance() *messagePool {
	once.Do(func() {
		msgPool = &messagePool{
			pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
		}
	})
	return msgPool
}

func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}

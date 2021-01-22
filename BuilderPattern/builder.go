package BuilderPattern

import "sync"

type builder struct {
	once *sync.Once
	msg  *Message
}

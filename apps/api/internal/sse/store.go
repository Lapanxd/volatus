package sse

import "sync"

type SSEEvent struct {
	EventType string
	Payload   any
}

var (
	connections = make(map[uint]chan SSEEvent)
	mutex       = sync.Mutex{}
)

func Register(userID uint, ch chan SSEEvent) {
	mutex.Lock()
	defer mutex.Unlock()
	connections[userID] = ch
}

func UnRegister(userID uint) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(connections, userID)
}

func SendEvent(userID uint, event SSEEvent) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if ch, ok := connections[userID]; ok {
		select {
		case ch <- event:
			return true
		default:
			return false
		}
	}
	return false
}

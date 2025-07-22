package storage

import (
	"sync"
	"time"
)

type Paste struct {
	Content   string
	ExpiresAt time.Time
	Quote     string
}

var (
	pasteStore = make(map[string]Paste)
	mu         sync.RWMutex
)

func SavePaste(id, content, quote string, duration time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	pasteStore[id] = Paste{
		Content:   content,
		Quote:     quote,
		ExpiresAt: time.Now().Add(duration),
	}
}

func GetPaste(id string) (Paste, bool) {
	mu.RLock()
	defer mu.RUnlock()
	paste, ok := pasteStore[id]

	if !ok || time.Now().After(paste.ExpiresAt) {
		delete(pasteStore, id)
		return Paste{}, false
	}
	return paste, true
}

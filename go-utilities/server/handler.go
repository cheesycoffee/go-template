package server

import "context"

// HandlerFunc types
type HandlerFunc func(ctx context.Context, message []byte) error

// ErrorHandler types
type ErrorHandler func(ctx context.Context, workerType Worker, workerName string, message []byte, err error)

// HandlerGroup group of worker handlers by pattern string
type HandlerGroup struct {
	Handlers []struct {
		Pattern      string
		HandlerFunc  HandlerFunc
		ErrorHandler []ErrorHandler
	}
}

// Add method from HandlerGroup, pattern can contains unique topic name, key, and task name
func (m *HandlerGroup) Add(pattern string, handlerFunc HandlerFunc, errHandlers ...ErrorHandler) {
	m.Handlers = append(m.Handlers, struct {
		Pattern      string
		HandlerFunc  HandlerFunc
		ErrorHandler []ErrorHandler
	}{
		Pattern: pattern, HandlerFunc: handlerFunc, ErrorHandler: errHandlers,
	})
}

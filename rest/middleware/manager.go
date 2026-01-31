package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{globalMiddlewares: []Middleware{}}
}
func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		next = middleware(next)
	}
	return next
}

func (m *Manager) WrapMux(next http.Handler) http.Handler {
	for _, middleware := range m.globalMiddlewares {
		next = middleware(next)
	}
	return next
}

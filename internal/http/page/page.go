package page

import "money/internal/core"

//Page - набор атрибутов для страницы
type Page struct {
	Attributes map[string]any
	Signals    map[string]string
	Session    *core.Session
}

// NewPage создает новую страницу с набором функциональных опций
func NewPage(opts ...func(*Page)) *Page {
	p := &Page{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// WithAttrs добавляет к странице мап атрибутов через замыкание
func WithAttrs(a map[string]any) func(*Page) {
	return func(p *Page) {
		p.Attributes = a
	}
}

// WithSignals добавляет к странице мап сигналов через замыкание
func WithSignals(s map[string]string) func(*Page) {
	return func(p *Page) {
		p.Signals = s
	}
}

func WithSession(s *core.Session) func(*Page) {
	return func(p *Page) {
		p.Session = s
	}
}

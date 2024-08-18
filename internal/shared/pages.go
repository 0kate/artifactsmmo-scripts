package shared

import (
	"math"
)

type Pages[T any] struct {
	data  []T
	total int
	page  int
	size  int
	pages int
}

func NewPages[T any](data []T, total, page, size, pages int) *Pages[T] {
	return &Pages[T]{
		data:  data,
		total: total,
		page:  page,
		size:  size,
		pages: pages,
	}
}

func NewDefaultPages[T any](page, size int) *Pages[T] {
	return NewPages[T](nil, 0, page, size, math.MaxInt32)
}

func (p *Pages[T]) Data() []T {
	return p.data
}

func (p *Pages[T]) Total() int {
	return p.total
}

func (p *Pages[T]) Page() int {
	return p.page
}

func (p *Pages[T]) Size() int {
	return p.size
}

func (p *Pages[T]) Pages() int {
	return p.pages
}

func (p *Pages[T]) HasNext() bool {
	return p.page < p.pages
}

func (p *Pages[T]) Next() *Pages[T] {
	return NewPages[T](nil, p.total, p.page+1, p.size, p.pages)
}

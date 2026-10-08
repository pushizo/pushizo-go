package herosignal

import "context"

// Iter, cursor tabanlı sayfalı listeleri (data / next_cursor / has_more) tek tek dolaşan yineleyicidir. Sayfalar
// ihtiyaç oldukça istenir.
//
//	it := client.Segments.ListAutoPaging(ctx, nil)
//	for it.Next() {
//		fmt.Println(it.Current().ID)
//	}
//	if err := it.Err(); err != nil { … }
type Iter[T any] struct {
	ctx    context.Context
	fetch  func(ctx context.Context, cursor string) ([]T, string, bool, error)
	cursor string
	buf    []T
	cur    T
	done   bool
	err    error
}

func newIter[T any](ctx context.Context, start string, fetch func(ctx context.Context, cursor string) ([]T, string, bool, error)) *Iter[T] {
	return &Iter[T]{ctx: ctx, fetch: fetch, cursor: start}
}

// Next, sıradaki öğeye ilerler; öğe yoksa veya hata oluştuysa false döner.
func (it *Iter[T]) Next() bool {
	for len(it.buf) == 0 {
		if it.done || it.err != nil {
			return false
		}
		items, next, hasMore, err := it.fetch(it.ctx, it.cursor)
		if err != nil {
			it.err = err
			return false
		}
		it.buf = items
		it.cursor = next
		if !hasMore || next == "" {
			it.done = true
		}
	}
	it.cur = it.buf[0]
	it.buf = it.buf[1:]
	return true
}

// Current, Next'in true döndürdüğü son öğedir.
func (it *Iter[T]) Current() T { return it.cur }

// Err, dolaşmayı durduran hatadır (varsa).
func (it *Iter[T]) Err() error { return it.err }

// Collect, kalan tüm öğeleri bir dilimde toplar.
func (it *Iter[T]) Collect() ([]T, error) {
	var out []T
	for it.Next() {
		out = append(out, it.Current())
	}
	return out, it.Err()
}

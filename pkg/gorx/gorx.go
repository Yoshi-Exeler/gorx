package gorx

import "sync"

type Observable[T any] struct {
	value    T
	handlers []func(value T)
	mutex    *sync.Mutex
}

func NewObservable[T any](value T) *Observable[T] {
	return &Observable[T]{
		value:    value,
		handlers: []func(value T){},
		mutex:    &sync.Mutex{},
	}
}

func (o *Observable[T]) Set(value T) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.value = value
	o.propagate(value)
}

func (o *Observable[T]) propagate(value T) {
	for _, handleFunc := range o.handlers {
		go handleFunc(value)
	}
}

func (o *Observable[T]) Get() T {
	return o.value
}

func (o Observable[T]) removeHandler(index int) {
	o.handlers = append(o.handlers[:index], o.handlers[index+1:]...)
}

func (o *Observable[T]) Subscribe(handler func(value T)) *Subscription[T] {
	o.mutex.Lock()
	idx := len(o.handlers) + 1
	o.handlers = append(o.handlers, handler)
	go handler(o.value)
	o.mutex.Unlock()
	return &Subscription[T]{
		target: o,
		index:  idx,
		once:   &sync.Once{},
	}
}

type Subscription[T any] struct {
	target *Observable[T]
	index  int
	once   *sync.Once
}

func (s *Subscription[T]) Unsubscribe() {
	s.once.Do(func() {
		s.target.removeHandler(s.index)
	})
}

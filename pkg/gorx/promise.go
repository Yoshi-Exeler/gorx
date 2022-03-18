package gorx

import "sync"

type Promise[T any] struct {
	resolveHandlers []func(value T)
	rejectHandlers  []func()
	mutex           *sync.Mutex
}

func NewPromise[T any]() *Promise[T] {
	return &Promise[T]{
		resolveHandlers: []func(value T){},
		rejectHandlers:  []func(){},
		mutex:           &sync.Mutex{},
	}
}

func (p *Promise[T]) Resolve(value T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, handleFunc := range p.resolveHandlers {
		go handleFunc(value)
	}
}

func (p *Promise[T]) Reject() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, handleFunc := range p.rejectHandlers {
		go handleFunc()
	}
}

// Then chains a onResolve and a onReject callback to the promise, returning a new dependant promise
func (p *Promise[T]) Then(onResolve func(value T), onReject func()) *Promise[T] {
	child := NewPromise[T]()
	p.mutex.Lock()
	p.resolveHandlers = append(p.resolveHandlers, func(value T) {
		// call the literal handler
		onResolve(value)
		// continue the chain
		child.Resolve(value)
	})
	p.rejectHandlers = append(p.rejectHandlers, func() {
		// call the literal handler
		onReject()
		// continue the chain
		child.Reject()
	})
	p.mutex.Unlock()
	return child
}

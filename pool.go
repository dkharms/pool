package pool

import (
	"errors"
	"sync"
)

type Task[T any] interface {
	Execute() (T, error)
	RetryAmount() int
}

type taskWrapper[T any] struct {
	Task[T]

	result T
	err    error

	done chan struct{}
}

func newTaskWrapper[T any](t Task[T]) *taskWrapper[T] {
	return &taskWrapper[T]{Task: t, done: make(chan struct{}, 1)}
}

func (t *taskWrapper[T]) Finish(v T, err error) {
	t.result, t.err = v, err
	t.done <- struct{}{}
	close(t.done)
}

func (t *taskWrapper[T]) Wait() {
	_ = <-t.done
}

func (t *taskWrapper[T]) Result() T {
	t.Wait()
	return t.result
}

func (t *taskWrapper[T]) Error() error {
	t.Wait()
	return t.err
}

type pool[T any] struct {
	qCapacity int
	wCount    int

	once     sync.Once
	inFlight sync.WaitGroup
	enqueued chan *taskWrapper[T]
}

func New[T any](wCount, qCapacity int) *pool[T] {
	if wCount <= 0 || qCapacity <= 0 {
		panic("invalid opts for pool")
	}

	return &pool[T]{
		qCapacity: qCapacity,
		wCount:    wCount,
		enqueued:  make(chan *taskWrapper[T], qCapacity),
	}
}

func (p *pool[T]) Init() {
	for i := 0; i < p.wCount; i++ {
		p.inFlight.Add(1)

		go func() {
			defer p.inFlight.Done()
			for t := range p.enqueued {
				for r := 0; ; r++ {
					result, err := t.Execute()
					if err == nil || r == t.RetryAmount() {
						t.Finish(result, err)
						break
					}
				}
			}
		}()

	}
}

func (p *pool[T]) Enqueue(t Task[T]) (*taskWrapper[T], error) {
	tw := newTaskWrapper(t)
	select {
	case p.enqueued <- tw:
		return tw, nil
	default:
		return nil, errors.New("queue is fulfilled")
	}
}

func (p *pool[T]) Close() error {
	p.once.Do(func() {
		close(p.enqueued)
	})
	p.inFlight.Wait()
	return nil
}

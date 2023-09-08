package pool

type taskWrapper[T any] struct {
	t Task[T]

	result T
	err    error

	done chan struct{}
}

func newTaskWrapper[T any](t Task[T]) *taskWrapper[T] {
	return &taskWrapper[T]{t: t, done: make(chan struct{}, 1)}
}

func (t *taskWrapper[T]) finish(v T, err error) {
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

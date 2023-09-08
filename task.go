package pool

type Task[T any] interface {
	Execute() (T, error)
	RetryAmount() int
}

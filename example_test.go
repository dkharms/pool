package pool_test

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/dkharms/pool"
)

var _ pool.Task[int] = (*DummyTask[int])(nil)

type DummyTask[T any] func() (T, error)

func (t *DummyTask[T]) Execute() (T, error) {
	return (*t)()
}

func (*DummyTask[T]) RetryAmount() int {
	return 0
}

func ExampleDummyTask() {
	p := pool.New[int](1, 1)
	defer p.Close()

	p.Init()

	var t DummyTask[int] = func() (int, error) {
		if rand.Int()%2 == 0 {
			return 1, nil
		}
		return 0, errors.New("oopsie, you are not lucky enough")
	}

	tw, err := p.Enqueue(&t)
	if err != nil {
		log.Fatal(err)
	}

	if err := tw.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(tw.Result())
}

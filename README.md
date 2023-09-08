<div align="center">

  <img width="256" height="256" src="https://github.com/dkharms/pool/assets/29202384/0645ee42-4319-4fc5-8d8c-568c2897cf9d">

</div>

<div align="center">

  <a href="">![GitHub](https://img.shields.io/github/license/dkharms/pool)</a>
  <a href="">![Go Report Card](https://goreportcard.com/badge/github.com/dkharms/pool)</a>
  <a href="">![Test Workflow](https://github.com/dkharms/pool/actions/workflows/test.yml/badge.svg)</a>

</div>

### About

Just a worker pool with regulation of the number of simultaneously running tasks and the number of maximum tasks in the queue.

### Features

#### Ability to return calculation results.

You don't have to use closures and manage synchronization by yourself – pool will take care of it.

### How To Start

1. Declare task, which satisfies interface `pool.Task`:
  ```go
  var _ pool.Task[int] = (*DummyTask[int])(nil)

  type DummyTask[T any] func() (T, error)

  func (t *DummyTask[T]) Execute() (T, error) {
    return (*t)()
  }

  func (*DummyTask[T]) RetryAmount() int {
    return 0
  }
  ```

2. Initialize worker pool:
  ```go
  p := pool.New[int](1, 1)
  defer p.Cloe()
  p.Init()
  ```

3. Enqueue you task and wait for result or error:
  ```go
  tw, err := p.Enqueue(&t)
  if err != nil {
      return nil, err
  }
  return tw.Result(), tw.Error()
  ```

### What's Next

Checkout [docs](https://pkg.go.dev/github.com/dkharms/pool).


package executor

import (
	"context"
	"errors"
	"os/exec"
	"sync"
	"time"
)

var (
	ErrClosed       = errors.New("executor is closed")
	errEmptyCommand = errors.New("command cannot be empty")
)

type Runner struct {
	ctx       context.Context
	cancel    context.CancelFunc
	waitDelay time.Duration

	mu     sync.Mutex
	closed bool
	wg     sync.WaitGroup
}

func New(parent context.Context, waitDelay time.Duration) *Runner {
	ctx, cancel := context.WithCancel(parent)

	return &Runner{
		ctx:       ctx,
		cancel:    cancel,
		waitDelay: waitDelay,
	}
}

func (r *Runner) Run(command, env []string) ([]byte, error) {
	if len(command) == 0 {
		return nil, errEmptyCommand
	}

	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil, ErrClosed
	}
	if err := r.ctx.Err(); err != nil {
		r.mu.Unlock()
		return nil, err
	}
	r.wg.Add(1)
	r.mu.Unlock()

	defer r.wg.Done()

	c := exec.CommandContext(r.ctx, command[0], command[1:]...)
	c.Env = env
	c.WaitDelay = r.waitDelay

	return c.CombinedOutput()
}

func (r *Runner) Shutdown() {
	r.mu.Lock()
	if !r.closed {
		r.closed = true
		r.cancel()
	}
	r.mu.Unlock()

	r.wg.Wait()
}

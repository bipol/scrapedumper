package circuitbreaker

import (
	"time"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

type CircuitState int

const (
	Closed   CircuitState = 0
	Open     CircuitState = 1
	HalfOpen CircuitState = 2
)

type BooleanRollingWindow struct {
	Vals []bool
	size int
}

func NewBooleanWindow(size int) *BooleanRollingWindow {
	vals := make([]bool, size)
	return &BooleanRollingWindow{
		vals,
		size,
	}
}

func (r *BooleanRollingWindow) Add(x bool) {
	if len(r.Vals) >= r.size {
		r.Vals = r.Vals[1:]
	}

	r.Vals = append(r.Vals, x)
}

func (r *BooleanRollingWindow) All(x bool) bool {
	for _, y := range r.Vals {
		if x != y {
			return false
		}
	}
	return true
}

var (
	//ErrOpenCircuit is an error returned when the circuit is opened
	ErrOpenCircuit = errors.New("circuit is open")
	//ErrSystemFailure is returned when a half open state exceeds the error window
	ErrSystemFailure = errors.New("poor recovery - half open state reverted back to failure")
)

type CircuitBreaker struct {
	state    CircuitState
	window   *BooleanRollingWindow
	openedAt time.Time
	waitTime time.Duration
	logger   *zap.Logger
}

// New will initialize a new circuit breaker, which will "OPEN" whenever it reaches the window val
//after being opened, we will wait until the waitTime has expired before going into a "HALFOPEN" state, allowing whatever service we are hitting to recover.
//if we get enough successful requests at this point to be the window size, we'll go back to an "OPEN" state
func New(logger *zap.Logger, waitTime time.Duration, window int) *CircuitBreaker {
	return &CircuitBreaker{
		state:    Closed,
		window:   NewBooleanWindow(window),
		waitTime: waitTime,
		logger:   logger,
	}
}

// Run will run a given function, and record if there is an error.  It will control its state
//based on values given in the New function
func (c *CircuitBreaker) Run(cmd func() error) error {
	if c.state == Open {
		// If enough time has passed, go to a safety state
		if c.openedAt.Before(time.Now().Add(-c.waitTime)) {
			c.state = HalfOpen
		} else {
			return ErrOpenCircuit
		}
	}

	err := cmd()
	if err != nil {
		c.window.Add(true)
		// if we have exceeded our error threshold, open the circuit
		if c.window.All(true) {
			// if we are at half open, we are in a system failure state
			if c.state == HalfOpen {
				return errors.Wrap(ErrSystemFailure, err.Error())
			}
			c.state = Open
			c.openedAt = time.Now()
			return ErrOpenCircuit
		}
	} else {
		c.window.Add(false)
		// if we are half open, we can revert back to closed if everything is good now
		if c.window.All(false) && c.state == HalfOpen {
			c.state = Closed
		}
	}

	return nil
}

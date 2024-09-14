package circuit_braker

import (
	"sync"
	"time"
)

const (
	StateClosed = iota
	StateOpen
	StateHalfOpen
)

type SingletonCircuitBreaker struct {
	State            int
	FailureCount     int
	FailureThreshold int
	Timeout          time.Duration
	LastFailureTime  time.Time
}

var instance *SingletonCircuitBreaker
var once sync.Once

func NewCircuitBreaker(failureThreshold int, timeout time.Duration) {
	once.Do(func() {
		instance = &SingletonCircuitBreaker{
			FailureThreshold: failureThreshold,
			Timeout:          timeout,
		}
	})
}

func GetInstance() *SingletonCircuitBreaker {
	return instance
}

func SetValueState(value int) {
	singleton := GetInstance()
	singleton.State = value
}

func SetValueFailureCount(value int) {
	singleton := GetInstance()
	singleton.FailureCount = value
}
func SetValueLastFailureTime(value time.Time) {
	singleton := GetInstance()
	singleton.LastFailureTime = value
}

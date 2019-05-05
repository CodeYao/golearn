package scheduler

import "github.com/golearn/crawler/engine"

type SimpleScheduler struct {
	workerChain chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func() {
		s.workerChain <- r
	}()
}

func (s *SimpleScheduler) Run() {
	s.workerChain = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(r chan engine.Request) {
}
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChain
}

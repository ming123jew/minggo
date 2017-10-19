
//分布式后台任务队列模拟(一)
//author: Xiong Chuan Liang
//date: 2015-3-24

package main


import (
"fmt"
"runtime"
"sync"
"time"
)
type workerFunc func(string, ...interface{}) error

type Workers struct {
	workers map[string]workerFunc
}
type OrdType int

const (
	PARALLEL = 1 << iota
	ORDER
)
type JobServer struct {
	Workers
	JobQueue []*WorkerClass
	interval time.Duration
	mt       sync.Mutex
	ord      OrdType
}

func NewJobServer() *JobServer {
	s := &JobServer{}
	s.workers = make(map[string]workerFunc, 0)
	return s
}

func (s *JobServer) RegisterWorkerClass(className string, f workerFunc) int {
	if _, found := s.workers[className]; found {
		return 1
	}
	s.workers[className] = f
	return 0
}

type WorkerClass struct {
	ClassName string
	Args      []interface{}
}

func (s *JobServer) Enqueue(className string, args ...interface{}) bool {
	s.mt.Lock()
	w := &WorkerClass{className, args}
	s.JobQueue = append(s.JobQueue, w)
	s.mt.Unlock()
	return true
}

//poller
func (s *JobServer) poll(quit <-chan bool) <-chan *WorkerClass {
	jobs := make(chan *WorkerClass)

	go func() {
		defer close(jobs)
		for {
			switch {
			case s.JobQueue == nil:
				timeout := time.After(time.Second * 2)
				select {
				case <-quit:
					fmt.Println("[JobServer] [poll] quit")
					return
				case <-timeout:
					fmt.Println("[JobServer] [poll] polling")
				}
			default:

				s.mt.Lock()
				j := s.JobQueue[0]
				if len(s.JobQueue)-1 <= 0 {
					s.JobQueue = nil
				} else {
					s.JobQueue = s.JobQueue[1:len(s.JobQueue)]
				}
				s.mt.Unlock()

				select {
				case jobs <- j:
				case <-quit:
					fmt.Println("[JobServer] [poll] quit")
					return
				}

			}
		}
	}()
	return jobs
}

//worker
func (s *JobServer) work(id int, jobs <-chan *WorkerClass, monitor *sync.WaitGroup) {
	monitor.Add(1)

	f := func() {
		defer monitor.Done()
		for job := range jobs {
			if f, found := s.workers[job.ClassName]; found {
				s.run(f, job)
			} else {
				fmt.Println("[JobServer] [poll] ", job.ClassName, " not found")
			}
		}
	}

	switch s.ord {
	case ORDER:
		f()
	default:
		go f()
	}
}

func (s *JobServer) run(f workerFunc, w *WorkerClass) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[JobServer] [run] Panicking %s\n", fmt.Sprint(r))
		}
	}()

	f(w.ClassName, w.Args...)
}

func (s *JobServer) StartServer(interval time.Duration, ord OrdType) {

	s.interval = interval
	s.ord = ord

	quit := signals()
	jobs := s.poll(quit)

	var monitor sync.WaitGroup

	switch s.ord {
	case ORDER: //顺序执行
		s.work(0, jobs, &monitor)
	default: //并发执行
		concurrency := runtime.NumCPU()
		for id := 0; id < concurrency; id++ {
			s.work(id, jobs, &monitor)
		}
	}

	monitor.Wait()
}
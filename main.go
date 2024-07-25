package main

import (
	"fmt"
	"time"
)

type Job interface {
	Execute()
}

type PrintJob struct {
	Message string
}

func (p PrintJob) Execute() {
	fmt.Println(p.Message)
}

type JobScheduler struct {
	JobQueue chan Job
	Interval time.Duration
}

func NewJobScheduler(interval time.Duration) *JobScheduler {
	return &JobScheduler{
		JobQueue: make(chan Job),
		Interval: interval,
	}
}

func (s *JobScheduler) Start() {
	go func() {
		ticker := time.NewTicker(s.Interval)

		for {
			select {
			case job := <-s.JobQueue:
				job.Execute()
			case <-ticker.C:
				for job := range s.JobQueue {
					job.Execute()
				}
			}
		}
	}()
}

func (s *JobScheduler) ScheduleOnce(duration time.Duration, job Job) {
	go func() {
		time.Sleep(duration)
		s.JobQueue <- job
	}()
}

func main() {
	jobScheduler := NewJobScheduler(1 * time.Minute)
	jobScheduler.Start()

	job := PrintJob{Message: "Hello, World!"}
	jobScheduler.ScheduleOnce(5*time.Second, job)

	job2 := PrintJob{Message: "Hello, World!"}
	go jobScheduler.ScheduleOnce(10*time.Second, job2)

	time.Sleep(15 * time.Second)
} 
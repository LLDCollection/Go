package scheduler

import (
	"time"
)

type Scheduler struct {
	workers chan struct{}
	tasks []*Task
}

func NewScheduler(noOfWorkers int) *Scheduler {
	workers := make(chan struct{}, noOfWorkers)
	return &Scheduler{
		workers: workers,
	}
}

func (s *Scheduler) Schedule(task *Task) {
	task.UpdateStatus(Scheduled)
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) Execute() {
	for {
		now := time.Now()
		for _, t := range s.tasks {
			if t.status != Scheduled { continue }
			if t.executionTime.Before(now) || t.executionTime.Equal(now) {
				s.workers <- struct{}{}
				go func (t *Task)  {
					defer func() { <-s.workers }()
					
					t.UpdateStatus(Executing)

					t.Run()

					t.UpdateStatus(Completed)

					if t.recurring {
						t.executionTime = time.Now().Add(time.Duration(t.intervalInMS) * time.Millisecond)
						s.Schedule(t)
					}
				}(t)
			}
		}
	}
}
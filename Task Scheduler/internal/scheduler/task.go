package scheduler

import "time"

type Runnable func()

type Status string
const (
	Created   Status = "CREATED"
	Scheduled Status = "SCHEDULED"
	Executing Status = "EXECUTING"
	Completed Status = "COMPLETED"
)

type Task struct {
	id            string
	executionTime time.Time
	recurring     bool
	intervalInMS  int
	api           Runnable
	status        Status
}

func NewTask(id string, executionTime time.Time, recurring bool, intervalInMs int, api Runnable) *Task {
	return &Task{
		id: id,
		executionTime: executionTime,
		recurring: recurring,
		intervalInMS: intervalInMs,
		api: api,
		status: Created,
	}
}

func (t *Task) UpdateStatus(status Status) {
	t.status = status
}

func (t *Task) Run() {
	t.api()
}

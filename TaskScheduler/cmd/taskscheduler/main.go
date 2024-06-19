package main

import (
	"fmt"
	"taskscheduler/internal/scheduler"
	"time"
)

func main() {
	sch := scheduler.NewScheduler(3)

	task1 := scheduler.NewTask("1", time.Now().Add(5 * time.Second), false, 0, dummyTask("1"))
	sch.Schedule(task1)

	task2 := scheduler.NewTask("2", time.Now().Add(2 * time.Second), false, 0, dummyTask("2"))
	sch.Schedule(task2)

	task3 := scheduler.NewTask("3", time.Now().Add(2 * time.Second), true, 1, dummyTask("3"))
	sch.Schedule(task3)

	task4 := scheduler.NewTask("4", time.Now().Add(3 * time.Second), false, 0, dummyTask("4"))
	sch.Schedule(task4)

	sch.Execute()
}

func dummyTask(id string) scheduler.Runnable {
	return func() {
		fmt.Printf("Executing task: %v\n", id)
		time.Sleep(5 * time.Second)
	}
}
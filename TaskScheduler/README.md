# Task Scheduler Design Problem

## Problem Statement
Design and implement a multi-threaded task scheduler that can manage both one-time and recurring tasks using multiple worker threads. The task scheduler should allow jobs to be scheduled either at a specific time or at regular intervals. 

### Functional Requirements
1. **Task Types:**
   - **One-time tasks:** Execute exactly once at a specified time.
   - **Recurring tasks:** Execute repeatedly at regular intervals starting from a specified time.

2. **Worker Threads:**
   - Initialize the scheduler with a specified number of worker threads.
   - Maintain a pool of threads that pick up and execute tasks as they become due.

3. **Task Execution Priority:**
   - When multiple tasks are due and there are insufficient idle threads, prioritize tasks based on the nearest next execution time.
   - Tasks with the same next execution time can be executed in any order.

### Input Specification
The input to your program will be a series of commands in the following format:
- `INIT <number_of_threads>`: Initialize the scheduler with the given number of worker threads.
- `ONCE <task_id> <time> <duration>`: Schedule a one-time task with a unique ID that runs for a specified duration starting at the given time.
- `RECURRING <task_id> <start_time> <interval> <duration>`: Schedule a recurring task with a unique ID that starts at the given time and recurs every interval, running for a specified duration each time.

### Output Specification
For each task that starts execution, output should be:
- `EXECUTED <task_id> <actual_start_time>`

### Example Input

```
INIT 3
ONCE 1 15:00 10
RECURRING 2 14:00 10 5
ONCE 3 15:05 5

EXECUTED 1 15:00
EXECUTED 2 15:00
EXECUTED 3 15:05
```

### Constraints
- Assume all times are in `HH:MM` format and represent minutes from the start of the day.
- Assume the system clock provides accurate timing services.
- Efficiently handle a large number of tasks that become due simultaneously.

## Evaluation Criteria
1. **Correctness:** The implementation should correctly interpret commands, schedule, and execute tasks as specified.
2. **Efficiency:** Effective use of threads to handle high volumes of tasks with minimal overhead.
3. **Robustness:** Handling edge cases such as tasks scheduled in the past or tasks with durations longer than their recurrence intervals.

### Design Suggestions
- Use a priority queue to manage tasks, sorting them by their next scheduled execution time.
- Worker threads act as consumers, pulling tasks from the queue and executing them.
- Consider thread safety and synchronization mechanisms to manage concurrency effectively.

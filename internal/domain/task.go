package domain

import "time"

type Task struct {
	ID        TaskID
	StoryID   StoryID
	Title     string
	Body      string
	Status    TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskStatus string

const (
	TaskQueued  TaskStatus = "queued"
	TaskRunning TaskStatus = "running"
	TaskReview  TaskStatus = "review"
	TaskMerged  TaskStatus = "merged"
	TaskFailed  TaskStatus = "failed"
)

type TaskTransition struct {
	From TaskStatus
	To   []TaskStatus
}

var TaskTransitions = map[TaskStatus][]TaskStatus{
	TaskQueued:  {TaskRunning, TaskFailed},
	TaskRunning: {TaskReview, TaskFailed},
	TaskReview:  {TaskMerged, TaskQueued, TaskFailed},
	TaskMerged:  {},
	TaskFailed:  {TaskQueued},
}

func ValidTransition(from, to TaskStatus) bool {
	for _, t := range TaskTransitions[from] {
		if t == to {
			return true
		}
	}
	return false
}

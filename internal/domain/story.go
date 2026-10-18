package domain

import "time"

type StoryID string
type TaskID string
type RunID string

type Story struct {
	ID        StoryID
	Title     string
	Body      string
	Status    StoryStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StoryStatus string

const (
	StoryDraft     StoryStatus = "draft"
	StoryTriaged   StoryStatus = "triaged"
	StoryPlanned   StoryStatus = "planned"
	StoryDone      StoryStatus = "done"
	StoryCancelled StoryStatus = "cancelled"
)

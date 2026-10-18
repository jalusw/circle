package domain

import (
	"encoding/json"
	"time"
)

type EventKind string

const (
	EventRunStarted   EventKind = "run_started"
	EventRunFinished  EventKind = "run_finished"
	EventToolCall     EventKind = "tool_call"
	EventToolResult   EventKind = "tool_result"
	EventAssistantMsg EventKind = "assistant_msg"
	EventPatch        EventKind = "patch"
	EventError        EventKind = "error"
)

type Event struct {
	ID      int64
	RunID   RunID
	Kind    EventKind
	Payload json.RawMessage
	TS      time.Time
}

type RunSpec struct {
	StoryID  StoryID
	TaskID   TaskID
	Runner   string
	Title    string
	Prompt   string
	MaxRetry int
}

type RunResult struct {
	Success bool
	Diff    string
	Summary string
}

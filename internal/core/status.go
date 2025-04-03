package core

import (
	"time"
)

type StatusHistoryRecord struct {
	Modifier       User
	PreviousStatus Status
	NewStatus      Status
	Time           time.Time
}

type Status struct {
	Level    Level
	Name     string
	ID       int
	Parents  []Status
	Children []Status
}

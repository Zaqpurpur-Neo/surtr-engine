package commons

type StatusProject string

const (
	StatusProjectActive    StatusProject = "active"
	StatusProjectPaused    StatusProject = "paused"
	StatusProjectCompleted StatusProject = "completed"
	StatusProjectCancelled StatusProject = "cancelled"
)

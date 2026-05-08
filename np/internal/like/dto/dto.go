package dto

type CountRequest struct {
	UserID  int
	Targets []Target
}

type Target struct {
	TargetID   int
	TargetType string
}

type Reaction struct {
	UserID     int
	TargetID   int
	TargetType string
}

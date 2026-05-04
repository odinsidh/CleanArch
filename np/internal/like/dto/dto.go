package dto

type CountInput struct {
	TargetID   int
	TargetType string
}

type CountOutput struct {
	TargetID   int
	TargetType string
	Count      map[string]int
}

type Reaction struct {
	UserID     int
	TargetID   int
	TargetType string
}

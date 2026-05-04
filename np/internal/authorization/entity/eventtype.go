package entity

type EventType int

const (
	Grant EventType = iota + 1
	Revoke
)

package entity

import "fmt"

// =====================================================================
//  Username errors
// =====================================================================

var (
	ErrUsernameNotFound error = fmt.Errorf("username not found")
)

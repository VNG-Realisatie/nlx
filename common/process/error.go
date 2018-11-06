package process

import (
	"errors"
)

var (
	// ErrMaxLevelReached is thrown when trying to add a process to a level above the max level, max level is defined when creating a new process
	ErrMaxLevelReached = errors.New("maximum number of levels reached")
)

package process

import (
	"errors"
	"strconv"
)

var (
	ErrMaxLevelReached = errors.New("maximum supported level is " + strconv.Itoa(maxLevel))
)

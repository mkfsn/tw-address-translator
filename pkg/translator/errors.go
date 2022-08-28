package translator

import (
	"errors"
)

var (
	ErrNoCounty = errors.New("no country")
	ErrNoRoad   = errors.New("no road")
)

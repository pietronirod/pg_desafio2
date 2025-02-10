package types

import "time"

type Result struct {
	StatusCode int
	Duration   time.Duration
}

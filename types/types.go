package types

import (
	"time"
)

func String(a string) *string {
	return &a
}

func Int(a int) *int {
	return &a
}

func Bool(a bool) *bool {
	return &a
}

func Time(a time.Time) *time.Time {
	return &a
}

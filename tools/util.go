package util

import "time"

func Elapsed(startTime time.Time) int64 {
	end := time.Now().UnixNano() / int64(time.Millisecond)
	start := startTime.UnixNano() / int64(time.Millisecond)
	diff := end - start
	return diff
}
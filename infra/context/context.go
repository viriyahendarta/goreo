package context

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

const (
	//StartTime is constant for start time context
	StartTime contextKey = "start_time"
)

//GetElapsedTime returns elapsed time during execution
func GetElapsedTime(ctx context.Context) string {
	if ctx != nil {
		if start, ok := ctx.Value(StartTime).(time.Time); ok {
			elapsed := time.Since(start).Seconds() * 1000
			return fmt.Sprintf("%.2fms", elapsed)
		}
	}
	return "-1ms"
}

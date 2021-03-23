package utils

import "context"

// ContextDone returns true if we can exit due to a cancelled context.
func ContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

package utils

import (
	"context"

	"go.uber.org/zap"
)

// ContextDone returns true if we can exit due to a cancelled context.
func ContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// LogContextDone is like ContextDone but logs when the context is done
func LogContextDone(logger *zap.Logger, ctx context.Context) bool {
	if ContextDone(ctx) {
		logger.Info("context cancelled")
		return true
	} else {
		return false
	}
}

package pipelineinternal

import "context"

// Worker is an excuter of a pipe line
type Worker interface {
	Run(ctx context.Context) (context.Context, error)
}

package main

import (
	"context"
	"generator/internal/entrypoints"
)

func main() {
	context := context.Background()
	entrypoints.SyntheticContinuousOrdersProcessorEntryPoint(context)
	<-context.Done()
}

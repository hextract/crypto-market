package main

import (
	"context"
	"executor/internal/entrypoints"
)

func main() {
	context := context.Background()
	entrypoints.OrderProcessorEntryPoint(context)
	select {}
}

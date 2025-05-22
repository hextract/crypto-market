package main

import (
	"context"
	"reporter/internal/entrypoints"
)

func main() {
	context := context.Background()
	entrypoints.EntryPoint(context)
	select {}
}

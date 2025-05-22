package main

import (
	"context"
	"receiver/internal/entrypoints"
)

func main() {
	context := context.Background()
	entrypoints.WebEntryPoint(context)
}

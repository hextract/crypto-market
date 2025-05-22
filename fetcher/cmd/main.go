package main

import (
	"context"
	"fetcher/internal/entrypoints"
)

func main() {
	context := context.Background()
	entrypoints.BinancePollerEntryPoint(context)
	entrypoints.WebEntryPoint(context)
}

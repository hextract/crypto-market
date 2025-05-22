package entrypoints

import (
	"context"
	"fetcher/internal/presentation/web"
	"log"
)

func WebEntryPoint(context context.Context) {
	web.EntryPoint(context)
	log.Println("Web is started up")
}

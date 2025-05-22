package entrypoints

import (
	"context"
	"log"
	"receiver/internal/presentation/web"
)

func WebEntryPoint(context context.Context) {
	web.EntryPoint(context)
	log.Println("Web is started up")
}

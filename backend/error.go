package backend

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
)

var wailsContext context.Context

const errorEvent = "error"

func BindContext(ctx context.Context) {
	if wailsContext != nil {
		log.Println("Error: context already bound")
		return
	}

	wailsContext = ctx
}

func EmitError() {
	runtime.EventsEmit(wailsContext, errorEvent)
}

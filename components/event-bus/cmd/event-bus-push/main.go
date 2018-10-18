package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mszostok/kyma/components/event-bus/cmd/event-bus-push/application"
	"github.com/mszostok/kyma/components/event-bus/internal/push/opts"
)

func main() {
	pushOpts := opts.ParseFlags()

	pushApplication := application.NewPushApplication(pushOpts)
	defer pushApplication.Stop()

	log.Printf("HTTP server starting on port %v", pushOpts.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", pushOpts.Port), pushApplication.ServerMux))
}

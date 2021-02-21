package events

import (
	"fmt"

	"opensvc.com/opensvc/core/client"
	"opensvc.com/opensvc/core/event"
	"opensvc.com/opensvc/core/output"
)

// Do renders the cluster status
func Do(color string, format string) {
	api := client.New()
	opts := client.NewEventsCmdConfig()
	events, _ := api.Events(*opts)
	defer close(events)
	for event := range events {
		doOne(event, color, format)
	}
}

func doOne(e event.Event, color string, format string) {
	human := func() {
		fmt.Print(event.Render(e))
	}
	output.Switch(format, color, e, human)
}

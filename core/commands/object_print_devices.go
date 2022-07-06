package commands

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"opensvc.com/opensvc/core/client"
	"opensvc.com/opensvc/core/clientcontext"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/object"
	"opensvc.com/opensvc/core/objectdevice"
	"opensvc.com/opensvc/core/output"
	"opensvc.com/opensvc/core/rawconfig"
)

type (
	// CmdObjectPrintDevices is the cobra flag set of the print devices command.
	CmdObjectPrintDevices struct {
		OptsGlobal
		Roles string `flag:"devroles"`
	}
	devicer interface {
		PrintDevices(roles objectdevice.Role) objectdevice.L
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdObjectPrintDevices) Init(kind string, parent *cobra.Command, selector *string) {
	cmd := t.cmd(kind, selector)
	parent.AddCommand(cmd)
	flag.Install(cmd, t)
}

func (t *CmdObjectPrintDevices) cmd(kind string, selector *string) *cobra.Command {
	return &cobra.Command{
		Use:     "devices",
		Short:   "Print selected objects and resources exposed, used, base and claimed block devices",
		Aliases: []string{"device", "devic", "devi", "dev", "devs", "de"},
		Run: func(cmd *cobra.Command, args []string) {
			t.run(selector, kind)
		},
	}
}

func (t *CmdObjectPrintDevices) extract(selector string, c *client.T) objectdevice.L {
	if t.Local || (t.NodeSelector == "" && !clientcontext.IsSet()) {
		return t.extractLocal(selector)
	}
	if data, err := t.extractFromDaemon(selector, c); err == nil {
		return data
	}
	if clientcontext.IsSet() {
		log.Error().Msg("can not fetch daemon data")
		return objectdevice.NewList()
	}
	return t.extractLocal(selector)
}

func (t *CmdObjectPrintDevices) extractLocal(selector string) objectdevice.L {
	data := objectdevice.NewList()
	sel := object.NewSelection(
		selector,
		object.SelectionWithLocal(true),
	)
	paths, err := sel.Expand()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return data
	}
	for _, p := range paths {
		obj, err := object.New(p)
		if err != nil {
			continue
		}
		i, ok := obj.(devicer)
		if !ok {
			continue
		}
		roles := objectdevice.ParseRoles(t.Roles)
		table := i.PrintDevices(roles)
		data = data.Add(table)
	}
	return data
}

func (t *CmdObjectPrintDevices) extractFromDaemon(selector string, c *client.T) (objectdevice.L, error) {
	data := objectdevice.NewList()
	/*
		req := c.NewGetDevicess()
		req.ObjectSelector = selector
		b, err := req.Do()
		if err != nil {
			return data, err
		}
		err = json.Unmarshal(b, &data)
		if err != nil {
			log.Debug().Err(err).Msg("unmarshal GET /schedules")
			return data, err
		}
	*/
	return data, nil
}

func (t *CmdObjectPrintDevices) run(selector *string, kind string) {
	mergedSelector := mergeSelector(*selector, t.ObjectSelector, kind, "")
	c, err := client.New(client.WithURL(t.Server))
	if err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}
	data := t.extract(mergedSelector, c)

	output.Renderer{
		Format:   t.Format,
		Color:    t.Color,
		Data:     data,
		Colorize: rawconfig.Colorize,
		HumanRenderer: func() string {
			return data.Render()
		},
	}.Print()
}

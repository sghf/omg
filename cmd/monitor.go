package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/api/daemon/status"
	"opensvc.com/opensvc/core/api/getevent"
	"opensvc.com/opensvc/core/client"
	"opensvc.com/opensvc/core/entrypoints/monitor"
	"os"
)

var (
	monWatchFlag    bool
	monSelectorFlag string
)

var monCmd = &cobra.Command{
	Use:     "monitor",
	Aliases: []string{"m", "mo", "mon", "moni", "monit", "monito"},
	Short:   "Print the cluster status",
	Long:    monitor.CmdLong,
	Run:     monCmdRun,
}

func init() {
	rootCmd.AddCommand(monCmd)
	monCmd.Flags().StringVarP(&monSelectorFlag, "selector", "s", "*", "An object selector expression")
	monCmd.Flags().BoolVarP(&monWatchFlag, "watch", "w", false, "Watch the monitor changes")
}

func monCmdRun(_ *cobra.Command, _ []string) {
	m := monitor.New()
	m.SetColor(colorFlag)
	m.SetFormat(formatFlag)

	cli, err := client.New().SetURL(serverFlag).Configure()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}
	if monWatchFlag {
		getter := getevent.New(*cli, monSelectorFlag, true)
		if err = m.DoWatch(getter, os.Stdout); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return
		}
	} else {
		getter := status.New(*cli, monSelectorFlag)
		m.Do(getter, os.Stdout)
	}
}

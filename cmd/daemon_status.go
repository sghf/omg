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
	daemonStatusWatchFlag    bool
	daemonStatusSelectorFlag string
)

var daemonStatusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Print the cluster status",
	Long:    monitor.CmdLong,
	Aliases: []string{"statu"},
	Run:     daemonStatusCmdRun,
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
	daemonStatusCmd.Flags().BoolVarP(&daemonStatusWatchFlag, "watch", "w", false, "Watch the monitor changes")
	daemonStatusCmd.Flags().StringVarP(&daemonStatusSelectorFlag, "selector", "s", "**", "Select opensvc objects (ex: **/db*,*/svc/db*)")
}

func daemonStatusCmdRun(cmd *cobra.Command, args []string) {
	m := monitor.New()
	m.SetColor(colorFlag)
	m.SetFormat(formatFlag)

	cli, err := client.New().SetURL(serverFlag).Configure()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if daemonStatusWatchFlag {
		getter := getevent.New(*cli, daemonStatusSelectorFlag, true)
		m.DoWatch(getter, os.Stdout)
	} else {
		getter := status.New(*cli, daemonStatusSelectorFlag)
		m.Do(getter, os.Stdout)
	}
}

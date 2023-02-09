package commands

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/pkg/errors"

	"github.com/opensvc/om3/core/client"
	"github.com/opensvc/om3/core/rawconfig"
	"github.com/opensvc/om3/util/command"
	"github.com/opensvc/om3/util/file"
)

type (
	CmdDaemonCommon struct{}
)

func (t *CmdDaemonCommon) startDaemon() (err error) {
	cmd := command.New(
		command.WithName(os.Args[0]),
		command.WithArgs([]string{"daemon", "start"}),
	)
	_, _ = fmt.Fprintf(os.Stdout, "Start daemon\n")
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "%s", cmd.String())
	}
	return nil
}

func (t *CmdDaemonCommon) stopDaemon() (err error) {
	cmd := command.New(
		command.WithName(os.Args[0]),
		command.WithArgs([]string{"daemon", "stop"}),
	)
	_, _ = fmt.Fprintf(os.Stdout, "Stop daemon\n")
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "%s", cmd.String())
	}
	return nil
}

// isRunning returns true if daemon is running
func (t *CmdDaemonCommon) isRunning() bool {
	cli, err := client.New()
	if err != nil {
		return false
	}
	if _, err := cli.NewGetDaemonRunning().Do(); err == nil {
		return true
	}
	return false
}

func (t *CmdDaemonCommon) nodeDrain() (err error) {
	cmd := command.New(
		command.WithName(os.Args[0]),
		command.WithArgs([]string{"node", "drain", "--wait"}),
	)
	_, _ = fmt.Fprintf(os.Stdout, "Draining node\n")
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "%s", cmd.String())
	}
	return nil
}

func (t *CmdDaemonCommon) backupLocalConfig(name string) (err error) {
	pathEtc := rawconfig.Paths.Etc
	if !file.ExistsAndDir(pathEtc) {
		_, _ = fmt.Fprintf(os.Stdout, "Empty %s, skip backup\n", pathEtc)
		return nil
	}
	cmd := command.New(
		command.WithBufferedStdout(),
		command.WithName(os.Args[0]),
		command.WithArgs([]string{"**", "print", "config", "--local", "--format", "json"}),
		// allow exit code 1 (Error: no match)
		command.WithIgnoredExitCodes(0, 1),
	)
	_, _ = fmt.Fprintln(os.Stdout, "Dump all configs")
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "%s", cmd.String())
	}

	backup := path.Join(pathEtc, name+time.Now().Format(name+"-2006-01-02T15:04:05.json"))
	_, _ = fmt.Fprintf(os.Stdout, "Save configs to %s\n", backup)
	if err := os.WriteFile(backup, cmd.Stdout(), 0o400); err != nil {
		return err
	}
	return nil
}

func (t *CmdDaemonCommon) deleteLocalConfig() (err error) {
	pathEtc := rawconfig.Paths.Etc
	if file.ExistsAndDir(pathEtc) {
		cmd := command.New(
			command.WithName(os.Args[0]),
			command.WithArgs([]string{"**", "delete", "--local"}),
		)
		_, _ = fmt.Fprintf(os.Stdout, "Delete all config\n")
		if err := cmd.Run(); err != nil {
			return errors.Wrapf(err, "%s", cmd.String())
		}
	} else {
		_, _ = fmt.Fprintf(os.Stdout, "Empty %s, skip delete local config\n", pathEtc)
	}
	return rawconfig.CreateMandatoryDirectories()
}

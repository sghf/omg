package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"opensvc.com/opensvc/core/flag"
	"opensvc.com/opensvc/core/object"
	"opensvc.com/opensvc/core/objectaction"
	"opensvc.com/opensvc/core/path"
)

type (
	// CmdKeystoreKeys is the cobra flag set of the keys command.
	CmdKeystoreKeys struct {
		object.OptsKeys
	}
)

// Init configures a cobra command and adds it to the parent command.
func (t *CmdKeystoreKeys) Init(kind string, parent *cobra.Command, selector *string) {
	cmd := t.cmd(kind, selector)
	parent.AddCommand(cmd)
	flag.Install(cmd, &t.OptsKeys)
}

func (t *CmdKeystoreKeys) cmd(kind string, selector *string) *cobra.Command {
	return &cobra.Command{
		Use:   "keys",
		Short: "list the object key names",
		Run: func(cmd *cobra.Command, args []string) {
			t.run(selector, kind)
		},
	}
}

func (t *CmdKeystoreKeys) run(selector *string, kind string) {
	mergedSelector := mergeSelector(*selector, t.ObjectSelector, kind, "")
	objectaction.New(
		objectaction.LocalFirst(),
		objectaction.WithLocal(t.Local),
		objectaction.WithColor(t.Color),
		objectaction.WithFormat(t.Format),
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithRemoteNodes(t.NodeSelector),
		objectaction.WithRemoteAction("keys"),
		objectaction.WithRemoteOptions(map[string]interface{}{
			"match": t.Match,
		}),
		objectaction.WithLocalRun(func(p path.T) (interface{}, error) {
			o, err := object.NewFromPath(p)
			if err != nil {
				return nil, err
			}
			store, ok := o.(object.Keystorer)
			if !ok {
				return nil, fmt.Errorf("%s is not a keystore", o)
			}
			return store.Keys(t.OptsKeys)
		}),
	).Do()
}

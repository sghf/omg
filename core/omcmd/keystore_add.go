package omcmd

import (
	"context"

	"github.com/opensvc/om3/core/naming"
	"github.com/opensvc/om3/core/object"
	"github.com/opensvc/om3/core/objectaction"
	"github.com/opensvc/om3/util/uri"
)

type (
	CmdKeystoreAdd struct {
		OptsGlobal
		OptsLock
		Key   string
		From  string
		Value string
	}
)

func (t *CmdKeystoreAdd) Run(selector, kind string) error {
	mergedSelector := mergeSelector(selector, t.ObjectSelector, kind, "")
	return objectaction.New(
		objectaction.LocalFirst(),
		objectaction.WithLocal(t.Local),
		objectaction.WithColor(t.Color),
		objectaction.WithOutput(t.Output),
		objectaction.WithObjectSelector(mergedSelector),
		objectaction.WithLocalFunc(func(ctx context.Context, p naming.Path) (interface{}, error) {
			store, err := object.NewKeystore(p)
			if err != nil {
				return nil, err
			}
			if t.From != "" {
				b, err := uri.ReadAllFrom(t.From)
				if err != nil {
					return nil, err
				}
				return nil, store.AddKey(t.Key, b)
			}
			return nil, store.AddKey(t.Key, []byte(t.Value))
		}),
	).Do()
}

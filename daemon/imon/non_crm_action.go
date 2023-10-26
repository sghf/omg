package imon

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/opensvc/om3/daemon/msgbus"
	"github.com/opensvc/om3/util/file"
)

func (o *imon) getFrozen() time.Time {
	return file.ModTime(filepath.Join(o.path.VarDir(), "frozen"))
}

// freeze creates missing instance frozen flag file, and publish InstanceFrozenFileUpdated
// local instance status cache frozen value is updated with value read from file system
func (o *imon) freeze() error {
	frozen := o.getFrozen()

	o.log.Infof("daemon action freeze")
	p := filepath.Join(o.path.VarDir(), "frozen")

	if !file.Exists(p) {
		d := filepath.Dir(p)
		if !file.Exists(d) {
			if err := os.MkdirAll(d, os.ModePerm); err != nil {
				o.log.Errorf("freeze: %s", err)
				return err
			}
		}
		f, err := os.Create(p)
		if err != nil {
			o.log.Errorf("freeze: %s", err)
			return err
		}
		_ = f.Close()
	}
	frozen = file.ModTime(p)
	if instanceStatus, ok := o.instStatus[o.localhost]; ok {
		instanceStatus.FrozenAt = frozen
		o.instStatus[o.localhost] = instanceStatus
	}
	if frozen.IsZero() {
		err := fmt.Errorf("unexpected frozen reset on %s", p)
		o.log.Errorf("freeze: %s", err)
		return err
	}
	o.pubsubBus.Pub(&msgbus.InstanceFrozenFileUpdated{Path: o.path, At: frozen},
		o.labelPath,
		o.labelLocalhost,
	)
	return nil
}

// freeze removes instance frozen flag file, and publish InstanceFrozenFileUpdated
// local instance status cache frozen value is updated with value read from file system
func (o *imon) unfreeze() error {
	o.log.Infof("daemon action unfreeze")
	p := filepath.Join(o.path.VarDir(), "frozen")
	if !file.Exists(p) {
		o.log.Infof("already thawed")
	} else {
		err := os.Remove(p)
		if err != nil {
			o.log.Errorf("unfreeze: %s", err)
			return err
		}
	}
	if instanceStatus, ok := o.instStatus[o.localhost]; ok {
		instanceStatus.FrozenAt = time.Time{}
		o.instStatus[o.localhost] = instanceStatus
	}
	o.pubsubBus.Pub(&msgbus.InstanceFrozenFileRemoved{Path: o.path, At: time.Now()},
		o.labelLocalhost,
		o.labelPath,
	)
	return nil
}

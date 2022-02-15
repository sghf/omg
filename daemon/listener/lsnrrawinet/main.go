package lsnrrawinet

import (
	"net"
	"net/http"
	"time"

	"github.com/rs/zerolog"

	"opensvc.com/opensvc/daemon/listener/mux/rawmux"
	"opensvc.com/opensvc/daemon/routinehelper"
	"opensvc.com/opensvc/daemon/subdaemon"
	"opensvc.com/opensvc/util/funcopt"
)

type (
	T struct {
		*subdaemon.T
		routinehelper.TT
		listener     *net.Listener
		log          zerolog.Logger
		routineTrace routineTracer
		mux          rawServer
		httpHandler  http.Handler
		addr         string
	}

	routineTracer interface {
		Trace(string) func()
		Stats() routinehelper.Stat
	}

	rawServer interface {
		Serve(rawmux.ReadWriteCloseSetDeadliner)
	}
)

func New(opts ...funcopt.O) *T {
	t := &T{}
	t.SetTracer(routinehelper.NewTracerNoop())
	if err := funcopt.Apply(t, opts...); err != nil {
		t.log.Error().Err(err).Msg("listener funcopt.Apply")
		return nil
	}
	t.T = subdaemon.New(
		subdaemon.WithName("lsnr-raw-inet"),
		subdaemon.WithMainManager(t),
		subdaemon.WithRoutineTracer(&t.TT),
	)
	t.log = t.Log().With().
		Str("addr", t.addr).
		Str("sub", t.Name()).
		Logger()
	t.mux = rawmux.New(t.httpHandler, t.log, 5*time.Second)
	return t
}

func (t *T) MainStart() error {
	t.log.Debug().Msg("mgr starting")
	started := make(chan bool)
	go func() {
		defer t.Trace(t.Name())()
		if err := t.start(); err != nil {
			t.log.Error().Err(err).Msgf("mgr start failure")
		}
		started <- true
	}()
	<-started
	t.log.Debug().Msg("mgr started")
	return nil
}

func (t *T) MainStop() error {
	t.log.Debug().Msg("mgr stopping")
	if err := t.stop(); err != nil {
		t.log.Error().Err(err).Msg("mgr stop failure")
	}
	t.log.Debug().Msg("mgr stopped")
	return nil
}

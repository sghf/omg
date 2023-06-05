package daemonapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/opensvc/om3/daemon/api"
	"github.com/opensvc/om3/daemon/daemondata"
	"github.com/opensvc/om3/daemon/subdaemon"
	"github.com/opensvc/om3/util/hostname"
	"github.com/opensvc/om3/util/pubsub"
)

type DaemonApi struct {
	Daemon     subdaemon.RootManager
	Daemondata *daemondata.T
	EventBus   *pubsub.Bus
}

var (
	labelApi  = pubsub.Label{"origin", "api"}
	labelNode = pubsub.Label{"node", hostname.Hostname()}
)

func JSONProblem(ctx echo.Context, code int, title, detail string) error {
	return ctx.JSON(code, api.Problem{
		Detail: detail,
		Title:  title,
		Status: code,
	})
}

func JSONProblemf(ctx echo.Context, code int, title, detail string, argv ...any) error {
	return ctx.JSON(code, api.Problem{
		Detail: fmt.Sprintf(detail, argv...),
		Title:  title,
		Status: code,
	})
}

func setStreamHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-control", "no-store")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")
}

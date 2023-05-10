package daemonapi

import (
	"encoding/json"
	"net/http"

	"github.com/opensvc/om3/core/network"
	"github.com/opensvc/om3/core/object"
	"github.com/opensvc/om3/daemon/api"
)

// GetDaemonDNSDump returns the DNS zone content.
func (a *DaemonApi) GetNetworks(w http.ResponseWriter, r *http.Request, params api.GetNetworksParams) {
	n, err := object.NewNode(object.WithVolatile(true))
	if err != nil {
		sendError(w, http.StatusInternalServerError, "new node")
		return
	}
	var l network.StatusList
	if params.Name != nil {
		l = network.ShowNetworksByName(n, *params.Name)
	} else {
		l = network.ShowNetworks(n)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(l)
	w.WriteHeader(http.StatusOK)
}

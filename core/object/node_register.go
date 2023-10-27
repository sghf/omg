package object

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/opensvc/om3/core/keyop"
	"github.com/opensvc/om3/util/hostname"
	"github.com/opensvc/om3/util/key"
	"golang.org/x/term"
)

type (
	// registerReq structures the POST /register request body
	registerReq struct {
		Nodename string `json:"nodename"`
		App      string `json:"app,omitempty"`
	}

	// registerRes structures the POST /register response body
	registerRes struct {
		Data  registerResData `json:"data"`
		Info  string          `json:"info"`
		Error string          `json:"error"`
	}
	registerResData struct {
		UUID string `json:"uuid"`
	}
)

// Register logs in to the collector using the provided user credentials.
//
// If the node is not already known to the collector, a new node uuid is
// generated by the collector and stored by the agent in the node config
// as a node.uuid key.
//
// If user credentials are given, the POST /register collector handler is
// used, and the app code is supported.
// If no user credentials are given, the "register_node" jsonrpc handler
// is used, and app code is ignored.
//
// If app is not set, the node is added to any app under the user's
// responsibility.
func (t Node) Register(user, password, app string) error {
	if err := t.register(user, password, app); err != nil {
		return err
	}
	if _, err := t.PushAsset(); err != nil {
		return err
	} else {
		t.Log().Infof("sent initial asset discovery")
	}
	if data, err := t.Checks(); err != nil {
		return err
	} else {
		t.Log().Infof("sent initial checks (%d)", data.Len())
	}
	if data, err := t.PushPkg(); err != nil {
		return err
	} else {
		t.Log().Infof("sent initial package inventory (%d)", len(data))
	}
	if data, err := t.PushPatch(); err != nil {
		return err
	} else {
		t.Log().Infof("sent initial patch inventory (%d)", len(data))
	}
	if _, err := t.PushDisks(); err != nil {
		return err
	}
	if err := t.Sysreport(); err != nil {
		return err
	}
	return nil
}

func (t Node) register(user, password, app string) error {
	if user == "" {
		return t.registerAsNode()
	} else {
		return t.registerAsUser(user, password, app)
	}
}

func (t Node) registerAsUser(user, password, app string) error {
	if password == "" {
		fmt.Printf("Password: ")
		if b, err := term.ReadPassword(int(os.Stdin.Fd())); err != nil {
			return err
		} else {
			password = string(b)
			fmt.Println("")
		}
	}
	client := t.CollectorRestAPIClient()
	url, err := t.CollectorRestAPIURL()
	if err != nil {
		return err
	}
	url.Path += "/register"
	if app == "" {
		app = t.MergedConfig().GetString(key.Parse("node.app"))
	}
	reqData := registerReq{
		Nodename: hostname.Hostname(),
		App:      app,
	}
	b, err := json.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("encode request body: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(b))
	req.SetBasicAuth(user, password)
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		if b, err := io.ReadAll(response.Body); err != nil {
			return fmt.Errorf("%d: %s", response.StatusCode, response.Status)
		} else {
			return fmt.Errorf("%s", string(b))
		}
	}
	dec := json.NewDecoder(response.Body)
	data := registerRes{}
	if err := dec.Decode(&data); err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	if data.Error != "" {
		return fmt.Errorf(data.Error)
	}
	if data.Info != "" {
		t.Log().Infof("%s", data.Info)
	}
	return t.writeUUID(data.Data.UUID)
}

func (t Node) registerAsNode() error {
	client, err := t.CollectorFeedClient()
	if err != nil {
		return err
	}
	if response, err := client.Call("register_node"); err != nil {
		return err
	} else if response.Error != nil {
		return fmt.Errorf("rpc: %s: %s", response.Error.Message, response.Error.Data)
	} else if response.Result != nil {
		switch v := response.Result.(type) {
		case []interface{}:
			for _, e := range v {
				s, ok := e.(string)
				if !ok {
					continue
				}
				if strings.Contains(s, "already") {
					t.Log().Infof(s)
				} else {
					return fmt.Errorf(s)
				}
			}
		case string:
			return t.writeUUID(v)
		default:
			return fmt.Errorf("unknown response result type: %+v", v)
		}
	} else {
		return fmt.Errorf("unexpected rpc response: %+v", response)
	}
	return nil
}

func (t Node) writeUUID(s string) error {
	if current := t.Config().GetString(key.Parse("node.uuid")); current == s {
		return nil
	}
	op := keyop.Parse("node.uuid=" + s)
	if err := t.Config().Set(*op); err != nil {
		return err
	}
	return t.Config().Commit()
}

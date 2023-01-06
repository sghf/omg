// Package daemonapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package daemonapi

import (
	"time"
)

const (
	BasicAuthScopes  = "basicAuth.Scopes"
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for Orchestrate.
const (
	OrchestrateFalse Orchestrate = "false"
	OrchestrateHa    Orchestrate = "ha"
	OrchestrateStart Orchestrate = "start"
)

// Defines values for Placement.
const (
	PlacementLoadAvg    Placement = "load avg"
	PlacementNodesOrder Placement = "nodes order"
	PlacementNone       Placement = "none"
	PlacementScore      Placement = "score"
	PlacementShift      Placement = "shift"
	PlacementSpread     Placement = "spread"
)

// Defines values for PostDaemonLogsControlLevel.
const (
	PostDaemonLogsControlLevelDebug PostDaemonLogsControlLevel = "debug"
	PostDaemonLogsControlLevelError PostDaemonLogsControlLevel = "error"
	PostDaemonLogsControlLevelFatal PostDaemonLogsControlLevel = "fatal"
	PostDaemonLogsControlLevelInfo  PostDaemonLogsControlLevel = "info"
	PostDaemonLogsControlLevelNone  PostDaemonLogsControlLevel = "none"
	PostDaemonLogsControlLevelPanic PostDaemonLogsControlLevel = "panic"
	PostDaemonLogsControlLevelWarn  PostDaemonLogsControlLevel = "warn"
)

// Defines values for PostDaemonSubActionAction.
const (
	PostDaemonSubActionActionStart PostDaemonSubActionAction = "start"
	PostDaemonSubActionActionStop  PostDaemonSubActionAction = "stop"
)

// Defines values for Provisioned.
const (
	False Provisioned = "false"
	Mixed Provisioned = "mixed"
	Na    Provisioned = "n/a"
	True  Provisioned = "true"
)

// Defines values for Topology.
const (
	Failover Topology = "failover"
	Flex     Topology = "flex"
)

// App defines model for app.
type App = string

// Cluster defines model for cluster.
type Cluster struct {
	Config ClusterConfig `json:"config"`
	Node   ClusterNode   `json:"node"`
	Object ClusterObject `json:"object"`
	Status ClusterStatus `json:"status"`
}

// ClusterConfig defines model for clusterConfig.
type ClusterConfig = map[string]interface{}

// ClusterNode defines model for clusterNode.
type ClusterNode = map[string]interface{}

// ClusterObject defines model for clusterObject.
type ClusterObject = map[string]interface{}

// ClusterStatus defines model for clusterStatus.
type ClusterStatus = map[string]interface{}

// Daemon defines model for daemon.
type Daemon struct {
	Goroutines int `json:"goroutines"`
	Sub        Sub `json:"sub"`
}

// Error defines model for error.
type Error struct {
	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// InstanceStatus defines model for instanceStatus.
type InstanceStatus struct {
	App         *App          `json:"app,omitempty"`
	Avail       Status        `json:"avail"`
	Children    *PathRelation `json:"children,omitempty"`
	Constraints *bool         `json:"constraints,omitempty"`
	Csum        *string       `json:"csum,omitempty"`
	Drp         *bool         `json:"drp,omitempty"`
	Env         *string       `json:"env,omitempty"`
	FlexMax     *int          `json:"flex_max,omitempty"`
	FlexMin     *int          `json:"flex_min,omitempty"`
	FlexTarget  *int          `json:"flex_target,omitempty"`
	Frozen      time.Time     `json:"frozen"`
	Kind        Kind          `json:"kind"`
	Optional    *Status       `json:"optional,omitempty"`
	Orchestrate *Orchestrate  `json:"orchestrate,omitempty"`
	Overall     Status        `json:"overall"`
	Parents     *PathRelation `json:"parents,omitempty"`

	// object placement policy
	Placement *Placement `json:"placement,omitempty"`

	// preserve is true if this status has not been updated due to a
	// heartbeat downtime covered by a maintenance period.
	// when the maintenance period ends, the status should be unchanged,
	// and preserve will be set to false.
	Preserved *bool `json:"preserved,omitempty"`

	// scheduling priority of an object instance on a its node
	Priority *Priority `json:"priority,omitempty"`

	// service, instance or resource provisioned state
	Provisioned Provisioned              `json:"provisioned"`
	Resources   *[]ResourceExposedStatus `json:"resources,omitempty"`
	Running     *[]string                `json:"running,omitempty"`
	Scale       *int                     `json:"scale,omitempty"`
	Slaves      *PathRelation            `json:"slaves,omitempty"`
	StatusGroup *string                  `json:"status_group,omitempty"`

	// resources properties
	Subsets *[]struct {
		Parallel bool   `json:"parallel"`
		Rid      string `json:"rid"`
	} `json:"subsets,omitempty"`

	// object topology
	Topology *Topology `json:"topology,omitempty"`
	Updated  time.Time `json:"updated"`
}

// Kind defines model for kind.
type Kind = string

// NodeInfo defines model for nodeInfo.
type NodeInfo struct {
	// labels is the list of node labels.
	Labels []NodeLabel `json:"labels"`

	// nodename is the name of the node where the labels and paths are coming from.
	Nodename string `json:"nodename"`

	// paths is the list of node to storage array san paths.
	Paths []SanPath `json:"paths"`
}

// NodeLabel defines model for nodeLabel.
type NodeLabel struct {
	// name is the label name.
	Name string `json:"name"`

	// value is the label value.
	Value string `json:"value"`
}

// NodesInfo defines model for nodesInfo.
type NodesInfo = []NodeInfo

// ObjectConfig defines model for objectConfig.
type ObjectConfig struct {
	Data  *map[string]interface{} `json:"data,omitempty"`
	Mtime *time.Time              `json:"mtime,omitempty"`
}

// ObjectFile defines model for objectFile.
type ObjectFile struct {
	Data  []byte    `json:"data"`
	Mtime time.Time `json:"mtime"`
}

// ObjectPath defines model for objectPath.
type ObjectPath = string

// ObjectSelector defines model for objectSelector.
type ObjectSelector = []ObjectPath

// Orchestrate defines model for orchestrate.
type Orchestrate string

// PathRelation defines model for pathRelation.
type PathRelation = []string

// object placement policy
type Placement string

// PostDaemonLogsControl defines model for postDaemonLogsControl.
type PostDaemonLogsControl struct {
	Level PostDaemonLogsControlLevel `json:"level"`
}

// PostDaemonLogsControlLevel defines model for PostDaemonLogsControl.Level.
type PostDaemonLogsControlLevel string

// PostDaemonSubAction defines model for postDaemonSubAction.
type PostDaemonSubAction struct {
	Action PostDaemonSubActionAction `json:"action"`

	// daemon component list
	Subs []string `json:"subs"`
}

// PostDaemonSubActionAction defines model for PostDaemonSubAction.Action.
type PostDaemonSubActionAction string

// PostNodeMonitor defines model for postNodeMonitor.
type PostNodeMonitor struct {
	GlobalExpect *string `json:"global_expect,omitempty"`
	LocalExpect  *string `json:"local_expect,omitempty"`
	State        *string `json:"state,omitempty"`
}

// PostObjectAbort defines model for postObjectAbort.
type PostObjectAbort struct {
	Path string `json:"path"`
}

// PostObjectClear defines model for postObjectClear.
type PostObjectClear struct {
	Path string `json:"path"`
}

// PostObjectMonitor defines model for postObjectMonitor.
type PostObjectMonitor struct {
	GlobalExpect *string `json:"global_expect,omitempty"`
	LocalExpect  *string `json:"local_expect,omitempty"`
	Path         string  `json:"path"`
	State        *string `json:"state,omitempty"`
}

// PostObjectProgress defines model for postObjectProgress.
type PostObjectProgress struct {
	IsPartial *bool  `json:"is_partial,omitempty"`
	Path      string `json:"path"`
	SessionId string `json:"session_id"`
	State     string `json:"state"`
}

// PostObjectStatus defines model for postObjectStatus.
type PostObjectStatus struct {
	Path   string         `json:"path"`
	Status InstanceStatus `json:"status"`
}

// PostObjectSwitchTo defines model for postObjectSwitchTo.
type PostObjectSwitchTo struct {
	Destination []string `json:"destination"`
	Path        string   `json:"path"`
}

// PostRelayMessage defines model for postRelayMessage.
type PostRelayMessage struct {
	ClusterId   string `json:"cluster_id"`
	ClusterName string `json:"cluster_name"`
	Msg         string `json:"msg"`
	Nodename    string `json:"nodename"`
}

// scheduling priority of an object instance on a its node
type Priority = int

// service, instance or resource provisioned state
type Provisioned string

// RelayMessage defines model for relayMessage.
type RelayMessage struct {
	Addr        string    `json:"addr"`
	ClusterId   string    `json:"cluster_id"`
	ClusterName string    `json:"cluster_name"`
	Msg         string    `json:"msg"`
	Nodename    string    `json:"nodename"`
	Updated     time.Time `json:"updated"`
}

// RelayMessageList defines model for relayMessageList.
type RelayMessageList = []RelayMessage

// RelayMessages defines model for relayMessages.
type RelayMessages struct {
	Messages RelayMessageList `json:"messages"`
}

// ResourceExposedStatus defines model for resourceExposedStatus.
type ResourceExposedStatus struct {
	// hints the resource ignores all state transition actions
	Disable *bool `json:"disable,omitempty"`

	// indicates that the resource is handled by the encapsulated agents,
	// and ignored at the hypervisor level
	Encap *bool `json:"encap,omitempty"`

	// key-value pairs providing interesting information to collect
	// site-wide about this resource
	Info  *map[string]interface{} `json:"info,omitempty"`
	Label string                  `json:"label"`
	Log   *[]struct {
		Level   string `json:"level"`
		Message string `json:"message"`
	} `json:"log,omitempty"`

	// tells the daemon if it should trigger a monitor action when the
	// resource is not up
	Monitor *bool `json:"monitor,omitempty"`

	// is resource status aggregated into Overall instead of Avail instance status.
	// Errors in optional resource don't stop a state transition action
	Optional    *bool                    `json:"optional,omitempty"`
	Provisioned *ResourceProvisionStatus `json:"provisioned,omitempty"`
	Restart     *int                     `json:"restart,omitempty"`
	Rid         ResourceId               `json:"rid"`

	// resource should always be up, even after a stop state transition action
	Standby *bool  `json:"standby,omitempty"`
	Status  Status `json:"status"`

	// the name of the subset this resource is assigned to
	Subset *string   `json:"subset,omitempty"`
	Tags   *[]string `json:"tags,omitempty"`
	Type   string    `json:"type"`
}

// ResourceId defines model for resourceId.
type ResourceId = string

// ResourceProvisionStatus defines model for resourceProvisionStatus.
type ResourceProvisionStatus struct {
	Mtime *time.Time `json:"mtime,omitempty"`

	// service, instance or resource provisioned state
	State Provisioned `json:"state"`
}

// ResponseDaemonStatus defines model for responseDaemonStatus.
type ResponseDaemonStatus struct {
	Cluster Cluster `json:"cluster"`
	Daemon  Daemon  `json:"daemon"`
}

// responseEventList is a list of sse
type ResponseEventList = string

// ResponseInfoStatus defines model for responseInfoStatus.
type ResponseInfoStatus struct {
	Info   int    `json:"info"`
	Status string `json:"status"`
}

// ResponseMuxBool defines model for responseMuxBool.
type ResponseMuxBool struct {
	Data []struct {
		Data     bool   `json:"data"`
		Endpoint string `json:"endpoint"`
	} `json:"data"`
	Entrypoint string `json:"entrypoint"`
	Status     int    `json:"status"`
}

// ResponsePostAuthToken defines model for responsePostAuthToken.
type ResponsePostAuthToken struct {
	Token         string    `json:"token"`
	TokenExpireAt time.Time `json:"token_expire_at"`
}

// ResponseText defines model for responseText.
type ResponseText = string

// SanPath defines model for sanPath.
type SanPath struct {
	// initiator is the host side san path endpoint.
	Initiator SanPathInitiator `json:"initiator"`

	// target is the storage array side san path endpoint.
	Target SanPathTarget `json:"target"`
}

// initiator is the host side san path endpoint.
type SanPathInitiator struct {
	// name is a worldwide unique path endpoint identifier.
	Name *string `json:"name,omitempty"`

	// type is the endpoint type.
	Type *string `json:"type,omitempty"`
}

// target is the storage array side san path endpoint.
type SanPathTarget struct {
	// name is a worldwide unique path endpoint identifier.
	Name *string `json:"name,omitempty"`

	// type is a the endpoint type.
	Type *string `json:"type,omitempty"`
}

// Severity defines model for severity.
type Severity = string

// Status defines model for status.
type Status = string

// Sub defines model for sub.
type Sub struct {
	Heartbeats *[]SubHeartbeats `json:"heartbeats,omitempty"`
}

// SubAlert defines model for subAlert.
type SubAlert struct {
	Message  string   `json:"message"`
	Severity Severity `json:"severity"`
}

// SubBase defines model for subBase.
type SubBase struct {
	Alerts     []SubAlert `json:"alerts"`
	Configured time.Time  `json:"configured"`
	Created    time.Time  `json:"created"`
	Id         string     `json:"id"`
	State      string     `json:"state"`
}

// SubHeartbeatPeerStatus defines model for subHeartbeatPeerStatus.
type SubHeartbeatPeerStatus struct {
	Beating bool      `json:"beating"`
	Last    time.Time `json:"last"`
}

// SubHeartbeats defines model for subHeartbeats.
type SubHeartbeats struct {
	Alerts     []SubAlert `json:"alerts"`
	Beating    bool       `json:"beating"`
	Configured time.Time  `json:"configured"`
	Created    time.Time  `json:"created"`
	Id         string     `json:"id"`
	Last       time.Time  `json:"last"`
	State      string     `json:"state"`
}

// object topology
type Topology string

// QueryDuration defines model for queryDuration.
type QueryDuration = string

// QueryEventFilter defines model for queryEventFilter.
type QueryEventFilter = []string

// QueryLimit defines model for queryLimit.
type QueryLimit = int64

// QueryNamespaceOptional defines model for queryNamespaceOptional.
type QueryNamespaceOptional = string

// QueryObjectPath defines model for queryObjectPath.
type QueryObjectPath = string

// QueryObjectSelector defines model for queryObjectSelector.
type QueryObjectSelector = string

// QueryRelativesOptional defines model for queryRelativesOptional.
type QueryRelativesOptional = bool

// QueryRelayClusterId defines model for queryRelayClusterId.
type QueryRelayClusterId = string

// QueryRelayNodename defines model for queryRelayNodename.
type QueryRelayNodename = string

// QuerySelectorOptional defines model for querySelectorOptional.
type QuerySelectorOptional = string

// GetDaemonEventsParams defines parameters for GetDaemonEvents.
type GetDaemonEventsParams struct {
	// max duration
	Duration *QueryDuration `form:"duration,omitempty" json:"duration,omitempty"`

	// limit items count
	Limit *QueryLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// list of event filter
	Filter *QueryEventFilter `form:"filter,omitempty" json:"filter,omitempty"`
}

// PostDaemonLogsControlJSONBody defines parameters for PostDaemonLogsControl.
type PostDaemonLogsControlJSONBody = PostDaemonLogsControl

// GetDaemonStatusParams defines parameters for GetDaemonStatus.
type GetDaemonStatusParams struct {
	// namespace
	Namespace *QueryNamespaceOptional `form:"namespace,omitempty" json:"namespace,omitempty"`

	// relatives
	Relatives *QueryRelativesOptional `form:"relatives,omitempty" json:"relatives,omitempty"`

	// selector
	Selector *QuerySelectorOptional `form:"selector,omitempty" json:"selector,omitempty"`
}

// PostDaemonSubActionJSONBody defines parameters for PostDaemonSubAction.
type PostDaemonSubActionJSONBody = PostDaemonSubAction

// PostNodeMonitorJSONBody defines parameters for PostNodeMonitor.
type PostNodeMonitorJSONBody = PostNodeMonitor

// PostObjectAbortJSONBody defines parameters for PostObjectAbort.
type PostObjectAbortJSONBody = PostObjectAbort

// PostObjectClearJSONBody defines parameters for PostObjectClear.
type PostObjectClearJSONBody = PostObjectClear

// GetObjectConfigParams defines parameters for GetObjectConfig.
type GetObjectConfigParams struct {
	// object path
	Path QueryObjectPath `form:"path" json:"path"`

	// evaluate
	Evaluate *bool `form:"evaluate,omitempty" json:"evaluate,omitempty"`

	// impersonate the evaluation as node
	Impersonate *string `form:"impersonate,omitempty" json:"impersonate,omitempty"`
}

// GetObjectFileParams defines parameters for GetObjectFile.
type GetObjectFileParams struct {
	// object path
	Path QueryObjectPath `form:"path" json:"path"`
}

// PostObjectMonitorJSONBody defines parameters for PostObjectMonitor.
type PostObjectMonitorJSONBody = PostObjectMonitor

// PostObjectProgressJSONBody defines parameters for PostObjectProgress.
type PostObjectProgressJSONBody = PostObjectProgress

// GetObjectSelectorParams defines parameters for GetObjectSelector.
type GetObjectSelectorParams struct {
	// object selector
	Selector QueryObjectSelector `form:"selector" json:"selector"`
}

// PostObjectStatusJSONBody defines parameters for PostObjectStatus.
type PostObjectStatusJSONBody = PostObjectStatus

// PostObjectSwitchToJSONBody defines parameters for PostObjectSwitchTo.
type PostObjectSwitchToJSONBody = PostObjectSwitchTo

// GetRelayMessageParams defines parameters for GetRelayMessage.
type GetRelayMessageParams struct {
	// the nodename component of the slot id on the relay
	Nodename *QueryRelayNodename `form:"nodename,omitempty" json:"nodename,omitempty"`

	// the cluster id component of the slot id on the relay
	ClusterId *QueryRelayClusterId `form:"cluster_id,omitempty" json:"cluster_id,omitempty"`
}

// PostRelayMessageJSONBody defines parameters for PostRelayMessage.
type PostRelayMessageJSONBody = PostRelayMessage

// PostDaemonLogsControlJSONRequestBody defines body for PostDaemonLogsControl for application/json ContentType.
type PostDaemonLogsControlJSONRequestBody = PostDaemonLogsControlJSONBody

// PostDaemonSubActionJSONRequestBody defines body for PostDaemonSubAction for application/json ContentType.
type PostDaemonSubActionJSONRequestBody = PostDaemonSubActionJSONBody

// PostNodeMonitorJSONRequestBody defines body for PostNodeMonitor for application/json ContentType.
type PostNodeMonitorJSONRequestBody = PostNodeMonitorJSONBody

// PostObjectAbortJSONRequestBody defines body for PostObjectAbort for application/json ContentType.
type PostObjectAbortJSONRequestBody = PostObjectAbortJSONBody

// PostObjectClearJSONRequestBody defines body for PostObjectClear for application/json ContentType.
type PostObjectClearJSONRequestBody = PostObjectClearJSONBody

// PostObjectMonitorJSONRequestBody defines body for PostObjectMonitor for application/json ContentType.
type PostObjectMonitorJSONRequestBody = PostObjectMonitorJSONBody

// PostObjectProgressJSONRequestBody defines body for PostObjectProgress for application/json ContentType.
type PostObjectProgressJSONRequestBody = PostObjectProgressJSONBody

// PostObjectStatusJSONRequestBody defines body for PostObjectStatus for application/json ContentType.
type PostObjectStatusJSONRequestBody = PostObjectStatusJSONBody

// PostObjectSwitchToJSONRequestBody defines body for PostObjectSwitchTo for application/json ContentType.
type PostObjectSwitchToJSONRequestBody = PostObjectSwitchToJSONBody

// PostRelayMessageJSONRequestBody defines body for PostRelayMessage for application/json ContentType.
type PostRelayMessageJSONRequestBody = PostRelayMessageJSONBody

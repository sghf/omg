package object

import (
	"opensvc.com/opensvc/core/keywords"
	"opensvc.com/opensvc/core/placement"
	"opensvc.com/opensvc/util/converters"
	"opensvc.com/opensvc/util/key"
)

var keywordStore = keywords.Store{
	{
		Section:     "DEFAULT",
		Option:      "nodes",
		Converter:   converters.ListLowercase,
		Text:        "A node selector expression specifying the list of cluster nodes hosting service instances.",
		DefaultText: "The lowercased hostname of the evaluating node.",
		Example:     "n1 n*",
	},
	{
		Section:   "DEFAULT",
		Option:    "drpnodes",
		Converter: converters.ListLowercase,
		Text:      "The backup node where the service is activated in a DRP situation. This node is also a data synchronization target for :c-res:`sync` resources.",
		Example:   "n1 n2",
	},
	{
		Section:   "DEFAULT",
		Option:    "encapnodes",
		Converter: converters.ListLowercase,
		Text:      "The list of `containers` handled by this service and with an OpenSVC agent installed to handle the encapsulated resources. With this parameter set, parameters can be scoped with the ``@encapnodes`` suffix.",
		Example:   "n1 n2",
	},
	{
		Section: "DEFAULT",
		Option:  "app",
		Default: "default",
		Text:    "Used to identify who is responsible for this service, who is billable and provides a most useful filtering key. Better keep it a short code.",
	},
	{
		Section:     "DEFAULT",
		Option:      "env",
		DefaultText: "Same as the node env",
		Candidates:  []string{"CERT", "DEV", "DRP", "FOR", "INT", "PRA", "PRD", "PRJ", "PPRD", "QUAL", "REC", "STG", "TMP", "TST", "UAT"},
		Text:        "A non-PRD service can not be brought up on a PRD node, but a PRD service can be startup on a non-PRD node (in a DRP situation). The default value is the node :kw:`env`.",
	},
	{
		Section:    "DEFAULT",
		Option:     "placement",
		Default:    "nodes order",
		Candidates: placement.Names(),
		Text: `Set a service instances placement policy:

* none        no placement policy. a policy for dummy, observe-only, services.
* nodes order the left-most available node is allowed to start a service instance when necessary.
* load avg    the least loaded node takes precedences.
* shift       shift the nodes order ranking by the service prefix converter to an integer.
* spread      a spread policy tends to perfect leveling with many services.
* score       the highest scoring node takes precedence (the score is a composite indice of load, mem and swap).
`,
	},
	{
		Section:    "DEFAULT",
		Option:     "topology",
		Default:    "failover",
		Candidates: []string{"failover", "flex"},
		Text:       "``failover`` the service is allowed to be up on one node at a time. ``flex`` the service can be up on :kw:`flex_target` nodes, where :kw:`flex_target` must be in the [flex_min, flex_max] range.",
	},
	{
		Section:   "DEFAULT",
		Option:    "flex_min",
		Default:   "1",
		Converter: converters.Int,
		//Depends: []keyval.T{
		//	{key.Parse("topology"), "flex"},
		//},
		Text: "Minimum number of up instances in the cluster. Below this number the aggregated service status is degraded to warn..",
	},
	{
		Section:     "DEFAULT",
		Option:      "flex_max",
		DefaultText: "Number of svc nodes",
		Converter:   converters.Int,
		//Depends: []keyval.T{
		//	{key.Parse("topology"), "flex"},
		//},
		Text: "Maximum number of up instances in the cluster. Above this number the aggregated service status is degraded to warn. ``0`` means unlimited.",
	},
	{
		Section:     "DEFAULT",
		Option:      "flex_target",
		DefaultText: "The value of flex_min",
		Converter:   converters.Int,
		//Depends: []keyval.T{
		//	{key.Parse("topology"), "flex"},
		//},
		Text: "Optimal number of up instances in the cluster. The value must be between :kw:`flex_min` and :kw:`flex_max`. If ``orchestrate=ha``, the monitor ensures the :kw:`flex_target` is met.",
	},
	{
		Section:   "DEFAULT",
		Option:    "priority",
		Default:   "50",
		Converter: converters.Int,
		Text:      "A scheduling priority (the smaller the more priority) used by the monitor thread to trigger actions for the top priority services, so that the :kw:`node.max_parallel` constraint doesn't prevent prior services to start first. The priority setting is dropped from a service configuration injected via the api by a user not granted the prioritizer role.",
	},
}

func (t Base) KeywordLookup(k key.T) keywords.Keyword {
	switch k.Section {
	case "data", "env":
		return keywords.Keyword{
			Option:   "*", // trick IsZero()
			Scopable: true,
			Required: false,
		}
	}
	return keywordStore.Lookup(k)
}

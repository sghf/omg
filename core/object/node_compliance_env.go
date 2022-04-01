package object

import (
	"opensvc.com/opensvc/util/compliance"
)

type (
	// OptsNodeComplianceEnv is the options of the ComplianceEnv function.
	OptsNodeComplianceEnv struct {
		Global    OptsGlobal
		Moduleset OptModuleset
		Module    OptModule
	}
)

func (t Node) ComplianceEnv(options OptsNodeComplianceEnv) (compliance.Envs, error) {
	client, err := t.collectorComplianceClient()
	if err != nil {
		return nil, err
	}
	comp := compliance.New()
	comp.SetCollectorClient(client)
	run := comp.NewRun()
	run.SetModulesetsExpr(options.Moduleset.Moduleset)
	run.SetModulesExpr(options.Module.Module)
	return run.Env()
}

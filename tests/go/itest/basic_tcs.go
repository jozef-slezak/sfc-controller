package itest

import (
	"testing"
	sfccore "github.com/ligato/sfc-controller/controller/core"
	"github.com/golang/protobuf/proto"
)

type basicTCSuite struct {
	T     *testing.T
	AgentTestHelper
	Given Given
	When  When
	Then  Then
}

// DefaultSetup injects Given dependencies
func (t *basicTCSuite) DefaultSetup() {
	t.AgentTestHelper.DefaultSetup(t.T)
	t.Given.agentT = &t.AgentTestHelper
	t.Then.agentT = &t.AgentTestHelper
}

// TC01ResyncEmptyVpp1Agent asserts that vpp agent writes properly vpp-agent configuration
// This TC assumes that vpp-agent configuration was empty before the test.
// Then a specific configuration is written to ETCD and after that SFC Controller starts.
func (t *basicTCSuite) TC01ResyncEmptyVpp1Agent(sfcCfg *sfccore.YamlConfig, vppAgentCfg ... proto.Message) {
	t.DefaultSetup()
	defer t.Teardown()

	t.Given.EmptyETCD()
	t.Given.ConfigSFCviaETCD(sfcCfg)
	t.Given.StartAgent()
	t.Then.VppAgentCfgContains("HOST-1", vppAgentCfg...)
	t.Then.HTTPGetEntities(sfcCfg)
}

// TC02HTTPPostasserts that vpp agent writes properly vpp-agent configuration
// This TC assumes that vpp-agent configuration was empty before the test.
// Then SFC Controller starts and after that SFC Controller is configured via REST HTTP posts.
func (t *basicTCSuite) TC02HTTPPost(sfcCfg *sfccore.YamlConfig, vppAgentCfg ... proto.Message) {
	t.DefaultSetup()
	defer t.Teardown()

	t.Given.EmptyETCD()
	t.Given.StartAgent()
	t.Given.ConfigSFCviaREST(sfcCfg)
	t.Then.VppAgentCfgContains("HOST-1", vppAgentCfg...)
	t.Then.HTTPGetEntities(sfcCfg)
}

// TC03CleanupAtStartupFlag checks that SFC Controller deletes existing ETCD configuration at startup
// if the clean config was set
func (t *basicTCSuite) TC03CleanupAtStartupFlag(etcdConfig *sfccore.YamlConfig) {
	t.DefaultSetup()
	defer t.Teardown()

	t.Given.EmptyETCD()
	t.Given.ConfigSFCviaETCD(etcdConfig)
	t.Given.SetFlagCleanConfigSFC()
	t.Given.StartAgent()
	t.Then.ConfigSFCisEmpty("HOST-1")
}

// TC04LoadConfigFile checks that SFC Controller loads the configuration from config (ETCD is empty)
func (t *basicTCSuite) TC04LoadConfigFile(sfcCfg *sfccore.YamlConfig, vppAgentCfg ... proto.Message) {
	t.DefaultSetup()
	defer t.Teardown()

	t.Given.EmptyETCD()
	t.Given.ConfigSFCviaFile(sfcCfg)
	t.Given.StartAgent()
	t.Then.VppAgentCfgContains("HOST-1", vppAgentCfg...)
	t.Then.HTTPGetEntities(sfcCfg)
}

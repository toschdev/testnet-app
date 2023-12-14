package cmd

import (
	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/spf13/cobra"

	"github.com/toschdev/ignite-testnet/network"
)

const flagOut = "out"

// NewNetworkCoordinatorManageShow creates a new chain show
// command to show a chain details on SPN.
func NewNetworkCoordinatorManageShow() *cobra.Command {
	c := &cobra.Command{
		Use:   "show",
		Short: "Show details of a chain",
	}
	c.AddCommand(
		NewNetworkCoordinatorManageShowInfo(),
		NewNetworkCoordinatorManageShowGenesis(),
		NewNetworkCoordinatorManageShowAccounts(),
		NewNetworkCoordinatorManageShowValidators(),
		NewNetworkCoordinatorManageShowPeers(),
	)
	return c
}

func networkChainLaunch(cmd *cobra.Command, args []string, session *cliui.Session) (NetworkBuilder, uint64, error) {
	nb, err := newNetworkBuilder(cmd, CollectEvents(session.EventBus()))
	if err != nil {
		return nb, 0, err
	}
	// parse launch ID.
	launchID, err := network.ParseID(args[0])
	if err != nil {
		return nb, launchID, err
	}
	return nb, launchID, err
}
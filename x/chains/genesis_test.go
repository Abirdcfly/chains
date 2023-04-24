package chains_test

import (
	"testing"

	keepertest "github.com/Abirdcfly/chains/testutil/keeper"
	"github.com/Abirdcfly/chains/testutil/nullify"
	"github.com/Abirdcfly/chains/x/chains"
	"github.com/Abirdcfly/chains/x/chains/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChainsKeeper(t)
	chains.InitGenesis(ctx, *k, genesisState)
	got := chains.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package keeper_test

import (
	"testing"

	testkeeper "github.com/Abirdcfly/chains/testutil/keeper"
	"github.com/Abirdcfly/chains/x/chains/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChainsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

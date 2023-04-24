package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Abirdcfly/chains/testutil/keeper"
	"github.com/Abirdcfly/chains/x/chains/keeper"
	"github.com/Abirdcfly/chains/x/chains/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChainsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

package keeper_test

import (
	"context"
	"testing"

	keepertest "athlete/testutil/keeper"
	"athlete/x/athlete/keeper"
	"athlete/x/athlete/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AthleteKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

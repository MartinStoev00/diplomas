package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/MartinStoev00/diplomas/testutil/keeper"
	"github.com/MartinStoev00/diplomas/x/diplomas/keeper"
	"github.com/MartinStoev00/diplomas/x/diplomas/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DiplomasKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

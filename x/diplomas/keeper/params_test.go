package keeper_test

import (
	"testing"

	testkeeper "github.com/MartinStoev00/diplomas/testutil/keeper"
	"github.com/MartinStoev00/diplomas/x/diplomas/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DiplomasKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

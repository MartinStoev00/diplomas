package diplomas_test

import (
	"testing"

	keepertest "github.com/MartinStoev00/diplomas/testutil/keeper"
	"github.com/MartinStoev00/diplomas/testutil/nullify"
	"github.com/MartinStoev00/diplomas/x/diplomas"
	"github.com/MartinStoev00/diplomas/x/diplomas/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DiplomasKeeper(t)
	diplomas.InitGenesis(ctx, *k, genesisState)
	got := diplomas.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

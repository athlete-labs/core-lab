package athlete_test

import (
	"testing"

	keepertest "athlete/testutil/keeper"
	"athlete/testutil/nullify"
	"athlete/x/athlete"
	"athlete/x/athlete/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AthleteKeeper(t)
	athlete.InitGenesis(ctx, *k, genesisState)
	got := athlete.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package lone_test

import (
	"testing"

	keepertest "L-ONE/testutil/keeper"
	"L-ONE/testutil/nullify"
	lone "L-ONE/x/lone/module"
	"L-ONE/x/lone/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoneKeeper(t)
	lone.InitGenesis(ctx, k, genesisState)
	got := lone.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

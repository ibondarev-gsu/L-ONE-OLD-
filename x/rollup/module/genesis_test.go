package rollup_test

import (
	"testing"

	keepertest "L-ONE/testutil/keeper"
	"L-ONE/testutil/nullify"
	rollup "L-ONE/x/rollup/module"
	"L-ONE/x/rollup/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RollupKeeper(t)
	rollup.InitGenesis(ctx, k, genesisState)
	got := rollup.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}

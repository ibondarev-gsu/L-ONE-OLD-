package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "L-ONE/testutil/keeper"
	"L-ONE/x/lone/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LoneKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mezum/testutil/keeper"
	"mezum/x/mezum/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MezumKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package mezum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mezum/testutil/keeper"
	"mezum/testutil/nullify"
	"mezum/x/mezum"
	"mezum/x/mezum/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MezumKeeper(t)
	mezum.InitGenesis(ctx, *k, genesisState)
	got := mezum.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

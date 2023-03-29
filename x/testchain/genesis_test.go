package testchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testchain/testutil/keeper"
	"testchain/testutil/nullify"
	"testchain/x/testchain"
	"testchain/x/testchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestchainKeeper(t)
	testchain.InitGenesis(ctx, *k, genesisState)
	got := testchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

package nfl

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmbed(t *testing.T) {
	tcs := []struct {
		txs                   func() Transactions
		expectedAllDayAddress string
		expectedNftAddress    string
	}{
		{
			txs:                   func() Transactions { return ResolvedTransactions },
			expectedAllDayAddress: "0x4dfd62c88d1b6462",
			expectedNftAddress:    "0x631e88ae7f1d7c20",
		},
		{
			txs: func() Transactions {
				txs, err := NewTransactions(context.Background(), Config{
					AllDayAddress: "0x0000000000000001",
					NFTAddress:    "0x0000000000000002",
				})
				assert.NoError(t, err)
				return txs
			},
			expectedAllDayAddress: "0x0000000000000001",
			expectedNftAddress:    "0x0000000000000002",
		},
		{
			txs: func() Transactions {
				conf, err := ConfigForEnv("testnet")
				assert.NoError(t, err)
				txs, err := NewTransactions(context.Background(), conf)
				assert.NoError(t, err)
				return txs
			},
			expectedAllDayAddress: "0x4dfd62c88d1b6462",
			expectedNftAddress:    "0x631e88ae7f1d7c20",
		},
	}

	for _, tc := range tcs {
		txs := tc.txs()
		assert.Contains(t, string(txs.EditionsCreateEdition), formatImport("AllDay", tc.expectedAllDayAddress))

		assert.Contains(t, string(txs.NftsMintMomentNft), formatImport("AllDay", tc.expectedAllDayAddress))
		//assert.Contains(t, string(txs.NftsMintPinNft), formatImport("NonFungibleToken", tc.expectedNftAddress))

		assert.Contains(t, string(txs.EditionsCreateEdition), formatImport("AllDay", tc.expectedAllDayAddress))
		//assert.Contains(t, string(txs.NftsAddXpToNft), formatImport("NonFungibleToken", tc.expectedNftAddress))
	}
}

func formatImport(contractName string, expected string) string {
	return fmt.Sprintf("import %v from %v", contractName, expected)
}

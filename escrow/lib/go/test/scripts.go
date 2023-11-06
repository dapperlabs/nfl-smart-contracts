package test

import (
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
)

// Accounts
func accountIsSetup(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	address flow.Address,
) bool {
	script := loadEscrowAccountIsSetupScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.BytesToAddress(address.Bytes()))})

	return result.ToGoValue().(bool)
}

func getSeriesData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) SeriesData {
	script := loadEscrowReadSeriesByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parseSeriesData(result)
}

func getSetData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) SetData {
	script := loadEscrowReadSetByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parseSetData(result)
}

func getPlayData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) PlayData {
	script := loadEscrowReadPlayByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parsePlayData(result)
}

func getEditionData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) EditionData {
	script := loadEscrowReadEditionByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parseEditionData(result)
}

func getMomentNFTSupply(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
) uint64 {
	script := loadEscrowReadMomentNFTSupplyScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{})

	return result.ToGoValue().(uint64)
}

func getMomentNFTProperties(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	collectionAddress flow.Address,
	nftID uint64,
) OurNFTData {
	script := loadEscrowReadMomentNFTPropertiesScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{
		jsoncdc.MustEncode(cadence.BytesToAddress(collectionAddress.Bytes())),
		jsoncdc.MustEncode(cadence.UInt64(nftID)),
	})

	return parseNFTProperties(result)
}

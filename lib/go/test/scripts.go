package test

import (
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
)

func getCurrentSeriesData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
) SeriesData {
	script := loadGeniesReadCurrentSeriesScript(contracts)
	result := executeScriptAndCheck(t, b, script, nil)

	return parseSeriesData(result)
}

func getSeriesData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint32,
) SeriesData {
	script := loadGeniesReadSeriesByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt32(id))})

	return parseSeriesData(result)
}

func getGeniesCollectionData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint32,
) GeniesCollectionData {
	script := loadGeniesReadGeniesCollectionScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt32(id))})

	return parseGeniesCollectionData(result)
}

func getEditionData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint32,
) EditionData {
	script := loadGeniesReadEditionScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt32(id))})

	return parseEditionData(result)
}

func getGeniesSupply(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
) uint64 {
	script := loadGeniesReadGeniesSupplyScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{})

	return result.ToGoValue().(uint64)
}

func getGeniesNFTProperties(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	collectionAddress flow.Address,
	nftID uint64,
) OurNFTData {
	script := loadGeniesReadNFTPropertiesScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{
		jsoncdc.MustEncode(cadence.BytesToAddress(collectionAddress.Bytes())),
		jsoncdc.MustEncode(cadence.UInt64(nftID)),
	})

	return parseNFTProperties(result)
}

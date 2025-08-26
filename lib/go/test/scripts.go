package test

import (
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-emulator/emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/require"
)

// Accounts
func accountIsSetup(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	address flow.Address,
) bool {
	script := loadAllDayAccountIsSetupScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.BytesToAddress(address.Bytes()))})

	return bool(result.(cadence.Bool))
}

func getSeriesData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) SeriesData {
	script := loadAllDayReadSeriesByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parseSeriesData(result)
}

func getSetData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) SetData {
	script := loadAllDayReadSetByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parseSetData(result)
}

func getPlayData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) PlayData {
	script := loadAllDayReadPlayByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parsePlayData(result)
}

func getEditionData(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint64,
) EditionData {
	script := loadAllDayReadEditionByIDScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.UInt64(id))})

	return parseEditionData(result)
}

func getMomentNFTSupply(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
) uint64 {
	script := loadAllDayReadMomentNFTSupplyScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{})

	return uint64(result.(cadence.UInt64))
}

func getMomentNFTProperties(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	collectionAddress flow.Address,
	nftID uint64,
) OurNFTData {
	script := loadAllDayReadMomentNFTPropertiesScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{
		jsoncdc.MustEncode(cadence.BytesToAddress(collectionAddress.Bytes())),
		jsoncdc.MustEncode(cadence.UInt64(nftID)),
	})

	return parseNFTProperties(result)
}

func getMomentNFTMetadata(t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	address flow.Address,
	nftID uint64,
	shouldRevert bool,
) []cadence.Struct {
	script := loadAllDayReadMomentNFTMetadataScript(contracts)
	result := executeScriptAndCheck(t, b, script,
		[][]byte{jsoncdc.MustEncode(cadence.BytesToAddress(address.Bytes())), jsoncdc.MustEncode(cadence.UInt64(nftID))})

	cArray := result.(cadence.Array).Values
	resultArray := make([]cadence.Struct, len(cArray))
	for i, val := range cArray {
		resultArray[i] = val.(cadence.Struct)
	}

	return resultArray
}

// Badges
func getBadgeBySlug(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	slug string,
) *BadgeData {
	script := loadAllDayGetBadgeBySlugScript(contracts)
	slugStr, err := cadence.NewString(slug)
	require.NoError(t, err)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(slugStr)})

	if optional, ok := result.(cadence.Optional); ok {
		if optional.Value == nil {
			return nil
		}
		badge := parseBadgeData(optional.Value)
		return &badge
	}
	return nil
}

func getNftAllBadges(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	account flow.Address,
	nftID uint64,
) []BadgeData {
	script := loadAllDayGetNftAllBadgesScript(contracts)
	result := executeScriptAndCheck(t, b, script, [][]byte{
		jsoncdc.MustEncode(cadence.BytesToAddress(account.Bytes())),
		jsoncdc.MustEncode(cadence.UInt64(nftID)),
	})

	return parseBadgeArray(result)
}

func badgeExists(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	slug string,
) bool {
	script := loadAllDayBadgeExistsScript(contracts)
	slugStr, err := cadence.NewString(slug)
	require.NoError(t, err)
	result := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(slugStr)})

	return bool(result.(cadence.Bool))
}

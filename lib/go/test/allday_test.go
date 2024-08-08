package test

import (
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/fixedpoint"
	"testing"

	"github.com/onflow/flow-emulator/emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------
// Setup
// ------------------------------------------------------------
func TestAllDayDeployContracts(t *testing.T) {
	b := newEmulator()
	AllDayDeployContracts(t, b)
}

func TestAllDaySetupAccount(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDay(t, b, userAddress, userSigner, contracts)

	t.Run("Account should be set up", func(t *testing.T) {
		accountIsSetUp := accountIsSetup(
			t,
			b,
			contracts,
			userAddress,
		)
		assert.Equal(t, true, accountIsSetUp)
	})
}

// ------------------------------------------------------------
// Series
// ------------------------------------------------------------
func TestSeries(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	createTestSeries(t, b, contracts)
}

func createTestSeries(t *testing.T, b *emulator.Blockchain, contracts Contracts) {
	t.Run("Should be able to create a new series", func(t *testing.T) {
		testCreateSeries(
			t,
			b,
			contracts,
			"Series One",
			1,
			false,
		)
	})

	t.Run("Should be able to create a new series with an incrementing ID from the first", func(t *testing.T) {
		testCreateSeries(
			t,
			b,
			contracts,
			"Series Two",
			2,
			false,
		)
	})

	t.Run("Should be able to close a series", func(t *testing.T) {
		testCloseSeries(
			t,
			b,
			contracts,
			2,
			false,
		)
	})
}

func testCreateSeries(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	seriesName string,
	shouldBeID uint64,
	shouldRevert bool,
) {
	createSeries(
		t,
		b,
		contracts,
		seriesName,
		false,
	)

	if !shouldRevert {
		series := getSeriesData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, series.ID)
		assert.Equal(t, seriesName, series.Name)
		assert.Equal(t, true, series.Active)
	}
}

func testCloseSeries(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	seriesID uint64,
	shouldRevert bool,
) {
	wasActive := getSeriesData(t, b, contracts, seriesID).Active
	closeSeries(
		t,
		b,
		contracts,
		seriesID,
		shouldRevert,
	)

	series := getSeriesData(t, b, contracts, seriesID)
	assert.Equal(t, seriesID, series.ID)
	if !shouldRevert {
		assert.Equal(t, false, series.Active)
	} else {
		assert.Equal(t, wasActive, series.Active)
	}
}

// ------------------------------------------------------------
// Sets
// ------------------------------------------------------------
func TestSets(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	createTestSets(t, b, contracts)

}

func createTestSets(t *testing.T, b *emulator.Blockchain, contracts Contracts) {
	t.Run("Should be able to create a new set", func(t *testing.T) {
		testCreateSet(
			t,
			b,
			contracts,
			"Set One",
			1,
			false,
		)
	})

	t.Run("Should be able to create a new set with an incrementing ID from the first", func(t *testing.T) {
		testCreateSet(
			t,
			b,
			contracts,
			"Set Two",
			2,
			false,
		)
	})
}

func testCreateSet(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	setName string,
	shouldBeID uint64,
	shouldRevert bool,
) {
	createSet(
		t,
		b,
		contracts,
		setName,
		false,
	)

	if !shouldRevert {
		set := getSetData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, set.ID)
		assert.Equal(t, setName, set.Name)
	}
}

// ------------------------------------------------------------
// Plays
// ------------------------------------------------------------
func TestPlays(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	createTestPlays(t, b, contracts)
}

func createTestPlays(t *testing.T, b *emulator.Blockchain, contracts Contracts) {
	t.Run("Should be able to create a new play", func(t *testing.T) {
		metadata := map[string]string{
			"playerFirstName": "Apple",
			"playerLastName":  "Alpha",
			"playType":        "Interception",
			"description":     "Fabulous diving interception by AA",
		}
		testCreatePlay(
			t,
			b,
			contracts,
			"TEST_CLASSIFICATION",
			metadata,
			1,
			false,
		)
	})

	t.Run("Should be able to create a new play with an incrementing ID from the first", func(t *testing.T) {
		metadata := map[string]string{
			"playerFirstName": "Bear",
			"playerLastName":  "Bravo",
			"playType":        "Rush",
		}
		testCreatePlay(
			t,
			b,
			contracts,
			"TEST_CLASSIFICATION",
			metadata,
			2,
			false,
		)
	})
}

func testCreatePlay(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	classification string,
	metadata map[string]string,
	shouldBeID uint64,
	shouldRevert bool,
) {
	createPlay(
		t,
		b,
		contracts,
		classification,
		metadata,
		false,
	)

	if !shouldRevert {
		play := getPlayData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, play.ID)
		assert.Equal(t, classification, play.Classification)
		assert.Equal(t, metadata, play.Metadata)
	}
}

// ------------------------------------------------------------
// Editions
// ------------------------------------------------------------
func TestEditions(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	createTestEditions(t, b, contracts)
}

func testCreateEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	seriesID uint64,
	setID uint64,
	playID uint64,
	maxMintSize *uint64,
	tier string,
	shouldBeID uint64,
	shouldRevert bool,
) {
	createEdition(
		t,
		b,
		contracts,
		seriesID,
		setID,
		playID,
		maxMintSize,
		tier,
		shouldRevert,
	)

	if !shouldRevert {
		edition := getEditionData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, edition.ID)
		assert.Equal(t, seriesID, edition.SeriesID)
		assert.Equal(t, setID, edition.SetID)
		assert.Equal(t, playID, edition.PlayID)
		assert.Equal(t, tier, edition.Tier)
		if maxMintSize != nil {
			assert.Equal(t, &maxMintSize, &edition.MaxMintSize)
		}
	}
}

func testCloseEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionID uint64,
	shouldBeID uint64,
	shouldRevert bool,
) {
	closeEdition(
		t,
		b,
		contracts,
		editionID,
		false,
	)

	if !shouldRevert {
		edition := getEditionData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, edition.ID)
	}
}

func createTestEditions(t *testing.T, b *emulator.Blockchain, contracts Contracts) {
	var maxMintSize uint64 = 2
	createTestSeries(t, b, contracts)
	createTestSets(t, b, contracts)
	createTestPlays(t, b, contracts)

	t.Run("Should be able to create a new edition with series/set/play IDs and a max mint size of 100", func(t *testing.T) {
		testCreateEdition(
			t,
			b,
			contracts,
			1,
			1,
			1,
			&maxMintSize,
			"COMMON",
			1,
			false,
		)
	})

	t.Run("Should be able to create another new edition with series/set/play IDs and no max mint size", func(t *testing.T) {
		testCreateEdition(
			t,
			b,
			contracts,
			1,
			2,
			1,
			nil,
			"COMMON",
			2,
			false,
		)
	})

	t.Run("Should be able to create a new edition with series/set/play IDs and no max mint size", func(t *testing.T) {
		testCreateEdition(
			t,
			b,
			contracts,
			1,
			1,
			2,
			nil,
			"COMMON",
			3,
			false,
		)
	})

	t.Run("Should not be able to create a new edition with a closed series", func(t *testing.T) {
		testCreateEdition(
			t,
			b,
			contracts,
			2,
			1,
			1,
			nil,
			"COMMON",
			4,
			true,
		)
	})

	t.Run("Should be able to create an Edition with a Set/Play combination that already exists but with a different tier", func(t *testing.T) {
		//Mint LEGENDARY edition
		testCreateEdition(t, b, contracts, 1 /*seriesID*/, 1 /*setID*/, 2 /*playID*/, nil,
			"LEGENDARY" /*tier*/, 4 /*shouldBEID*/, false /*shouldRevert*/)
	})

	t.Run("Should NOT be able to mint new edition using the same set/play with new tier", func(t *testing.T) {
		//Mint COMMON edition again, tx should revert
		testCreateEdition(t, b, contracts, 1 /*seriesID*/, 1 /*setID*/, 2 /*playID*/, nil,
			"COMMON" /*tier*/, 5 /*shouldBEID*/, true /*shouldRevert*/)
	})

	t.Run("Should be able to close and edition that has no max mint size", func(t *testing.T) {
		testCloseEdition(
			t,
			b,
			contracts,
			3,
			3,
			false,
		)
	})
}

// ------------------------------------------------------------
// MomentNFTs
// ------------------------------------------------------------
func TestMomentNFTs(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDay(t, b, userAddress, userSigner, contracts)

	createTestEditions(t, b, contracts)

	t.Run("Should be able to mint a new MomentNFT from an edition that has a maxMintSize", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(1),
			nil,
			userAddress,
			uint64(1),
			uint64(1),
			false,
		)
	})

	t.Run("Should be able to mint a second new MomentNFT from an edition that has a maxmintSize", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(1),
			nil,
			userAddress,
			uint64(2),
			uint64(2),
			false,
		)
	})

	t.Run("Should be able to mint a new MomentNFT from an edition with no max mint size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(2),
			uint64Ptr(2023),
			userAddress,
			uint64(3),
			uint64(2023),
			false,
		)
	})

	t.Run("Should be able to mint a second new MomentNFT from an edition with no max mint size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(2),
			uint64Ptr(2023),
			userAddress,
			uint64(4),
			uint64(2023),
			false,
		)
	})

	t.Run("Should not be able to mint an edition that has reached max minting size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(1),
			nil,
			userAddress,
			uint64(3),
			uint64(3),
			true,
		)
	})

	t.Run("Should not be able to mint an edition that is already closed", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(3),
			nil,
			userAddress,
			uint64(1),
			uint64(1),
			true,
		)
	})
}

func testMintMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionID uint64,
	serialNumber *uint64,
	userAddress flow.Address,
	shouldBeID uint64,
	shouldBeSerialNumber uint64,
	shouldRevert bool,
) {
	// Make sure the total supply of NFTs is tracked correctly
	previousSupply := getMomentNFTSupply(t, b, contracts)

	mintMomentNFT(
		t,
		b,
		contracts,
		userAddress,
		editionID,
		serialNumber,
		shouldRevert,
	)

	newSupply := getMomentNFTSupply(t, b, contracts)
	if !shouldRevert {
		assert.Equal(t, previousSupply+uint64(1), newSupply)

		nftProperties := getMomentNFTProperties(
			t,
			b,
			contracts,
			userAddress,
			shouldBeID,
		)
		assert.Equal(t, shouldBeID, nftProperties.ID)
		assert.Equal(t, editionID, nftProperties.EditionID)
		assert.Equal(t, shouldBeSerialNumber, nftProperties.SerialNumber)
		assert.Equal(t, shouldBeSerialNumber, nftProperties.SerialNumber)
		//FIXME: query the block time and check equality.
		//       Here we just make sure it's nonzero.
		assert.Less(t, uint64(0), nftProperties.MintingDate)
	} else {
		assert.Equal(t, previousSupply, newSupply)
	}
}

// ------------------------------------------------------------
// MomentNFT MetadataViews
// ------------------------------------------------------------
func TestMomentNFTMetadataViews(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDay(t, b, userAddress, userSigner, contracts)
	createTestEditions(t, b, contracts)
	mintMomentNFT(t, b, contracts, userAddress /*editionID*/, 1, nil /*shouldRevert*/, false)

	t.Run("Should be able to get moment's metadata", func(t *testing.T) {
		result := getMomentNFTMetadata(t, b, contracts, userAddress, 1, false)

		//Validate Display
		displayView := result[0].FieldsMappedByName()
		assert.Equal(t, "Apple Alpha Interception", string(displayView["name"].(cadence.String)))
		assert.Equal(t, "Fabulous diving interception by AA", string(displayView["description"].(cadence.String)))
		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image?format=jpeg&width=256",
			string(displayView["thumbnail"].(cadence.Struct).FieldsMappedByName()["url"].(cadence.String)))

		//Validate Editions
		editionsView := result[1]
		edition := editionsView.FieldsMappedByName()["infoList"].(cadence.Array).Values[0].(cadence.Struct).FieldsMappedByName()
		assert.Equal(t, "Set One: #1", string(edition["name"].(cadence.Optional).Value.(cadence.String)))
		assert.Equal(t, uint64(1), uint64(edition["number"].(cadence.UInt64)))
		assert.Equal(t, uint64(2), uint64(edition["max"].(cadence.Optional).Value.(cadence.UInt64)))

		// Validate External URL
		externalURLView := result[2]
		assert.Equal(t, "https://nflallday.com/moments/1", string(externalURLView.FieldsMappedByName()["url"].(cadence.String)))

		//Validate Medias
		mediasView := result[3]
		medias := mediasView.FieldsMappedByName()["items"].(cadence.Array)
		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image?format=jpeg&width=512",
			getMediaPath(medias.Values[0]))
		assert.Equal(t, "image/jpeg", getMediaType(medias.Values[0]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image-details?format=jpeg&width=512",
			getMediaPath(medias.Values[1]))
		assert.Equal(t, "image/jpeg", getMediaType(medias.Values[1]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image-logo?format=jpeg&width=512",
			getMediaPath(medias.Values[2]))
		assert.Equal(t, "image/jpeg", getMediaType(medias.Values[2]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image-legal?format=jpeg&width=512",
			getMediaPath(medias.Values[3]))
		assert.Equal(t, "image/jpeg", getMediaType(medias.Values[3]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image-player?format=jpeg&width=512",
			getMediaPath(medias.Values[4]))
		assert.Equal(t, "image/jpeg", getMediaType(medias.Values[4]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image-scores?format=jpeg&width=512",
			getMediaPath(medias.Values[5]))
		assert.Equal(t, "image/jpeg", getMediaType(medias.Values[5]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/video",
			getMediaPath(medias.Values[6]))
		assert.Equal(t, "video/mp4", getMediaType(medias.Values[6]))

		assert.Equal(t, "https://media.nflallday.com/editions/1/media/video-idle",
			getMediaPath(medias.Values[7]))
		assert.Equal(t, "video/mp4", getMediaType(medias.Values[7]))

		//Validate NFTCollectionDisplay
		collectionDisplay := result[4].FieldsMappedByName()
		assert.Equal(t, "NFL All Day", string(collectionDisplay["name"].(cadence.String)))
		assert.Equal(t, "Officially Licensed Digital Collectibles Featuring the NFL’s Best Highlights. Buy, Sell and Collect Your Favorite NFL Moments",
			string(collectionDisplay["description"].(cadence.String)))
		assert.Equal(t, "https://nflallday.com/", string(collectionDisplay["externalURL"].(cadence.Struct).FieldsMappedByName()["url"].(cadence.String)))
		assert.Equal(t, "https://assets.nflallday.com/flow/catalogue/NFLAD_SQUARE.png",
			getMediaPath(collectionDisplay["squareImage"]))
		assert.Equal(t, "image/png", getMediaType(collectionDisplay["squareImage"]))
		assert.Equal(t, "https://assets.nflallday.com/flow/catalogue/NFLAD_BANNER.png",
			getMediaPath(collectionDisplay["bannerImage"]))
		assert.Equal(t, "image/png", getMediaType(collectionDisplay["bannerImage"]))
		socials := map[string]cadence.Struct{}
		for _, kvPair := range collectionDisplay["socials"].(cadence.Dictionary).Pairs {
			socials[string(kvPair.Key.(cadence.String))] = kvPair.Value.(cadence.Struct)
		}
		assert.Equal(t, "https://www.instagram.com/nflallday/", string(socials["instagram"].FieldsMappedByName()["url"].(cadence.String)))
		assert.Equal(t, "https://twitter.com/NFLAllDay", string(socials["twitter"].FieldsMappedByName()["url"].(cadence.String)))
		assert.Equal(t, "https://discord.com/invite/5K6qyTzj2k", string(socials["discord"].FieldsMappedByName()["url"].(cadence.String)))

		// Validate Royalties
		royaltiesView := result[5].FieldsMappedByName()
		royalty := royaltiesView["cutInfos"].(cadence.Array).Values[0].(cadence.Struct).FieldsMappedByName()
		assert.Equal(t, contracts.RoyaltyAddress, flow.HexToAddress(royalty["receiver"].(cadence.Capability).Address.Hex()))
		assert.Equal(t, uint64(0.05*fixedpoint.Fix64Factor), uint64(royalty["cut"].(cadence.UFix64)))
		assert.Equal(t, "NFL All Day marketplace royalty", string(royalty["description"].(cadence.String)))

		// Validate Serial
		serialView := result[6]
		assert.Equal(t, uint64(1), uint64(serialView.FieldsMappedByName()["number"].(cadence.UInt64)))

		// Validate Traits
		traitsView := result[7]
		traits := traitsView.FieldsMappedByName()["traits"].(cadence.Array)
		traitsMap := make(map[string]any)
		for _, trait := range traits.Values {
			traitsMap[string(trait.(cadence.Struct).FieldsMappedByName()["name"].(cadence.String))] = trait.(cadence.Struct).FieldsMappedByName()["value"]
		}
		assert.Equal(t, cadence.String("COMMON"), traitsMap["editionTier"])
		assert.Equal(t, cadence.String("Series One"), traitsMap["seriesName"])
		assert.Equal(t, cadence.String("Set One"), traitsMap["setName"])
		assert.Equal(t, cadence.NewUInt64(1), traitsMap["serialNumber"])
		assert.Equal(t, cadence.String("Interception"), traitsMap["playType"])
	})
}

func getMediaPath(media interface{}) interface{} {
	return string(media.(cadence.Struct).FieldsMappedByName()["file"].(cadence.Struct).FieldsMappedByName()["url"].(cadence.String))
}

func getMediaType(media interface{}) interface{} {
	return string(media.(cadence.Struct).FieldsMappedByName()["mediaType"].(cadence.String))
}

func TestUpdatePlayDescription(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDay(t, b, userAddress, userSigner, contracts)
	createTestEditions(t, b, contracts)
	mintMomentNFT(t, b, contracts, userAddress /*editionID*/, 1, nil /*shouldRevert*/, false)

	t.Run("Should be able to update play's description", func(t *testing.T) {
		result := getMomentNFTMetadata(t, b, contracts, userAddress, 1, false)

		//Validate Display
		displayView := result[0].FieldsMappedByName()
		assert.Equal(t, "Apple Alpha Interception", string(displayView["name"].(cadence.String)))
		assert.Equal(t, "Fabulous diving interception by AA", string(displayView["description"].(cadence.String)))
		assert.Equal(t, "https://media.nflallday.com/editions/1/media/image?format=jpeg&width=256",
			string(displayView["thumbnail"].(cadence.Struct).FieldsMappedByName()["url"].(cadence.String)))

		//Update play description
		newPlayDescription := "A new play description"
		updatePlayDescription(t, b, contracts, 1 /*playID*/, newPlayDescription, false /*shouldRevert*/)

		//Validate Display has been updated
		result = getMomentNFTMetadata(t, b, contracts, userAddress, 1, false)
		displayView = result[0].FieldsMappedByName()
		assert.Equal(t, newPlayDescription, string(displayView["description"].(cadence.String)))

	})
}

func TestUpdatePlayDynamicMetadata(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDay(t, b, userAddress, userSigner, contracts)
	createTestEditions(t, b, contracts)
	mintMomentNFT(t, b, contracts, userAddress /*editionID*/, 1, nil /*shouldRevert*/, false)

	t.Run("Should be able to update play's dynamic metadata", func(t *testing.T) {
		//Validate initial Display
		result := getMomentNFTMetadata(t, b, contracts, userAddress, 1, false)
		displayView := result[0].FieldsMappedByName()
		assert.Equal(t, "Apple Alpha Interception", string(displayView["name"].(cadence.String)))

		//Update play metadata
		teamName := "New Team"
		var playerFirstName *string
		playerLastName := "Charlie"
		var playerNumber *string
		var playerPosition *string
		updatePlayDynamicMetadata(t, b, contracts, 1 /*playID*/, &teamName, playerFirstName, &playerLastName,
			playerNumber, playerPosition, false /*shouldRevert*/)

		//Validate Display has been updated
		result = getMomentNFTMetadata(t, b, contracts, userAddress, 1, false)
		displayView = result[0].FieldsMappedByName()
		assert.Equal(t, "Apple Charlie Interception", string(displayView["name"].(cadence.String)))

		//Validate Play metadata has been updated
		traitsView := result[7]
		traits := traitsView.FieldsMappedByName()["traits"].(cadence.Array).Values
		for _, trait := range traits {
			ts := trait.(cadence.Struct).FieldsMappedByName()
			if string(ts["name"].(cadence.String)) == "teamName" {
				assert.Equal(t, teamName, string(ts["value"].(cadence.String)))
			}
			if string(ts["name"].(cadence.String)) == "playerFirstName" {
				assert.Equal(t, "Apple", string(ts["value"].(cadence.String)))
			}
			if string(ts["name"].(cadence.String)) == "playerLastName" {
				assert.Equal(t, playerLastName, string(ts["value"].(cadence.String)))
			}
		}
	})
}

func TestMintMomentMulti(t *testing.T) {
	b := newEmulator()
	contracts := AllDayDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDay(t, b, userAddress, userSigner, contracts)
	createTestEditions(t, b, contracts)
	// edition 1 has a maxSize while edition 2 does not
	editions := []uint64{1, 2}
	serialNumbers := []*uint64{nil, uint64Ptr(2023)}

	mintMomentNFTMulti(t, b, contracts, userAddress /*editionID*/, editions, serialNumbers /*shouldRevert*/, false)

	t.Run("Should have a serial number of 1", func(t *testing.T) {
		nft := getMomentNFTProperties(t, b, contracts, userAddress, 1)
		assert.Equal(t, uint64(1), nft.EditionID)
		assert.Equal(t, uint64(1), nft.SerialNumber)
	})
	t.Run("Should have a serial number of 2023", func(t *testing.T) {
		nft := getMomentNFTProperties(t, b, contracts, userAddress, 2)
		assert.Equal(t, uint64(2), nft.EditionID)
		assert.Equal(t, uint64(2023), nft.SerialNumber)
	})
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}

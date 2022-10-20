package test

import (
	"github.com/onflow/cadence"
	"github.com/onflow/cadence/fixedpoint"
	"testing"

	emulator "github.com/onflow/flow-emulator"
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
			"PlayerFirstName": "Apple",
			"PlayerLastName":  "Alpha",
			"PlayType":        "Interception",
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
			"PlayerFirstName": "Bear",
			"PlayerLastName":  "Bravo",
			"PlayType":        "Rush",
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

	t.Run("Should not be able to create an Edition with a Set/Play combination that already exists", func(t *testing.T) {
		testCreateEdition(
			t,
			b,
			contracts,
			1,
			1,
			2,
			nil,
			"COMMON",
			5,
			true,
		)
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
			userAddress,
			uint64(3),
			uint64(1),
			false,
		)
	})

	t.Run("Should be able to mint a second new MomentNFT from an edition with no max mint size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(2),
			userAddress,
			uint64(4),
			uint64(2),
			false,
		)
	})

	t.Run("Should not be able to mint an edition that has reached max minting size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint64(1),
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
	mintMomentNFT(t, b, contracts, userAddress /*editionID*/, 1 /*shouldRevert*/, false)

	t.Run("Should be able to get moment's metadata", func(t *testing.T) {
		result := getMomentNFTMetadata(t, b, contracts, userAddress, 1, false)

		//Validate Display
		displayView := result[0]
		assert.Equal(t, "Apple Alpha Interception", displayView.Fields[0].ToGoValue())
		assert.Equal(t, "Series One Set One moment with serial number 1", displayView.Fields[1].ToGoValue())
		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image?format=jpeg&width=256",
			displayView.Fields[2].(cadence.Struct).Fields[0].ToGoValue())

		//Validate Editions
		editionsView := result[1]
		edition := editionsView.Fields[0].(cadence.Array).Values[0].(cadence.Struct)
		assert.Equal(t, "Set One: #1", edition.Fields[0].ToGoValue())
		assert.Equal(t, uint64(1), edition.Fields[1].ToGoValue())
		assert.Equal(t, uint64(2), edition.Fields[2].ToGoValue())

		// Validate External URL
		externalURLView := result[2]
		assert.Equal(t, "https://nflallday.com/moments/1", externalURLView.Fields[0].ToGoValue())

		//Validate Medias
		mediasView := result[3]
		editions := mediasView.Fields[0].(cadence.Array)
		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image?format=jpeg&width=512",
			editions.Values[0].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", editions.Values[0].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image-details?format=jpeg&width=512",
			editions.Values[1].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", editions.Values[1].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image-logo?format=jpeg&width=512",
			editions.Values[2].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", editions.Values[2].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image-legal?format=jpeg&width=512",
			editions.Values[3].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", editions.Values[3].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image-player?format=jpeg&width=512",
			editions.Values[4].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", editions.Values[4].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/image-scores?format=jpeg&width=512",
			editions.Values[5].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", editions.Values[5].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/video",
			editions.Values[6].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "video/mp4", editions.Values[6].(cadence.Struct).Fields[1].ToGoValue())

		assert.Equal(t, "https://assets.nflallday.com/editions/1/media/video-idle",
			editions.Values[7].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "video/mp4", editions.Values[7].(cadence.Struct).Fields[1].ToGoValue())

		//Validate NFTCollectionDisplay
		collectionDisplay := result[4]
		assert.Equal(t, "NFL All Day", collectionDisplay.Fields[0].ToGoValue())
		assert.Equal(t, "Officially Licensed Digital Collectibles Featuring the NFL’s Best Highlights. Buy, Sell and Collect Your Favorite NFL Moments",
			collectionDisplay.Fields[1].ToGoValue())
		assert.Equal(t, "https://nflallday.com/", collectionDisplay.Fields[2].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "https://storage.googleapis.com/dl-nfl-assets-prod/static/images/flow-catalogue/NFLAD_SQUARE.png",
			collectionDisplay.Fields[3].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/png", collectionDisplay.Fields[3].(cadence.Struct).Fields[1].ToGoValue())
		assert.Equal(t, "https://storage.googleapis.com/dl-nfl-assets-prod/static/images/flow-catalogue/NFLAD_BANNER_1200x630.jpg",
			collectionDisplay.Fields[4].(cadence.Struct).Fields[0].(cadence.Struct).Fields[0].ToGoValue())
		assert.Equal(t, "image/jpeg", collectionDisplay.Fields[4].(cadence.Struct).Fields[1].ToGoValue())
		socials := map[string]cadence.Struct{}
		for _, kvPair := range collectionDisplay.Fields[5].(cadence.Dictionary).Pairs {
			socials[kvPair.Key.ToGoValue().(string)] = kvPair.Value.(cadence.Struct)
		}
		assert.Equal(t, "https://www.instagram.com/nflallday/", socials["instagram"].Fields[0].ToGoValue())
		assert.Equal(t, "https://twitter.com/NFLAllDay", socials["twitter"].Fields[0].ToGoValue())
		assert.Equal(t, "https://discord.com/invite/5K6qyTzj2k", socials["discord"].Fields[0].ToGoValue())

		// Validate Royalties
		royaltiesView := result[5]
		royalty := royaltiesView.Fields[0].(cadence.Array).Values[0].(cadence.Struct)
		assert.Equal(t, contracts.RoyaltyAddress, flow.HexToAddress(royalty.Fields[0].(cadence.Capability).Address.Hex()))
		assert.Equal(t, uint64(0.05*fixedpoint.Fix64Factor), royalty.Fields[1].ToGoValue())
		assert.Equal(t, "NFL All Day marketplace royalty", royalty.Fields[2].ToGoValue())

		// Validate Serial
		serialView := result[6]
		assert.Equal(t, uint64(1), serialView.Fields[0].ToGoValue())

		// Validate Traits
		traitsView := result[7]
		traits := traitsView.Fields[0].(cadence.Array)
		traitsMap := map[string]interface{}{}
		for _, trait := range traits.Values {
			traitsMap[trait.(cadence.Struct).Fields[0].ToGoValue().(string)] = trait.(cadence.Struct).Fields[1].ToGoValue()
		}
		assert.Equal(t, "COMMON", traitsMap["EditionTier"])
		assert.Equal(t, "Series One", traitsMap["SeriesName"])
		assert.Equal(t, "Set One", traitsMap["SetName"])
		assert.Equal(t, uint64(1), traitsMap["SerialNumber"])
		assert.Equal(t, "Interception", traitsMap["PlayType"])
	})
}

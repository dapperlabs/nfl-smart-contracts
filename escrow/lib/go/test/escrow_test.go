package test

import (
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestEscrowDeployContracts(t *testing.T) {
	b := newEmulator()
	EscrowContracts(t, b)
}

func TestEscrowSetupAccount(t *testing.T) {
	b := newEmulator()
	contracts := EscrowContracts(t, b)
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
	contracts := EscrowContracts(t, b)
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

// ------------------------------------------------------------
// Sets
// ------------------------------------------------------------
func TestSets(t *testing.T) {
	b := newEmulator()
	contracts := EscrowContracts(t, b)
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
	contracts := EscrowContracts(t, b)
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
	contracts := EscrowContracts(t, b)
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
}

// ------------------------------------------------------------
// MomentNFTs
// ------------------------------------------------------------
func TestMomentNFTs(t *testing.T) {
	b := newEmulator()
	contracts := EscrowContracts(t, b)
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
		assert.Less(t, uint64(0), nftProperties.MintingDate)
	} else {
		assert.Equal(t, previousSupply, newSupply)
	}
}

// ------------------------------------------------------------
// Escrow
// ------------------------------------------------------------
func TestEscrow(t *testing.T) {
	b := newEmulator()
	contracts := EscrowContracts(t, b)
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

	t.Run("Should confirm that 1 MomentNFT exists within users collection", func(t *testing.T) {
		// Get the MomentNFT data from the users collection.
		count := getMomentNFTLengthInAccount(t, b, contracts, userAddress)
		assert.Equal(t, big.NewInt(1), count)
	})

	t.Run("Should be able to create a leaderboard", func(t *testing.T) {
		testCreateLeaderboard(
			t,
			b,
			contracts,
			"leaderboardBurn-1",
		)
	})

	t.Run("Should get the leaderboard by name to confirm it exists", func(t *testing.T) {
		testGetLeaderboard(
			t,
			b,
			contracts,
			"leaderboardBurn-1",
		)
	})

	t.Run("Should be able to escrow moment to leaderboard", func(t *testing.T) {
		testEscrowMomentNFT(
			t,
			b,
			contracts,
			userSigner,
			userAddress,
			uint64(1),
		)
	})

	t.Run("Should check that the entries length on Leaderboard is 1", func(t *testing.T) {
		count := getEscrowNFTLengthInLeaderboard(t, b, contracts, "leaderboardBurn-1")
		assert.Equal(t, big.NewInt(1), count)
	})

	t.Run("Should confirm that 0 MomentNFTs exists within users collection due to escrow", func(t *testing.T) {
		// Get the MomentNFT data from the users collection.
		count := getMomentNFTLengthInAccount(t, b, contracts, userAddress)
		assert.Equal(t, big.NewInt(0), count)
	})

	t.Run("Should withdraw Moment from Leaderboard by name", func(t *testing.T) {
		testWithdrawMomentNFT(
			t,
			b,
			contracts,
			"leaderboardBurn-1",
			uint64(1),
		)
	})

	t.Run("Should check that the MomentNFT is back in the users collection", func(t *testing.T) {
		// Get the MomentNFT data from the users collection.
		count := getMomentNFTLengthInAccount(t, b, contracts, userAddress)
		assert.Equal(t, big.NewInt(1), count)
	})

	t.Run("Should check that the entries length on Leaderboard is 0", func(t *testing.T) {
		count := getEscrowNFTLengthInLeaderboard(t, b, contracts, "leaderboardBurn-1")
		assert.Equal(t, big.NewInt(0), count)
	})

	t.Run("Should escrow the moment again", func(t *testing.T) {
		testEscrowMomentNFT(
			t,
			b,
			contracts,
			userSigner,
			userAddress,
			uint64(1),
		)
	})

	t.Run("Should check that the MomentNFT is not in the users collection", func(t *testing.T) {
		// Get the MomentNFT data from the users collection.
		count := getMomentNFTLengthInAccount(t, b, contracts, userAddress)
		assert.Equal(t, big.NewInt(0), count)
	})

	t.Run("Should check that the entries length on Leaderboard is 1", func(t *testing.T) {
		count := getEscrowNFTLengthInLeaderboard(t, b, contracts, "leaderboardBurn-1")
		assert.Equal(t, big.NewInt(1), count)
	})

	t.Run("Should burn the moment in leaderboards", func(t *testing.T) {
		testBurnMomentNFT(
			t,
			b,
			contracts,
			"leaderboardBurn-1",
			uint64(1),
		)
	})

	t.Run("Should check that the entries length on Leaderboard is 1", func(t *testing.T) {
		count := getEscrowNFTLengthInLeaderboard(t, b, contracts, "leaderboardBurn-1")
		assert.Equal(t, big.NewInt(0), count)
	})

	t.Run("Should check that the MomentNFT is not in the users collection", func(t *testing.T) {
		// Get the MomentNFT data from the users collection.
		count := getMomentNFTLengthInAccount(t, b, contracts, userAddress)
		assert.Equal(t, big.NewInt(0), count)
	})
}

func testCreateLeaderboard(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	leaderboardName string,
) {
	createLeaderboard(
		t,
		b,
		contracts,
		leaderboardName,
	)
}

func testWithdrawMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	leaderboardName string,
	momentNftFlowID uint64,
) {
	withdrawMomentNFT(
		t,
		b,
		contracts,
		leaderboardName,
		momentNftFlowID,
	)
}

func testBurnMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	leaderboardName string,
	momentNftFlowID uint64,
) {
	burnMomentNFT(
		t,
		b,
		contracts,
		leaderboardName,
		momentNftFlowID,
	)
}

func testEscrowMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	userSigner crypto.Signer,
	userAddress flow.Address,
	momentNftFlowID uint64,
) {
	escrowMomentNFT(
		t,
		b,
		contracts,
		userSigner,
		userAddress,
		momentNftFlowID,
	)
}

func testGetLeaderboard(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	leaderboardName string,
) {
	getLeaderboard(
		t,
		b,
		contracts,
		leaderboardName,
	)
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}

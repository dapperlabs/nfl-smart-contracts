package test

import (
	"testing"

	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/assert"
)

//------------------------------------------------------------
// Setup
//------------------------------------------------------------
func TestShowdownDeployContracts(t *testing.T) {
	b := newEmulator()
	showdownDeployContracts(t, b)
}

func TestShowdownSetupAccount(t *testing.T) {
	b := newEmulator()
	contracts := showdownDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupShowdown(t, b, userAddress, userSigner, contracts)

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

//------------------------------------------------------------
// Series
//------------------------------------------------------------
func TestSeries(t *testing.T) {
	b := newEmulator()
	contracts := showdownDeployContracts(t, b)
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
	shouldBeID uint32,
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
	seriesID uint32,
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

//------------------------------------------------------------
// Sets
//------------------------------------------------------------
func TestSets(t *testing.T) {
	b := newEmulator()
	contracts := showdownDeployContracts(t, b)
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
	shouldBeID uint32,
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

//------------------------------------------------------------
// Plays
//------------------------------------------------------------
func TestPlays(t *testing.T) {
	b := newEmulator()
	contracts := showdownDeployContracts(t, b)
	createTestPlays(t, b, contracts)
}

func createTestPlays(t *testing.T, b *emulator.Blockchain, contracts Contracts) {
	t.Run("Should be able to create a new play", func(t *testing.T) {
		testCreatePlay(
			t,
			b,
			contracts,
			"TEST_CLASSIFICATION",
			map[string]string{"test play metadata a": "TEST PLAY METADATA A"},
			1,
			false,
		)
	})

	t.Run("Should be able to create a new play with an incrementing ID from the first", func(t *testing.T) {
		testCreatePlay(
			t,
			b,
			contracts,
			"TEST_CLASSIFICATION",
			map[string]string{"test play metadata b": "TEST PLAY METADATA B"},
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
	shouldBeID uint32,
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

//------------------------------------------------------------
// Editions
//------------------------------------------------------------
func TestEditions(t *testing.T) {
	b := newEmulator()
	contracts := showdownDeployContracts(t, b)
	createTestEditions(t, b, contracts)
}

func testCreateEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	seriesID uint32,
	setID uint32,
	playID uint32,
	maxMintSize *uint32,
	tier string,
	shouldBeID uint32,
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
	editionID uint32,
	shouldBeID uint32,
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
	var maxMintSize uint32 = 2
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
	contracts := showdownDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupShowdown(t, b, userAddress, userSigner, contracts)

	createTestEditions(t, b, contracts)

	t.Run("Should be able to mint a new MomentNFT from an edition that has a maxMintSize", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint32(1),
			userAddress,
			uint64(1),
			uint32(1),
			false,
		)
	})

	t.Run("Should be able to mint a second new MomentNFT from an edition that has a maxmintSize", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint32(1),
			userAddress,
			uint64(2),
			uint32(2),
			false,
		)
	})

	t.Run("Should be able to mint a new MomentNFT from an edition with no max mint size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint32(2),
			userAddress,
			uint64(3),
			uint32(1),
			false,
		)
	})

	t.Run("Should be able to mint a second new MomentNFT from an edition with no max mint size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint32(2),
			userAddress,
			uint64(4),
			uint32(2),
			false,
		)
	})

	t.Run("Should not be able to mint an edition that has reached max minting size", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint32(1),
			userAddress,
			uint64(3),
			uint32(3),
			true,
		)
	})

	t.Run("Should not be able to mint an edition that is already closed", func(t *testing.T) {
		testMintMomentNFT(
			t,
			b,
			contracts,
			uint32(3),
			userAddress,
			uint64(1),
			uint32(1),
			true,
		)
	})
}

func testMintMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionID uint32,
	userAddress flow.Address,
	shouldBeID uint64,
	shouldBeSerialNumber uint32,
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

//------------------------------------------------------------
// ShardedCollection
// ------------------------------------------------------------
func TestShardedCollection(t *testing.T) {
	b := newEmulator()

	contracts := showdownDeployContracts(t, b)
	// Set up for NFT Minting
	createTestEditions(t, b, contracts)
	user1Address, user1Signer := createAccount(t, b)
	user2Address, user2Signer := createAccount(t, b)
	numMomentNFTs := uint64(20)

	t.Run("Creating a sharded collection should work", func(t *testing.T) {
		setupShardedCollection(
			t,
			b,
			contracts,
			user1Address,
			user1Signer,
			uint64(75),
			false,
		)
	})

	t.Run("Creating a sharded collection twice for the same address should not work", func(t *testing.T) {
		setupShardedCollection(
			t,
			b,
			contracts,
			user1Address,
			user1Signer,
			uint64(75),
			true,
		)
	})

	t.Run("Minting to a sharded collection should work", func(t *testing.T) {
		numMomentNFTsAlready := getMomentNFTSupply(t, b, contracts)
		for i := uint64(1); i < numMomentNFTs; i++ {
			testMintMomentNFT(
				t,
				b,
				contracts,
				uint32(2),
				user1Address,
				uint64(numMomentNFTsAlready+i),
				uint32(numMomentNFTsAlready+i),
				false,
			)
		}
	})

	t.Run("Transferring from a sharded collection to a collection should work", func(t *testing.T) {
		setupShowdown(t, b, user2Address, user2Signer, contracts)
		// Transfer the first 10 moments from ShardedCollection to Collection
		for i := uint64(1); i <= 10; i++ {
			transferMomentNFTFromShardedCollection(
				t,
				b,
				contracts,
				user1Address,
				user1Signer,
				i,
				user2Address,
				false,
			)
		}
	})

	t.Run("Batch transferring from a sharded collection to a collection should work", func(t *testing.T) {
		user2Address, user2Signer := createAccount(t, b)
		setupShowdown(t, b, user2Address, user2Signer, contracts)

		// Make the list of IDs to transfer
		nftIDs := []uint64{}
		for i := uint64(11); i < 20; i++ {
			nftIDs = append(nftIDs, i)
		}
		// Transfer the next 10 moments from ShardedCollection to Collection
		batchTransferMomentNFTsFromShardedCollection(
			t,
			b,
			contracts,
			user1Address,
			user1Signer,
			nftIDs,
			user2Address,
			false,
		)
	})

	t.Run("Transferring from a collection to a collection should work", func(t *testing.T) {
		user3Address, user3Signer := createAccount(t, b)
		setupShowdown(t, b, user3Address, user3Signer, contracts)

		// Transfer the first 10 moments from Collection to Collection
		for i := uint64(1); i <= 10; i++ {
			transferMomentNFT(
				t,
				b,
				contracts,
				user2Address,
				user2Signer,
				i,
				user3Address,
				false,
			)
		}
	})
}

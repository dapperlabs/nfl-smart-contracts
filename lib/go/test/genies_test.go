package test

import (
	"testing"

	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
	"github.com/stretchr/testify/assert"
)

func testAdvanceSeries(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	seriesName string,
	metadata map[string]string,
	shouldBeID uint32,
	shouldRevert bool,
) {
	advanceSeries(
		t,
		b,
		contracts,
		seriesName,
		metadata,
		false,
	)

	if !shouldRevert {
		series := getCurrentSeriesData(t, b, contracts)
		assert.Equal(t, shouldBeID, series.ID)
		assert.Equal(t, seriesName, series.Name)
		assert.Equal(t, metadata, series.Metadata)
		assert.Equal(t, true, series.Active)
		assert.Equal(t, []uint32{}, series.CollectionIDs)
		assert.Equal(t, uint32(0), series.CollectionsOpen)
	}
}

func testAddGeniesCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	collectionName string,
	collectionSeriesID uint32,
	metadata map[string]string,
	shouldBeID uint32,
	shouldRevert bool,
) {
	addGeniesCollection(
		t,
		b,
		contracts,
		collectionName,
		collectionSeriesID,
		metadata,
		shouldRevert,
	)

	if !shouldRevert {
		collection := getGeniesCollectionData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, collection.ID)
		assert.Equal(t, collectionSeriesID, collection.SeriesID)
		assert.Equal(t, collectionName, collection.Name)
		assert.Equal(t, metadata, collection.Metadata)
		assert.Equal(t, true, collection.Open)
		assert.Equal(t, []uint32{}, collection.EditionIDs)
		assert.Equal(t, uint32(0), collection.EditionsActive)
	}
}

func testCloseGeniesCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	collectionID uint32,
	shouldRevert bool,
) {
	wasOpen := getGeniesCollectionData(t, b, contracts, collectionID).Open
	closeGeniesCollection(
		t,
		b,
		contracts,
		collectionID,
		shouldRevert,
	)

	collection := getGeniesCollectionData(t, b, contracts, collectionID)
	assert.Equal(t, collectionID, collection.ID)
	if !shouldRevert {
		assert.Equal(t, false, collection.Open)
	} else {
		assert.Equal(t, wasOpen, collection.Open)
	}
}

func testAddEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionName string,
	collectionID uint32,
	metadata map[string]string,
	shouldBeID uint32,
	shouldRevert bool,
) {
	addEdition(
		t,
		b,
		contracts,
		editionName,
		collectionID,
		metadata,
		shouldRevert,
	)

	if !shouldRevert {
		edition := getEditionData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, edition.ID)
		assert.Equal(t, collectionID, edition.CollectionID)
		assert.Equal(t, editionName, edition.Name)
		assert.Equal(t, metadata, edition.Metadata)
		assert.Equal(t, true, edition.Open)
		assert.Equal(t, uint32(0), edition.NumMinted)
	}
}

func testRetireEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionID uint32,
	shouldRevert bool,
) {
	editionWasOpen := getEditionData(t, b, contracts, editionID).Open
	retireEdition(
		t,
		b,
		contracts,
		editionID,
		shouldRevert,
	)

	edition := getEditionData(t, b, contracts, editionID)
	assert.Equal(t, editionID, edition.ID)
	if !shouldRevert {
		assert.Equal(t, false, edition.Open)
	} else {
		assert.Equal(t, editionWasOpen, edition.Open)
	}
}

func testMintGeniesNFT(
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
	previousSupply := getGeniesSupply(t, b, contracts)

	mintGeniesNFT(
		t,
		b,
		contracts,
		userAddress,
		editionID,
		shouldRevert,
	)

	newSupply := getGeniesSupply(t, b, contracts)

	if !shouldRevert {
		assert.Equal(t, previousSupply+uint64(1), newSupply)

		nftProperties := getGeniesNFTProperties(
			t,
			b,
			contracts,
			userAddress,
			shouldBeID,
		)
		assert.Equal(t, shouldBeID, nftProperties.ID)
		assert.Equal(t, editionID, nftProperties.EditionID)
		assert.Equal(t, shouldBeSerialNumber, nftProperties.SerialNumber)
		//FIXME: query the block time and check equality.
		//       Here we just make sure it's nonzero.
		assert.Less(t, uint64(0), nftProperties.MintingDate)
	} else {
		assert.Equal(t, previousSupply, newSupply)
	}
}

func TestGeniesDeployContracts(t *testing.T) {
	b := newEmulator()
	geniesDeployContracts(t, b)
}

func TestGeniesSetupAccount(t *testing.T) {
	b := newEmulator()

	contracts := geniesDeployContracts(t, b)

	userAddress, userSigner := createAccount(t, b)
	setupGenies(t, b, userAddress, userSigner, contracts)
}

func TestGeniesAdvanceSeries(t *testing.T) {
	b := newEmulator()

	contracts := geniesDeployContracts(t, b)

	t.Run("Should be able to create the first series", func(t *testing.T) {
		// Create the first series
		testAdvanceSeries(
			t,
			b,
			contracts,
			"Series One",
			map[string]string{"test a": "TEST A"},
			0,
			false,
		)
	})

	t.Run("Should be able to advance to the next series", func(t *testing.T) {
		// Create the next series
		testAdvanceSeries(
			t,
			b,
			contracts,
			"Series Two",
			map[string]string{"test b": "TEST B"},
			1,
			false,
		)

		// Check the previous series
		previousSeries := getSeriesData(t, b, contracts, 0)
		assert.Equal(t, uint32(0), previousSeries.ID)
		assert.Equal(t, "Series One", previousSeries.Name)
		assert.Equal(t, false, previousSeries.Active)
	})
}

func TestGeniesCollection(t *testing.T) {
	b := newEmulator()

	contracts := geniesDeployContracts(t, b)

	// Create the series that we will create collections in
	testAdvanceSeries(
		t,
		b,
		contracts,
		"Series One",
		map[string]string{"test c": "TEST C"},
		uint32(0),
		false,
	)

	t.Run("Should be able to add a collection to the series", func(t *testing.T) {
		testAddGeniesCollection(
			t,
			b,
			contracts,
			"Collection One",
			uint32(0),
			map[string]string{"test d": "TEST D"},
			uint32(0),
			false,
		)
	})

	t.Run("Should be able to add another collection to the series", func(t *testing.T) {
		testAddGeniesCollection(
			t,
			b,
			contracts,
			"Collection Two",
			uint32(0),
			map[string]string{"test e": "TEST E"},
			uint32(1),
			false,
		)
	})

	t.Run("Should be able to close a collection", func(t *testing.T) {
		testCloseGeniesCollection(
			t,
			b,
			contracts,
			uint32(0),
			false,
		)
	})

	t.Run("Should not be able to add a collection to a closed series", func(t *testing.T) {
		testCloseGeniesCollection(
			t,
			b,
			contracts,
			uint32(1),
			false,
		)

		testAdvanceSeries(
			t,
			b,
			contracts,
			"Series Two",
			map[string]string{"test f": "TEST F"},
			1,
			false,
		)

		testAddGeniesCollection(
			t,
			b,
			contracts,
			"Collection Two",
			uint32(0),
			map[string]string{"test f": "TEST F"},
			2,
			true,
		)
	})

}

func TestGeniesEdition(t *testing.T) {
	b := newEmulator()

	contracts := geniesDeployContracts(t, b)
	testAdvanceSeries(
		t,
		b,
		contracts,
		"Series One",
		map[string]string{"test g": "TEST G"},
		uint32(0),
		false,
	)

	testAddGeniesCollection(
		t,
		b,
		contracts,
		"Collection One",
		uint32(0),
		map[string]string{"test h": "TEST H"},
		uint32(0),
		false,
	)

	t.Run("Should be able to add an edition to the collection", func(t *testing.T) {
		testAddEdition(
			t,
			b,
			contracts,
			"GeniesNFT Line One",
			uint32(0),
			map[string]string{"test i": "TEST I"},
			uint32(0),
			false,
		)
	})

	t.Run("Should be able to add another edition to the collection", func(t *testing.T) {
		testAddEdition(
			t,
			b,
			contracts,
			"GeniesNFT Line Two",
			uint32(0),
			map[string]string{"test j": "TEST J"},
			uint32(1),
			false,
		)
	})

	t.Run("Should be able to retire an edition", func(t *testing.T) {
		testRetireEdition(
			t,
			b,
			contracts,
			uint32(0),
			false,
		)
	})

	t.Run("Should not be able to retire an already retired edition", func(t *testing.T) {
		testRetireEdition(
			t,
			b,
			contracts,
			uint32(0),
			true,
		)
	})

	t.Run("Should not be able to add an edition to a closed collection", func(t *testing.T) {
		collectionID := uint32(0)

		// Retire the other edition we created in tests, then close the collection
		retireEdition(t, b, contracts, 1, false)
		closeGeniesCollection(t, b, contracts, collectionID, false)

		testAddEdition(
			t,
			b,
			contracts,
			"GeniesNFT Line Three",
			uint32(0),
			map[string]string{"test k": "TEST K"},
			uint32(2),
			true,
		)
	})
}

func TestGeniesNFT(t *testing.T) {
	b := newEmulator()

	contracts := geniesDeployContracts(t, b)

	testAdvanceSeries(
		t,
		b,
		contracts,
		"Series One",
		map[string]string{"test l": "TEST l"},
		uint32(0),
		false,
	)

	testAddGeniesCollection(
		t,
		b,
		contracts,
		"Collection One",
		uint32(0),
		map[string]string{"test aa": "TEST AA"},
		uint32(0),
		false,
	)

	testAddEdition(
		t,
		b,
		contracts,
		"GeniesNFT Line One",
		uint32(0),
		map[string]string{"test bb": "TEST BB"},
		uint32(0),
		false,
	)

	testAddEdition(
		t,
		b,
		contracts,
		"GeniesNFT Line Two",
		uint32(0),
		map[string]string{"test cc": "TEST CC"},
		uint32(1),
		false,
	)

	userAddress, userSigner := createAccount(t, b)
	setupGenies(t, b, userAddress, userSigner, contracts)

	t.Run("Should be able to mint an NFT in an open edition", func(t *testing.T) {
		testMintGeniesNFT(
			t,
			b,
			contracts,
			uint32(0),
			userAddress,
			uint64(0),
			uint32(0),
			false,
		)
	})

	t.Run("Should be able to mint another NFT in an open edition", func(t *testing.T) {
		testMintGeniesNFT(
			t,
			b,
			contracts,
			uint32(0),
			userAddress,
			uint64(1),
			uint32(1),
			false,
		)
	})

	// Check nonzero/different values in the report struct.

	t.Run("Should be able to mint an NFT in another open edition", func(t *testing.T) {
		testMintGeniesNFT(
			t,
			b,
			contracts,
			uint32(1),
			userAddress,
			uint64(2),
			uint32(0),
			false,
		)
	})

	t.Run("Should be able to mint another NFT in another open edition", func(t *testing.T) {
		testMintGeniesNFT(
			t,
			b,
			contracts,
			uint32(1),
			userAddress,
			uint64(3),
			uint32(1),
			false,
		)
	})

	t.Run("Should not be able to mint an NFT in an closed edition", func(t *testing.T) {
		editionID := uint32(0)

		retireEdition(t, b, contracts, editionID, false)

		testMintGeniesNFT(
			t,
			b,
			contracts,
			uint32(0),
			userAddress,
			uint64(2),
			uint32(2),
			true,
		)
	})
}

func TestShardedCollection(t *testing.T) {
	b := newEmulator()

	contracts := geniesDeployContracts(t, b)

	// Set up the NFT minting
	testAdvanceSeries(
		t,
		b,
		contracts,
		"Series One",
		map[string]string{"test m": "TEST M"},
		uint32(0),
		false,
	)
	testAddGeniesCollection(
		t,
		b,
		contracts,
		"Collection One",
		uint32(0),
		map[string]string{"test n": "TEST N"},
		uint32(0),
		false,
	)
	testAddEdition(
		t,
		b,
		contracts,
		"GeniesNFT Line One",
		uint32(0),
		map[string]string{"test o": "TEST O"},
		uint32(0),
		false,
	)

	t.Run("Creating a sharded collection should work", func(t *testing.T) {
		user2Address, user2Signer := createAccount(t, b)
		setupShardedCollection(
			t,
			b,
			contracts,
			user2Address,
			user2Signer,
			uint64(75),
			false,
		)
	})

	t.Run("Creating a sharded collection twice for the same address should  notwork", func(t *testing.T) {
		user2Address, user2Signer := createAccount(t, b)
		setupShardedCollection(
			t,
			b,
			contracts,
			user2Address,
			user2Signer,
			uint64(75),
			false,
		)
		setupShardedCollection(
			t,
			b,
			contracts,
			user2Address,
			user2Signer,
			uint64(75),
			true,
		)
	})

	t.Run("Minting to a sharded collection should work", func(t *testing.T) {
		numGeniesNFTs := uint64(10)
		numGeniesNFTsAlready := getGeniesSupply(t, b, contracts)
		user1Address, user1Signer := createAccount(t, b)
		setupShardedCollection(
			t,
			b,
			contracts,
			user1Address,
			user1Signer,
			uint64(75),
			false,
		)
		for i := uint64(0); i < numGeniesNFTs; i++ {
			testMintGeniesNFT(
				t,
				b,
				contracts,
				uint32(0),
				user1Address,
				uint64(numGeniesNFTsAlready+i),
				uint32(numGeniesNFTsAlready+i),
				false,
			)
		}
	})

	t.Run("Transferring from a sharded collection to a collection should work", func(t *testing.T) {
		numGeniesNFTs := uint64(10)
		numGeniesNFTsAlready := getGeniesSupply(t, b, contracts)
		user1Address, user1Signer := createAccount(t, b)
		setupShardedCollection(
			t,
			b,
			contracts,
			user1Address,
			user1Signer,
			uint64(75),
			false,
		)
		for i := uint64(0); i < numGeniesNFTs; i++ {
			testMintGeniesNFT(
				t,
				b,
				contracts,
				uint32(0),
				user1Address,
				uint64(numGeniesNFTsAlready+i),
				uint32(numGeniesNFTsAlready+i),
				false,
			)
		}

		user2Address, user2Signer := createAccount(t, b)
		setupGenies(t, b, user2Address, user2Signer, contracts)
		// Transfer from ShardedCollection to Collection
		for i := uint64(0); i < numGeniesNFTs; i++ {
			transferGeniesNFTFromShardedCollection(
				t,
				b,
				contracts,
				user1Address,
				user1Signer,
				numGeniesNFTsAlready+i,
				user2Address,
				false,
			)
		}
	})

	t.Run("Batch transferring from a sharded collection to a collection should work", func(t *testing.T) {
		numGeniesNFTs := uint64(10)
		numGeniesNFTsAlready := getGeniesSupply(t, b, contracts)
		user1Address, user1Signer := createAccount(t, b)
		setupShardedCollection(
			t,
			b,
			contracts,
			user1Address,
			user1Signer,
			uint64(75),
			false,
		)
		for i := uint64(0); i < numGeniesNFTs; i++ {
			testMintGeniesNFT(
				t,
				b,
				contracts,
				uint32(0),
				user1Address,
				uint64(numGeniesNFTsAlready+i),
				uint32(numGeniesNFTsAlready+i),
				false,
			)
		}

		user2Address, user2Signer := createAccount(t, b)
		setupGenies(t, b, user2Address, user2Signer, contracts)

		// Make the list of IDs to transfer
		nftIDs := make([]uint64, numGeniesNFTs)
		// Transfer from ShardedCollection to Collection
		for i := uint64(0); i < numGeniesNFTs; i++ {
			nftIDs[i] = numGeniesNFTsAlready + i
		}

		batchTransferGeniesNFTsFromShardedCollection(
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
}

package test

import (
	"testing"

	emulator "github.com/onflow/flow-emulator"
	"github.com/stretchr/testify/assert"
)

//------------------------------------------------------------
// Setup
//------------------------------------------------------------
func TestAllDaySeasonalDeployContracts(t *testing.T) {
	b := newEmulator()
	AllDaySeasonalDeployContracts(t, b)
}

func TestAllDaySeasonalSetupAccount(t *testing.T) {
	b := newEmulator()
	contracts := AllDaySeasonalDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDaySeasonal(t, b, userAddress, userSigner, contracts)

	t.Run("Account should be set up", func(t *testing.T) {
		accountIsSetUp := accountSeasonalIsSetup(
			t,
			b,
			contracts,
			userAddress,
		)
		assert.Equal(t, true, accountIsSetUp)
	})
}

//------------------------------------------------------------
// Edition
//------------------------------------------------------------
func TestEdition(t *testing.T) {
	b := newEmulator()
	contracts := AllDaySeasonalDeployContracts(t, b)
	createTestSeasonalEditions(t, b, contracts)
}

func createTestSeasonalEditions(t *testing.T, b *emulator.Blockchain, contracts Contracts) {
	t.Run("Should be able to create a new edition", func(t *testing.T) {
		testCreateSeasonalEdition(
			t,
			b,
			contracts,
			map[string]string{"test play metadata a": "TEST PLAY METADATA A"},
			1,
			false,
		)
	})

	t.Run("Should be able to create a new edition with an incrementing ID from the first", func(t *testing.T) {
		testCreateSeasonalEdition(
			t,
			b,
			contracts,
			map[string]string{"test play metadata a": "TEST PLAY METADATA A"},
			2,
			false,
		)
	})

	t.Run("Should be able to close a edition", func(t *testing.T) {
		testCloseSeasonalEdition(
			t,
			b,
			contracts,
			2,
			false,
		)
	})
}

func testCreateSeasonalEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	metadata map[string]string,
	shouldBeID uint64,
	shouldRevert bool,
) {
	createSeasonalEdition(
		t,
		b,
		contracts,
		metadata,
		false,
	)

	if !shouldRevert {
		series := getSeasonalEditionData(t, b, contracts, shouldBeID)
		assert.Equal(t, shouldBeID, series.ID)
		assert.Equal(t, true, series.Active)
	}
}

func testCloseSeasonalEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionID uint64,
	shouldRevert bool,
) {
	wasActive := getSeasonalEditionData(t, b, contracts, editionID).Active
	closeSeasonalEdition(
		t,
		b,
		contracts,
		editionID,
		shouldRevert,
	)

	edition := getSeasonalEditionData(t, b, contracts, editionID)
	assert.Equal(t, editionID, edition.ID)
	if !shouldRevert {
		assert.Equal(t, false, edition.Active)
	} else {
		assert.Equal(t, wasActive, edition.Active)
	}
}

// ------------------------------------------------------------
// MomentNFTs
// ------------------------------------------------------------
func TestSeasonalNFTs(t *testing.T) {
	b := newEmulator()
	contracts := AllDaySeasonalDeployContracts(t, b)
	userAddress, userSigner := createAccount(t, b)
	setupAllDaySeasonal(t, b, userAddress, userSigner, contracts)

	createTestSeasonalEditions(t, b, contracts)

	t.Run("Should be able to mint a new NFT from an edition that has a maxMintSize", func(t *testing.T) {
		testMintSeasonalNFT(
			t,
			b,
			contracts,
			uint64(1),
			userAddress,
			uint64(1),
			false,
		)
	})

	t.Run("Should be able to mint a second new MomentNFT from an edition that has a maxmintSize", func(t *testing.T) {
		testMintSeasonalNFT(
			t,
			b,
			contracts,
			uint64(1),
			userAddress,
			uint64(2),
			false,
		)
	})

	closeSeasonalEdition(
		t,
		b,
		contracts,
		uint64(1),
		false,
	)

	t.Run("Should not be able to mint an edition that is already closed", func(t *testing.T) {
		testMintSeasonalNFT(
			t,
			b,
			contracts,
			uint64(1),
			userAddress,
			uint64(3),
			true,
		)
	})
}

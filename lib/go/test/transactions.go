package test

import (
	"testing"

	"github.com/onflow/cadence"
	emulator "github.com/onflow/flow-emulator"
	fttemplates "github.com/onflow/flow-ft/lib/go/templates"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

//------------------------------------------------------------
// Setup
//------------------------------------------------------------
func fundAccount(
	t *testing.T,
	b *emulator.Blockchain,
	receiverAddress flow.Address,
	amount string,
) {
	script := fttemplates.GenerateMintTokensScript(
		ftAddress,
		flowTokenAddress,
		flowTokenName,
	)

	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(b.ServiceKey().Address)

	tx.AddArgument(cadence.NewAddress(receiverAddress))
	tx.AddArgument(cadenceUFix64(amount))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address},
		[]crypto.Signer{b.ServiceKey().Signer()},
		false,
	)
}

/*
	NOTE: This requires extra storage, and higher gas.
*/

func setupShardedCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	recipientAddress flow.Address,
	recipientSigner crypto.Signer,
	numberOfBuckets uint64,
	shouldRevert bool,
) {
	// We need additional storage to hold the shards
	fundAccount(
		t,
		b,
		recipientAddress,
		"0.01",
	)

	tx := flow.NewTransaction().
		SetScript(loadShowdownSetupShardedCollectionTransaction(contracts)).
		SetGasLimit(800).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(recipientAddress)
	tx.AddArgument(cadence.NewUInt64(numberOfBuckets))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, recipientAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), recipientSigner},
		shouldRevert,
	)
}

func transferMomentNFTFromShardedCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	senderAddress flow.Address,
	senderSigner crypto.Signer,
	nftID uint64,
	recipientAddress flow.Address,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownTransferMomentNFTFromShardedCollectionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(senderAddress)
	tx.AddArgument(cadence.BytesToAddress(recipientAddress.Bytes()))
	tx.AddArgument(cadence.NewUInt64(nftID))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, senderAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), senderSigner},
		shouldRevert,
	)
}

/*
	NOTE: This requires higher gas.
*/

func batchTransferMomentNFTsFromShardedCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	senderAddress flow.Address,
	senderSigner crypto.Signer,
	nftIDs []uint64,
	recipientAddress flow.Address,
	shouldRevert bool,
) {
	cadenceIDs := make([]cadence.Value, len(nftIDs))
	for i, id := range nftIDs {
		cadenceIDs[i] = cadence.NewUInt64(id)
	}
	tx := flow.NewTransaction().
		SetScript(loadShowdownBatchTransferMomentNFTsFromShardedCollectionTransaction(contracts)).
		SetGasLimit(uint64(100+(50*len(nftIDs)))).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(senderAddress)
	tx.AddArgument(cadence.BytesToAddress(recipientAddress.Bytes()))
	tx.AddArgument(cadence.NewArray(cadenceIDs))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, senderAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), senderSigner},
		shouldRevert,
	)
}

//------------------------------------------------------------
// Series
//------------------------------------------------------------
func createSeries(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	name string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownCreateSeriesTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.NewString(name))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

func closeSeries(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint32,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownCloseSeriesTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.NewUInt32(id))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

//------------------------------------------------------------
// Sets
//------------------------------------------------------------
func createSet(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	name string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownCreateSetTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.NewString(name))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

//------------------------------------------------------------
// Plays
//------------------------------------------------------------
func createPlay(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	classification string,
	metadata map[string]string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownCreatePlayTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.NewString(classification))
	tx.AddArgument(metadataDict(metadata))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

//------------------------------------------------------------
// Editions
//------------------------------------------------------------
func createEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	seriesID uint32,
	setID uint32,
	playID uint32,
	maxMintSize *uint32,
	tier string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownCreateEditionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.NewUInt32(seriesID))
	tx.AddArgument(cadence.NewUInt32(setID))
	tx.AddArgument(cadence.NewUInt32(playID))
	tx.AddArgument(cadence.NewString(tier))
	if maxMintSize != nil {
		tx.AddArgument(cadence.NewUInt32(*maxMintSize))
	} else {
		tx.AddArgument(cadence.Optional{})
	}

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

func closeEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	editionID uint32,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownCloseEditionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.NewUInt32(editionID))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

//------------------------------------------------------------
// MomentNFTs
//------------------------------------------------------------
func mintMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	recipientAddress flow.Address,
	editionID uint32,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownMintMomentNFTTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.ShowdownAddress)
	tx.AddArgument(cadence.BytesToAddress(recipientAddress.Bytes()))
	tx.AddArgument(cadence.NewUInt32(editionID))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.ShowdownAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.ShowdownSigner},
		shouldRevert,
	)
}

func transferMomentNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	senderAddress flow.Address,
	senderSigner crypto.Signer,
	nftID uint64,
	recipientAddress flow.Address,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadShowdownTransferNFTTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(senderAddress)
	tx.AddArgument(cadence.BytesToAddress(recipientAddress.Bytes()))
	tx.AddArgument(cadence.NewUInt64(nftID))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, senderAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), senderSigner},
		shouldRevert,
	)
}

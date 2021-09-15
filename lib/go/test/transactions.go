package test

import (
	"testing"

	"github.com/onflow/cadence"
	emulator "github.com/onflow/flow-emulator"
	fttemplates "github.com/onflow/flow-ft/lib/go/templates"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

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

func advanceSeries(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	name string,
	metadata map[string]string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadGeniesAdvanceSeriesTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.GeniesAddress)
	tx.AddArgument(cadence.NewString(name))
	tx.AddArgument(metadataDict(metadata))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.GeniesAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.GeniesSigner},
		shouldRevert,
	)
}

func addGeniesCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	name string,
	seriesID uint32,
	metadata map[string]string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadGeniesAddGeniesCollectionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.GeniesAddress)
	tx.AddArgument(cadence.NewString(name))
	tx.AddArgument(cadence.NewUInt32(seriesID))
	tx.AddArgument(metadataDict(metadata))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.GeniesAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.GeniesSigner},
		shouldRevert,
	)
}

func closeGeniesCollection(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint32,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadGeniesCloseGeniesCollectionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.GeniesAddress)
	tx.AddArgument(cadence.NewUInt32(id))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.GeniesAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.GeniesSigner},
		shouldRevert,
	)
}

func addEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	name string,
	collectionID uint32,
	metadata map[string]string,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadGeniesAddEditionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.GeniesAddress)
	tx.AddArgument(cadence.NewString(name))
	tx.AddArgument(cadence.NewUInt32(collectionID))
	tx.AddArgument(metadataDict(metadata))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.GeniesAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.GeniesSigner},
		shouldRevert,
	)
}

func retireEdition(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	id uint32,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadGeniesRetireEditionTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.GeniesAddress)
	tx.AddArgument(cadence.NewUInt32(id))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.GeniesAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.GeniesSigner},
		shouldRevert,
	)
}

func mintGeniesNFT(
	t *testing.T,
	b *emulator.Blockchain,
	contracts Contracts,
	recipientAddress flow.Address,
	editionID uint32,
	shouldRevert bool,
) {
	tx := flow.NewTransaction().
		SetScript(loadGeniesMintGeniesNFTTransaction(contracts)).
		SetGasLimit(100).
		SetProposalKey(b.ServiceKey().Address, b.ServiceKey().Index, b.ServiceKey().SequenceNumber).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(contracts.GeniesAddress)
	tx.AddArgument(cadence.BytesToAddress(recipientAddress.Bytes()))
	tx.AddArgument(cadence.NewUInt32(editionID))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{b.ServiceKey().Address, contracts.GeniesAddress},
		[]crypto.Signer{b.ServiceKey().Signer(), contracts.GeniesSigner},
		shouldRevert,
	)
}

func transferGeniesNFT(
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
		SetScript(loadGeniesTransferNFTTransaction(contracts)).
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
		SetScript(loadGeniesSetupShardedCollectionTransaction(contracts)).
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

func transferGeniesNFTFromShardedCollection(
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
		SetScript(loadGeniesTransferGeniesNFTFromShardedCollectionTransaction(contracts)).
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

func batchTransferGeniesNFTsFromShardedCollection(
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
		SetScript(loadGeniesBatchTransferGeniesNFTsFromShardedCollectionTransaction(contracts)).
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

package test

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

// Handle relative paths by making these regular expressions

const (
	nftAddressPlaceholder                       = "\"[^\"]*NonFungibleToken.cdc\""
	showdownAddressPlaceholder                  = "\"[^\"]*Showdown.cdc\""
	showdownShardedCollectionAddressPlaceholder = "\"[^\"]*ShowdownShardedCollection.cdc\""
)

const (
	showdownPath                 = "../../../contracts/Showdown.cdc"
	showdownTransactionsRootPath = "../../../transactions"
	showdownScriptsRootPath      = "../../../scripts"

	// Accounts
	showdownSetupAccountPath   = showdownTransactionsRootPath + "/user/setup_showdown_account.cdc"
	showdownAccountIsSetupPath = showdownScriptsRootPath + "/user/account_is_setup.cdc"

	// ShardedCollection
	showdownShardedCollectionPath                            = "../../../contracts/ShowdownShardedCollection.cdc"
	showdownSetupShardedCollectionPath                       = showdownTransactionsRootPath + "/admin/sharded_collection/setup_sharded_collection.cdc"
	showdownTransferMomentNFTFromShardedCollectionPath       = showdownTransactionsRootPath + "/admin/sharded_collection/transfer_showdown_nft_from_sharded_collection.cdc"
	showdownBatchTransferMomentNFTsFromShardedCollectionPath = showdownTransactionsRootPath + "/admin/sharded_collection/batch_transfer_showdown_nfts_from_sharded_collection.cdc"

	// Series
	showdownCreateSeriesPath       = showdownTransactionsRootPath + "/admin/series/create_series.cdc"
	showdownCloseSeriesPath        = showdownTransactionsRootPath + "/admin/series/close_series.cdc"
	showdownReadAllSeriesPath      = showdownScriptsRootPath + "/series/read_all_series.cdc"
	showdownReadSeriesByIDPath     = showdownScriptsRootPath + "/series/read_series_by_id.cdc"
	showdownReadSeriesByNamePath   = showdownScriptsRootPath + "/series/read_series_by_name.cdc"
	showdownReadAllSeriesNamesPath = showdownScriptsRootPath + "/series/read_all_series_names.cdc"

	// Sets
	showdownCreateSetPath       = showdownTransactionsRootPath + "/admin/sets/create_set.cdc"
	showdownReadAllSetsPath     = showdownScriptsRootPath + "/sets/read_all_sets.cdc"
	showdownReadSetByIDPath     = showdownScriptsRootPath + "/sets/read_set_by_id.cdc"
	showdownReadSetsByNamePath  = showdownScriptsRootPath + "/sets/read_sets_by_name.cdc"
	showdownReadAllSetNamesPath = showdownScriptsRootPath + "/sets/read_all_set_names.cdc"

	// Plays
	showdownCreatePlayPath   = showdownTransactionsRootPath + "/admin/plays/create_play.cdc"
	showdownReadPlayByIDPath = showdownScriptsRootPath + "/plays/read_play_by_id.cdc"
	showdownReadAllPlaysPath = showdownScriptsRootPath + "/plays/read_all_plays.cdc"

	// Editions
	showdownCreateEditionPath   = showdownTransactionsRootPath + "/admin/editions/create_edition.cdc"
	showdownCloseEditionPath    = showdownTransactionsRootPath + "/admin/editions/close_edition.cdc"
	showdownReadEditionByIDPath = showdownScriptsRootPath + "/editions/read_edition_by_id.cdc"
	showdownReadAllEditionsPath = showdownScriptsRootPath + "/edition/read_all_editions.cdc"

	// Moment NFTs
	showdownMintMomentNFTPath           = showdownTransactionsRootPath + "/admin/nfts/mint_moment_nft.cdc"
	showdownMintMomentNFTMultiPath      = showdownTransactionsRootPath + "/admin/nfts/mint_moment_nft_multi.cdc"
	showdownTransferNFTPath             = showdownTransactionsRootPath + "/user/transfer_moment_nft.cdc"
	showdownReadMomentNFTSupplyPath     = showdownScriptsRootPath + "/nfts/read_moment_nft_supply.cdc"
	showdownReadMomentNFTPropertiesPath = showdownScriptsRootPath + "/nfts/read_moment_nft_properties.cdc"
	showdownReadCollectionNFTLengthPath = showdownScriptsRootPath + "/nfts/read_collection_nft_length.cdc"
	showdownReadCollectionNFTIDsPath    = showdownScriptsRootPath + "/nfts/read_collection_nft_ids.cdc"
)

//------------------------------------------------------------
// Accounts
//------------------------------------------------------------
func replaceAddresses(code []byte, contracts Contracts) []byte {
	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+contracts.NFTAddress.String()))

	showdownRe := regexp.MustCompile(showdownAddressPlaceholder)
	code = showdownRe.ReplaceAll(code, []byte("0x"+contracts.ShowdownAddress.String()))

	showdownShardedCollectionRe := regexp.MustCompile(showdownShardedCollectionAddressPlaceholder)
	code = showdownShardedCollectionRe.ReplaceAll(code, []byte("0x"+contracts.ShowdownShardedCollectionAddress.String()))

	return code
}

func LoadShowdown(nftAddress flow.Address) []byte {
	code := readFile(showdownPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	return code
}

func loadShowdownSetupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownSetupAccountPath),
		contracts,
	)
}

func loadShowdownAccountIsSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownAccountIsSetupPath),
		contracts,
	)
}

//------------------------------------------------------------
// Sharded Collection
//------------------------------------------------------------
func loadShowdownShardedCollection(nftAddress flow.Address, showdownAddress flow.Address) []byte {
	code := readFile(showdownShardedCollectionPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	showdownRe := regexp.MustCompile(showdownAddressPlaceholder)
	code = showdownRe.ReplaceAll(code, []byte("0x"+showdownAddress.String()))

	return code
}

func loadShowdownSetupShardedCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownSetupShardedCollectionPath),
		contracts,
	)
}

func loadShowdownTransferMomentNFTFromShardedCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownTransferMomentNFTFromShardedCollectionPath),
		contracts,
	)
}

func loadShowdownBatchTransferMomentNFTsFromShardedCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownBatchTransferMomentNFTsFromShardedCollectionPath),
		contracts,
	)
}

//------------------------------------------------------------
// Series
//------------------------------------------------------------
func loadShowdownCreateSeriesTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownCreateSeriesPath),
		contracts,
	)
}

func loadShowdownReadSeriesByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadSeriesByIDPath),
		contracts,
	)
}

func loadShowdownReadSeriesByNameScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadSeriesByNamePath),
		contracts,
	)
}

func loadShowdownReadAllSeriesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadAllSeriesPath),
		contracts,
	)
}

func loadShowdownReadAllSeriesNamesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadAllSeriesNamesPath),
		contracts,
	)
}

func loadShowdownCloseSeriesTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownCloseSeriesPath),
		contracts,
	)
}

//------------------------------------------------------------
// Sets
//------------------------------------------------------------
func loadShowdownCreateSetTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownCreateSetPath),
		contracts,
	)
}

func loadShowdownReadSetByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadSetByIDPath),
		contracts,
	)
}

func loadShowdownReadAllSetsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadAllSetsPath),
		contracts,
	)
}

func loadShowdownReadSetsByNameScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadSetsByNamePath),
		contracts,
	)
}

func loadShowdownReadAllSetNamesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadAllSetNamesPath),
		contracts,
	)
}

//------------------------------------------------------------
// Plays
//------------------------------------------------------------
func loadShowdownCreatePlayTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownCreatePlayPath),
		contracts,
	)
}

func loadShowdownReadPlayByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadPlayByIDPath),
		contracts,
	)
}

func loadShowdownReadAllPlaysScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadAllPlaysPath),
		contracts,
	)
}

//------------------------------------------------------------
// Editions
//------------------------------------------------------------
func loadShowdownCreateEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownCreateEditionPath),
		contracts,
	)
}

func loadShowdownReadEditionByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadEditionByIDPath),
		contracts,
	)
}

func loadShowdownCloseEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownCloseEditionPath),
		contracts,
	)
}

func loadShowdownReadAllEditionsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadAllEditionsPath),
		contracts,
	)
}

//------------------------------------------------------------
// Moment NFTs
//------------------------------------------------------------
func loadShowdownMintMomentNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownMintMomentNFTPath),
		contracts,
	)
}

func loadShowdownMintMomentNFTMultiTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownMintMomentNFTMultiPath),
		contracts,
	)
}

func loadShowdownReadMomentNFTSupplyScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadMomentNFTSupplyPath),
		contracts,
	)
}

func loadShowdownReadMomentNFTPropertiesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadMomentNFTPropertiesPath),
		contracts,
	)
}

func loadShowdownReadCollectionNFTLengthScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadCollectionNFTLengthPath),
		contracts,
	)
}

func loadShowdownReadCollectionNFTIDsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownReadCollectionNFTIDsPath),
		contracts,
	)
}

func loadShowdownTransferNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(showdownTransferNFTPath),
		contracts,
	)
}

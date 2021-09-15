package test

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

// Handle relative paths by making these regular expressions

const (
	nftAddressPlaceholder                     = "\"[^\"]*NonFungibleToken.cdc\""
	geniesAddressPlaceholder                  = "\"[^\"]*Genies.cdc\""
	geniesShardedCollectionAddressPlaceholder = "\"[^\"]*GeniesShardedCollection.cdc\""
)

const (
	geniesPath                  = "../../../contracts/Genies.cdc"
	geniesShardedCollectionPath = "../../../contracts/GeniesShardedCollection.cdc"
	geniesTransactionsRootPath  = "../../../transactions"
	geniesScriptsRootPath       = "../../../scripts"

	geniesAddEditionPath                                   = geniesTransactionsRootPath + "/admin/edition/add_edition.cdc"
	geniesRetireEditionPath                                = geniesTransactionsRootPath + "/admin/edition/retire_edition.cdc"
	geniesAddGeniesCollectionPath                          = geniesTransactionsRootPath + "/admin/geniesCollection/add_genies_collection.cdc"
	geniesCloseGeniesCollectionPath                        = geniesTransactionsRootPath + "/admin/geniesCollection/close_genies_collection.cdc"
	geniesMintGeniesNFTMultiPath                           = geniesTransactionsRootPath + "/admin/nfts/mint_genies_nft_multi.cdc"
	geniesMintGeniesNFTPath                                = geniesTransactionsRootPath + "/admin/nfts/mint_genies_nft.cdc"
	geniesAdvanceSeriesPath                                = geniesTransactionsRootPath + "/admin/series/advance_series.cdc"
	geniesSetupAccountPath                                 = geniesTransactionsRootPath + "/user/setup_genies_account.cdc"
	geniesTransferNFTPath                                  = geniesTransactionsRootPath + "/user/transfer_genies_nft.cdc"
	geniesSetupShardedCollectionPath                       = geniesTransactionsRootPath + "/admin/sharded_collection/setup_sharded_collection.cdc"
	geniesTransferGeniesNFTFromShardedCollectionPath       = geniesTransactionsRootPath + "/admin/sharded_collection/transfer_genies_nft_from_sharded_collection.cdc"
	geniesBatchTransferGeniesNFTsFromShardedCollectionPath = geniesTransactionsRootPath + "/admin/sharded_collection/batch_transfer_genies_nfts_from_sharded_collection.cdc"

	geniesReadAllEditionsPath          = geniesScriptsRootPath + "/edition/read_all_editions.cdc"
	geniesReadEditionPath              = geniesScriptsRootPath + "/edition/read_edition.cdc"
	geniesReadAllGeniesCollectionsPath = geniesScriptsRootPath + "/geniesCollection/read_all_genies_collections.cdc"
	geniesReadGeniesCollectionPath     = geniesScriptsRootPath + "/geniesCollection/read_genies_collection.cdc"
	geniesReadCollectionIDsPath        = geniesScriptsRootPath + "/nfts/read_collection_ids.cdc"
	geniesReadCollectionLengthPath     = geniesScriptsRootPath + "/nfts/read_collection_length.cdc"
	geniesReadGeniesSupplyPath         = geniesScriptsRootPath + "/nfts/read_genies_supply.cdc"
	geniesReadNFTPropertiesPath        = geniesScriptsRootPath + "/nfts/read_nft_properties.cdc"
	geniesReadAllSeriesNamesPath       = geniesScriptsRootPath + "/series/read_all_series_names.cdc"
	geniesReadAllSeriesPath            = geniesScriptsRootPath + "/series/read_all_series.cdc"
	geniesReadCurrentSeriesPath        = geniesScriptsRootPath + "/series/read_current_series.cdc"
	geniesReadSeriesByIDPath           = geniesScriptsRootPath + "/series/read_series_by_id.cdc"
	geniesReadSeriesByNamePath         = geniesScriptsRootPath + "/series/read_series_by_name.cdc"
	geniesAccountIsSetupPath           = geniesScriptsRootPath + "/user/account_is_setup.cdc"
)

func replaceAddresses(code []byte, contracts Contracts) []byte {
	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+contracts.NFTAddress.String()))

	geniesRe := regexp.MustCompile(geniesAddressPlaceholder)
	code = geniesRe.ReplaceAll(code, []byte("0x"+contracts.GeniesAddress.String()))

	geniesShardedCollectionRe := regexp.MustCompile(geniesShardedCollectionAddressPlaceholder)
	code = geniesShardedCollectionRe.ReplaceAll(code, []byte("0x"+contracts.GeniesShardedCollectionAddress.String()))

	return code
}

func loadGenies(nftAddress flow.Address) []byte {
	code := readFile(geniesPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	return code
}

func loadGeniesShardedCollection(nftAddress flow.Address, geniesAddress flow.Address) []byte {
	code := readFile(geniesShardedCollectionPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	geniesRe := regexp.MustCompile(geniesAddressPlaceholder)
	code = geniesRe.ReplaceAll(code, []byte("0x"+geniesAddress.String()))

	return code
}

func loadGeniesAddEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesAddEditionPath),
		contracts,
	)
}

func loadGeniesRetireEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesRetireEditionPath),
		contracts,
	)
}

func loadGeniesAddGeniesCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesAddGeniesCollectionPath),
		contracts,
	)
}

func loadGeniesCloseGeniesCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesCloseGeniesCollectionPath),
		contracts,
	)
}

func loadGeniesMintGeniesNFTMultiTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesMintGeniesNFTMultiPath),
		contracts,
	)
}

func loadGeniesMintGeniesNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesMintGeniesNFTPath),
		contracts,
	)
}

func loadGeniesAdvanceSeriesTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesAdvanceSeriesPath),
		contracts,
	)
}

func loadGeniesSetupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesSetupAccountPath),
		contracts,
	)
}

func loadGeniesTransferNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesTransferNFTPath),
		contracts,
	)
}

func loadGeniesSetupShardedCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesSetupShardedCollectionPath),
		contracts,
	)
}

func loadGeniesTransferGeniesNFTFromShardedCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesTransferGeniesNFTFromShardedCollectionPath),
		contracts,
	)
}

func loadGeniesBatchTransferGeniesNFTsFromShardedCollectionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesBatchTransferGeniesNFTsFromShardedCollectionPath),
		contracts,
	)
}

func loadGeniesReadAllEditionsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadAllEditionsPath),
		contracts,
	)
}

func loadGeniesReadEditionScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadEditionPath),
		contracts,
	)
}

func loadGeniesReadAllGeniesCollectionsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadAllGeniesCollectionsPath),
		contracts,
	)
}

func loadGeniesReadGeniesCollectionScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadGeniesCollectionPath),
		contracts,
	)
}

func loadGeniesReadCollectionIDsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadCollectionIDsPath),
		contracts,
	)
}

func loadGeniesReadCollectionLengthScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadCollectionLengthPath),
		contracts,
	)
}

func loadGeniesReadGeniesSupplyScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadGeniesSupplyPath),
		contracts,
	)
}
func loadGeniesReadNFTPropertiesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadNFTPropertiesPath),
		contracts,
	)
}

func loadGeniesReadAllSeriesNamesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadAllSeriesNamesPath),
		contracts,
	)
}

func loadGeniesReadAllSeriesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadAllSeriesPath),
		contracts,
	)
}

func loadGeniesReadCurrentSeriesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadCurrentSeriesPath),
		contracts,
	)
}

func loadGeniesReadSeriesByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadSeriesByIDPath),
		contracts,
	)
}

func loadGeniesReadSeriesByNameScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesReadSeriesByNamePath),
		contracts,
	)
}

func loadGeniesAccountIsSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(geniesAccountIsSetupPath),
		contracts,
	)
}

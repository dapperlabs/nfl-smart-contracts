package test

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

// Handle relative paths by making these regular expressions

const (
	nftAddressPlaceholder            = "\"[^\"]*NonFungibleToken.cdc\""
	AllDayAddressPlaceholder         = "\"[^\"]*AllDay.cdc\""
	AllDaySeasonalAddressPlaceholder = "\"[^\"]*AllDaySeasonal.cdc\""
)

const (
	AllDayPath                 = "../../../contracts/AllDay.cdc"
	AllDaySeasonalPath         = "../../../contracts/AllDaySeasonal.cdc"
	AllDayTransactionsRootPath = "../../../transactions"
	AllDayScriptsRootPath      = "../../../scripts"

	// Accounts
	AllDaySetupAccountPath           = AllDayTransactionsRootPath + "/user/setup_allday_account.cdc"
	AllDaySeasonalSetupAccountPath   = AllDayTransactionsRootPath + "/user/setup_allday_seasonal_account.cdc"
	AllDayAccountIsSetupPath         = AllDayScriptsRootPath + "/user/account_is_setup.cdc"
	AllDaySeasonalAccountIsSetupPath = AllDayScriptsRootPath + "/user/account_seasonal_is_setup.cdc"

	// Series
	AllDayCreateSeriesPath       = AllDayTransactionsRootPath + "/admin/series/create_series.cdc"
	AllDayCloseSeriesPath        = AllDayTransactionsRootPath + "/admin/series/close_series.cdc"
	AllDayReadAllSeriesPath      = AllDayScriptsRootPath + "/series/read_all_series.cdc"
	AllDayReadSeriesByIDPath     = AllDayScriptsRootPath + "/series/read_series_by_id.cdc"
	AllDayReadSeriesByNamePath   = AllDayScriptsRootPath + "/series/read_series_by_name.cdc"
	AllDayReadAllSeriesNamesPath = AllDayScriptsRootPath + "/series/read_all_series_names.cdc"

	// Sets
	AllDayCreateSetPath       = AllDayTransactionsRootPath + "/admin/sets/create_set.cdc"
	AllDayReadAllSetsPath     = AllDayScriptsRootPath + "/sets/read_all_sets.cdc"
	AllDayReadSetByIDPath     = AllDayScriptsRootPath + "/sets/read_set_by_id.cdc"
	AllDayReadSetsByNamePath  = AllDayScriptsRootPath + "/sets/read_sets_by_name.cdc"
	AllDayReadAllSetNamesPath = AllDayScriptsRootPath + "/sets/read_all_set_names.cdc"

	// Plays
	AllDayCreatePlayPath   = AllDayTransactionsRootPath + "/admin/plays/create_play.cdc"
	AllDayReadPlayByIDPath = AllDayScriptsRootPath + "/plays/read_play_by_id.cdc"
	AllDayReadAllPlaysPath = AllDayScriptsRootPath + "/plays/read_all_plays.cdc"

	// Editions
	AllDayCreateEditionPath         = AllDayTransactionsRootPath + "/admin/editions/create_edition.cdc"
	AllDaySeasonalCreateEditionPath = AllDayTransactionsRootPath + "/admin/editions/create_seasonal_edition.cdc"
	AllDayCloseEditionPath          = AllDayTransactionsRootPath + "/admin/editions/close_edition.cdc"
	AllDaySeasonalCloseEditionPath  = AllDayTransactionsRootPath + "/admin/editions/close_seasonal_edition.cdc"

	AllDayReadEditionByIDPath         = AllDayScriptsRootPath + "/editions/read_edition_by_id.cdc"
	AllDaySeasonalReadEditionByIDPath = AllDayScriptsRootPath + "/editions/read_seasonal_edition_by_id.cdc"
	AllDayReadAllEditionsPath         = AllDayScriptsRootPath + "/edition/read_all_editions.cdc"

	// Moment NFTs
	AllDayMintMomentNFTPath   = AllDayTransactionsRootPath + "/admin/nfts/mint_moment_nft.cdc"
	AllDayMintSeasonalNFTPath = AllDayTransactionsRootPath + "/admin/nfts/mint_seasonal_nft.cdc"

	AllDayMintMomentNFTMultiPath        = AllDayTransactionsRootPath + "/admin/nfts/mint_moment_nft_multi.cdc"
	AllDayTransferNFTPath               = AllDayTransactionsRootPath + "/user/transfer_moment_nft.cdc"
	AllDayReadMomentNFTSupplyPath       = AllDayScriptsRootPath + "/nfts/read_moment_nft_supply.cdc"
	AllDayReadSeasonalNFTSupplyPath     = AllDayScriptsRootPath + "/nfts/read_seasonal_nft_supply.cdc"
	AllDayReadMomentNFTPropertiesPath   = AllDayScriptsRootPath + "/nfts/read_moment_nft_properties.cdc"
	AllDayReadSeasonalNFTPropertiesPath = AllDayScriptsRootPath + "/nfts/read_seasonal_nft_properties.cdc"

	AllDayReadCollectionNFTLengthPath = AllDayScriptsRootPath + "/nfts/read_collection_nft_length.cdc"
	AllDayReadCollectionNFTIDsPath    = AllDayScriptsRootPath + "/nfts/read_collection_nft_ids.cdc"
)

//------------------------------------------------------------
// Accounts
//------------------------------------------------------------
func replaceAddresses(code []byte, contracts Contracts) []byte {
	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+contracts.NFTAddress.String()))

	AllDayRe := regexp.MustCompile(AllDayAddressPlaceholder)
	code = AllDayRe.ReplaceAll(code, []byte("0x"+contracts.AllDayAddress.String()))

	AllDaySeasonalRe := regexp.MustCompile(AllDaySeasonalAddressPlaceholder)
	code = AllDaySeasonalRe.ReplaceAll(code, []byte("0x"+contracts.AllDayAddress.String()))

	return code
}

func LoadAllDay(nftAddress flow.Address) []byte {
	code := readFile(AllDayPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	return code
}

func LoadAllDaySeasonal(nftAddress flow.Address) []byte {
	code := readFile(AllDaySeasonalPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	return code
}

func loadAllDaySetupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySetupAccountPath),
		contracts,
	)
}

func loadAllDaySeasonalSetupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySeasonalSetupAccountPath),
		contracts,
	)
}

func loadAllDayAccountIsSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayAccountIsSetupPath),
		contracts,
	)
}

func loadAllDaySeasonalAccountIsSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySeasonalAccountIsSetupPath),
		contracts,
	)
}

//------------------------------------------------------------
// Series
//------------------------------------------------------------
func loadAllDayCreateSeriesTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreateSeriesPath),
		contracts,
	)
}

func loadAllDayReadSeriesByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadSeriesByIDPath),
		contracts,
	)
}

func loadAllDayReadSeriesByNameScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadSeriesByNamePath),
		contracts,
	)
}

func loadAllDayReadAllSeriesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllSeriesPath),
		contracts,
	)
}

func loadAllDayReadAllSeriesNamesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllSeriesNamesPath),
		contracts,
	)
}

func loadAllDayCloseSeriesTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCloseSeriesPath),
		contracts,
	)
}

//------------------------------------------------------------
// Sets
//------------------------------------------------------------
func loadAllDayCreateSetTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreateSetPath),
		contracts,
	)
}

func loadAllDayReadSetByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadSetByIDPath),
		contracts,
	)
}

func loadAllDayReadAllSetsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllSetsPath),
		contracts,
	)
}

func loadAllDayReadSetsByNameScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadSetsByNamePath),
		contracts,
	)
}

func loadAllDayReadAllSetNamesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllSetNamesPath),
		contracts,
	)
}

//------------------------------------------------------------
// Plays
//------------------------------------------------------------
func loadAllDayCreatePlayTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreatePlayPath),
		contracts,
	)
}

func loadAllDayReadPlayByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadPlayByIDPath),
		contracts,
	)
}

func loadAllDayReadAllPlaysScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllPlaysPath),
		contracts,
	)
}

//------------------------------------------------------------
// Editions
//------------------------------------------------------------
func loadAllDayCreateEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreateEditionPath),
		contracts,
	)
}

func loadAllDaySeasonalCreateEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySeasonalCreateEditionPath),
		contracts,
	)
}

func loadAllDayReadEditionByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadEditionByIDPath),
		contracts,
	)
}

func loadAllDaySeasonalReadEditionByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySeasonalReadEditionByIDPath),
		contracts,
	)
}

func loadAllDayCloseEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCloseEditionPath),
		contracts,
	)
}

func loadAllDaySeasonalCloseEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySeasonalCloseEditionPath),
		contracts,
	)
}

func loadAllDayReadAllEditionsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllEditionsPath),
		contracts,
	)
}

//------------------------------------------------------------
// Moment NFTs
//------------------------------------------------------------
func loadAllDayMintMomentNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayMintMomentNFTPath),
		contracts,
	)
}

func loadAllDayMintSeasonalNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayMintMomentNFTPath),
		contracts,
	)
}

func loadAllDayMintMomentNFTMultiTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayMintMomentNFTMultiPath),
		contracts,
	)
}

func loadAllDayReadMomentNFTSupplyScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadMomentNFTSupplyPath),
		contracts,
	)
}

func loadAllDayReadMomentNFTPropertiesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadMomentNFTPropertiesPath),
		contracts,
	)
}

func loadAllDayReadSeasonalNFTPropertiesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadSeasonalNFTPropertiesPath),
		contracts,
	)
}

func loadAllDayReadCollectionNFTLengthScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadCollectionNFTLengthPath),
		contracts,
	)
}

func loadAllDayReadCollectionNFTIDsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadCollectionNFTIDsPath),
		contracts,
	)
}

func loadAllDayTransferNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayTransferNFTPath),
		contracts,
	)
}

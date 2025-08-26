package test

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

// Handle relative paths by making these regular expressions

const (
	nftAddressPlaceholder                      = "\"NonFungibleToken\""
	ftAddressPlaceholder                       = "\"FungibleToken\""
	mvAddressPlaceholder                       = "\"MetadataViews\""
	viewResolverAddressPlaceHolder             = "\"ViewResolver\""
	AllDayAddressPlaceholder                   = "\"AllDay\""
	FungibleTokenSwitchboardAddressPlaceholder = "\"FungibleTokenSwitchboard\""
	royaltyAddressPlaceholder                  = "0xALLDAYROYALTYADDRESS"
)

const (
	AllDayPath                 = "../../../contracts/AllDay.cdc"
	AllDayTransactionsRootPath = "../../../transactions"
	AllDayScriptsRootPath      = "../../../scripts"

	// Accounts
	AllDaySetupAccountPath     = AllDayTransactionsRootPath + "/user/setup_AllDay_account.cdc"
	AllDayAccountIsSetupPath   = AllDayScriptsRootPath + "/user/account_is_setup.cdc"
	AllDaySetupSwitchboardPath = AllDayTransactionsRootPath + "/user/setup_switchboard_account.cdc"

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
	AllDayCreatePlayPath                = AllDayTransactionsRootPath + "/admin/plays/create_play.cdc"
	AllDayUpdatePlayDescriptionPath     = AllDayTransactionsRootPath + "/admin/plays/update_play_description.cdc"
	AllDayUpdatePlayDynamicMetadataPath = AllDayTransactionsRootPath + "/admin/plays/update_play_dynamic_metadata.cdc"
	AllDayReadPlayByIDPath              = AllDayScriptsRootPath + "/plays/read_play_by_id.cdc"
	AllDayReadAllPlaysPath              = AllDayScriptsRootPath + "/plays/read_all_plays.cdc"

	// Editions
	AllDayCreateEditionPath   = AllDayTransactionsRootPath + "/admin/editions/create_edition.cdc"
	AllDayCloseEditionPath    = AllDayTransactionsRootPath + "/admin/editions/close_edition.cdc"
	AllDayReadEditionByIDPath = AllDayScriptsRootPath + "/editions/read_edition_by_id.cdc"
	AllDayReadAllEditionsPath = AllDayScriptsRootPath + "/edition/read_all_editions.cdc"

	// Moment NFTs
	AllDayMintMomentNFTPath           = AllDayTransactionsRootPath + "/admin/nfts/mint_moment_nft.cdc"
	AllDayMintMomentNFTMultiPath      = AllDayTransactionsRootPath + "/admin/nfts/mint_moment_nfts_multi.cdc"
	AllDayTransferNFTPath             = AllDayTransactionsRootPath + "/user/transfer_moment_nft.cdc"
	AllDayReadMomentNFTSupplyPath     = AllDayScriptsRootPath + "/nfts/read_moment_nft_supply.cdc"
	AllDayReadMomentNFTPropertiesPath = AllDayScriptsRootPath + "/nfts/read_moment_nft_properties.cdc"
	AllDayReadCollectionNFTLengthPath = AllDayScriptsRootPath + "/nfts/read_collection_nft_length.cdc"
	AllDayReadCollectionNFTIDsPath    = AllDayScriptsRootPath + "/nfts/read_collection_nft_ids.cdc"
	AllDayReadMomentNFTMetadataPath   = AllDayScriptsRootPath + "/nfts/read_moment_nft_metadata.cdc"

	// Badges
	AllDayCreateBadgePath            = AllDayTransactionsRootPath + "/admin/badges/create_badge.cdc"
	AllDayUpdateBadgePath            = AllDayTransactionsRootPath + "/admin/badges/update_badge.cdc"
	AllDayAddBadgeToPlayPath         = AllDayTransactionsRootPath + "/admin/badges/add_badge_to_play.cdc"
	AllDayAddBadgeToEditionPath      = AllDayTransactionsRootPath + "/admin/badges/add_badge_to_edition.cdc"
	AllDayAddBadgeToMomentPath       = AllDayTransactionsRootPath + "/admin/badges/add_badge_to_moment.cdc"
	AllDayRemoveBadgeFromPlayPath    = AllDayTransactionsRootPath + "/admin/badges/remove_badge_from_play.cdc"
	AllDayRemoveBadgeFromEditionPath = AllDayTransactionsRootPath + "/admin/badges/remove_badge_from_edition.cdc"
	AllDayRemoveBadgeFromMomentPath  = AllDayTransactionsRootPath + "/admin/badges/remove_badge_from_moment.cdc"
	AllDayGetBadgeBySlugPath         = AllDayScriptsRootPath + "/badges/get_badge_by_slug.cdc"
	AllDayGetNftAllBadgesPath        = AllDayScriptsRootPath + "/badges/get_nft_all_badges.cdc"
	AllDayBadgeExistsPath            = AllDayScriptsRootPath + "/badges/badge_exists.cdc"
)

// ------------------------------------------------------------
// Accounts
// ------------------------------------------------------------
func replaceAddresses(code []byte, contracts Contracts) []byte {
	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+contracts.NFTAddress.String()))

	ftRe := regexp.MustCompile(ftAddressPlaceholder)
	code = ftRe.ReplaceAll(code, []byte("0x"+ftAddress.String()))

	AllDayRe := regexp.MustCompile(AllDayAddressPlaceholder)
	code = AllDayRe.ReplaceAll(code, []byte("0x"+contracts.AllDayAddress.String()))

	mvRe := regexp.MustCompile(mvAddressPlaceholder)
	code = mvRe.ReplaceAll(code, []byte("0x"+contracts.MetadataViewsAddress.String()))

	switchboardRe := regexp.MustCompile(FungibleTokenSwitchboardAddressPlaceholder)
	code = switchboardRe.ReplaceAll(code, []byte("0x"+contracts.FungibleTokenSwitchboardAddress.String()))

	royaltyRe := regexp.MustCompile(royaltyAddressPlaceholder)
	code = royaltyRe.ReplaceAll(code, []byte("0x"+contracts.RoyaltyAddress.String()))

	return code
}

func LoadAllDay(nftAddress flow.Address, metaAddress flow.Address, royaltyAddress flow.Address, viewResolverAddress flow.Address) []byte {
	code := readFile(AllDayPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	ftRe := regexp.MustCompile(ftAddressPlaceholder)
	code = ftRe.ReplaceAll(code, []byte("0x"+ftAddress.String()))

	mvRe := regexp.MustCompile(mvAddressPlaceholder)
	code = mvRe.ReplaceAll(code, []byte("0x"+metaAddress.String()))

	viewResolverRe := regexp.MustCompile(viewResolverAddressPlaceHolder)
	code = viewResolverRe.ReplaceAll(code, []byte("0x"+viewResolverAddress.String()))

	royaltyRe := regexp.MustCompile(royaltyAddressPlaceholder)
	code = royaltyRe.ReplaceAll(code, []byte("0x"+royaltyAddress.String()))

	//fmt.Println(string(code))
	return code
}

func loadAllDaySetupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySetupAccountPath),
		contracts,
	)
}

func loadAllDayAccountIsSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayAccountIsSetupPath),
		contracts,
	)
}

func loadSetupSwitchboardAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDaySetupSwitchboardPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Series
// ------------------------------------------------------------
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

// ------------------------------------------------------------
// Sets
// ------------------------------------------------------------
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

// ------------------------------------------------------------
// Plays
// ------------------------------------------------------------
func loadAllDayCreatePlayTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreatePlayPath),
		contracts,
	)
}

func loadAllDayUpdatePlayDescriptionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayUpdatePlayDescriptionPath),
		contracts,
	)
}

func loadAllDayUpdateDayUpdatePlayDynamicMetadataTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayUpdatePlayDynamicMetadataPath),
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

// ------------------------------------------------------------
// Editions
// ------------------------------------------------------------
func loadAllDayCreateEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreateEditionPath),
		contracts,
	)
}

func loadAllDayReadEditionByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadEditionByIDPath),
		contracts,
	)
}

func loadAllDayCloseEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCloseEditionPath),
		contracts,
	)
}

func loadAllDayReadAllEditionsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadAllEditionsPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Moment NFTs
// ------------------------------------------------------------
func loadAllDayMintMomentNFTTransaction(contracts Contracts) []byte {
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

func loadAllDayReadMomentNFTMetadataScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayReadMomentNFTMetadataPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Badges
// ------------------------------------------------------------
func loadAllDayCreateBadgeTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayCreateBadgePath),
		contracts,
	)
}

func loadAllDayUpdateBadgeTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayUpdateBadgePath),
		contracts,
	)
}

func loadAllDayAddBadgeToPlayTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayAddBadgeToPlayPath),
		contracts,
	)
}

func loadAllDayAddBadgeToEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayAddBadgeToEditionPath),
		contracts,
	)
}

func loadAllDayAddBadgeToMomentTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayAddBadgeToMomentPath),
		contracts,
	)
}

func loadAllDayRemoveBadgeFromPlayTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayRemoveBadgeFromPlayPath),
		contracts,
	)
}

func loadAllDayRemoveBadgeFromEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayRemoveBadgeFromEditionPath),
		contracts,
	)
}

func loadAllDayRemoveBadgeFromMomentTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayRemoveBadgeFromMomentPath),
		contracts,
	)
}

func loadAllDayGetBadgeBySlugScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayGetBadgeBySlugPath),
		contracts,
	)
}

func loadAllDayGetNftAllBadgesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayGetNftAllBadgesPath),
		contracts,
	)
}

func loadAllDayBadgeExistsScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(AllDayBadgeExistsPath),
		contracts,
	)
}

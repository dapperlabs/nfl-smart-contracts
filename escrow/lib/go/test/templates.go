package test

import (
	"regexp"

	"github.com/onflow/flow-go-sdk"
)

// Handle relative paths by making these regular expressions

const (
	nftAddressPlaceholder     = "\"[^\"]*NonFungibleToken.cdc\""
	ftAddressPlaceholder      = "\"[^\"]*FungibleToken.cdc\""
	mvAddressPlaceholder      = "\"[^\"]*MetadataViews.cdc\""
	AllDayAddressPlaceholder  = "\"[^\"]*AllDay.cdc\""
	royaltyAddressPlaceholder = "0xALLDAYROYALTYADDRESS"
	escrowAddressPlaceholder  = "\"[^\"]*Escrow.cdc\""
)

const (
	EscrowPath                 = "../../../contracts/Escrow.cdc"
	AllDayPath                 = "../../../contracts/AllDay.cdc"
	EscrowTransactionsRootPath = "../../../transactions"
	EscrowScriptsRootPath      = "../../../scripts"

	// Accounts
	EscrowSetupAccountPath   = EscrowTransactionsRootPath + "/user/setup_AllDay_account.cdc"
	EscrowAccountIsSetupPath = EscrowScriptsRootPath + "/user/account_is_setup.cdc"

	// Series
	EscrowCreateSeriesPath   = EscrowTransactionsRootPath + "/admin/series/create_series.cdc"
	EscrowCloseSeriesPath    = EscrowTransactionsRootPath + "/admin/series/close_series.cdc"
	EscrowReadSeriesByIDPath = EscrowScriptsRootPath + "/series/read_series_by_id.cdc"

	// Sets
	EscrowCreateSetPath   = EscrowTransactionsRootPath + "/admin/sets/create_set.cdc"
	EscrowReadSetByIDPath = EscrowScriptsRootPath + "/sets/read_set_by_id.cdc"

	// Plays
	EscrowCreatePlayPath                = EscrowTransactionsRootPath + "/admin/plays/create_play.cdc"
	EscrowUpdatePlayDescriptionPath     = EscrowTransactionsRootPath + "/admin/plays/update_play_description.cdc"
	EscrowUpdatePlayDynamicMetadataPath = EscrowTransactionsRootPath + "/admin/plays/update_play_dynamic_metadata.cdc"
	EscrowReadPlayByIDPath              = EscrowScriptsRootPath + "/plays/read_play_by_id.cdc"

	// Editions
	EscrowCreateEditionPath   = EscrowTransactionsRootPath + "/admin/editions/create_edition.cdc"
	EscrowReadEditionByIDPath = EscrowScriptsRootPath + "/editions/read_edition_by_id.cdc"

	// Leaderboards
	EscrowCreateLeaderboardPath = EscrowTransactionsRootPath + "/admin/leaderboards/create_leaderboard.cdc"
	EscrowGetLeaderboardPath    = EscrowTransactionsRootPath + "/admin/leaderboards/get_leaderboard.cdc"

	// Moment NFTs
	EscrowMintMomentNFTPath           = EscrowTransactionsRootPath + "/admin/nfts/mint_moment_nft.cdc"
	EscrowReadMomentNFTSupplyPath     = EscrowScriptsRootPath + "/nfts/read_moment_nft_supply.cdc"
	EscrowReadMomentNFTPropertiesPath = EscrowScriptsRootPath + "/nfts/read_moment_nft_properties.cdc"
	EscrowReadCollectionLengthPath    = EscrowScriptsRootPath + "/nfts/read_collection_nft_length.cdc"

	// Escrow
	EscrowMomentNFTPath                  = EscrowTransactionsRootPath + "/user/add_entry.cdc"
	EscrowWithdrawMomentNFTPath          = EscrowTransactionsRootPath + "/admin/leaderboards/withdraw_entry.cdc"
	EscrowBurnNFTPath                    = EscrowTransactionsRootPath + "/admin/leaderboards/burn_nft.cdc"
	EscrowReadNFTLengthInLeaderboardPath = EscrowScriptsRootPath + "/leaderboards/read_entries_length.cdc"
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

	royaltyRe := regexp.MustCompile(royaltyAddressPlaceholder)
	code = royaltyRe.ReplaceAll(code, []byte("0x"+contracts.RoyaltyAddress.String()))

	escrowRe := regexp.MustCompile(escrowAddressPlaceholder)
	code = escrowRe.ReplaceAll(code, []byte("0x"+contracts.EscrowAddress.String()))

	return code
}

func LoadAllDay(nftAddress flow.Address, metaAddress flow.Address, royaltyAddress flow.Address) []byte {
	code := readFile(AllDayPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	ftRe := regexp.MustCompile(ftAddressPlaceholder)
	code = ftRe.ReplaceAll(code, []byte("0x"+ftAddress.String()))

	mvRe := regexp.MustCompile(mvAddressPlaceholder)
	code = mvRe.ReplaceAll(code, []byte("0x"+metaAddress.String()))

	royaltyRe := regexp.MustCompile(royaltyAddressPlaceholder)
	code = royaltyRe.ReplaceAll(code, []byte("0x"+royaltyAddress.String()))

	return code
}

func LoadEscrow(nftAddress flow.Address) []byte {
	code := readFile(EscrowPath)

	nftRe := regexp.MustCompile(nftAddressPlaceholder)
	code = nftRe.ReplaceAll(code, []byte("0x"+nftAddress.String()))

	return code
}

func loadEscrowSetupAccountTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowSetupAccountPath),
		contracts,
	)
}

func loadEscrowAccountIsSetupScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowAccountIsSetupPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Series
// ------------------------------------------------------------
func loadEscrowCreateSeriesTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowCreateSeriesPath),
		contracts,
	)
}

func loadEscrowReadSeriesByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadSeriesByIDPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Sets
// ------------------------------------------------------------
func loadEscrowCreateSetTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowCreateSetPath),
		contracts,
	)
}

func loadEscrowReadSetByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadSetByIDPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Plays
// ------------------------------------------------------------
func loadEscrowCreatePlayTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowCreatePlayPath),
		contracts,
	)
}

func loadEscrowReadPlayByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadPlayByIDPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Editions
// ------------------------------------------------------------
func loadEscrowCreateEditionTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowCreateEditionPath),
		contracts,
	)
}

func loadEscrowReadEditionByIDScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadEditionByIDPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Moment NFTs
// ------------------------------------------------------------
func loadEscrowMintMomentNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowMintMomentNFTPath),
		contracts,
	)
}

func loadEscrowReadCollectionLengthScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadCollectionLengthPath),
		contracts,
	)
}

func loadEscrowReadMomentNFTSupplyScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadMomentNFTSupplyPath),
		contracts,
	)
}

func loadEscrowReadMomentNFTPropertiesScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadMomentNFTPropertiesPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Escrow
// ------------------------------------------------------------
func loadEscrowMomentNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowMomentNFTPath),
		contracts,
	)
}

func loadEscrowReadNFTLengthInLeaderboardScript(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowReadNFTLengthInLeaderboardPath),
		contracts,
	)
}

func loadEscrowWithdrawMomentNFT(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowWithdrawMomentNFTPath),
		contracts,
	)
}

func loadEscrowBurnNFTTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowBurnNFTPath),
		contracts,
	)
}

// ------------------------------------------------------------
// Leaderboards
// ------------------------------------------------------------
func loadEscrowCreateLeaderboardTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowCreateLeaderboardPath),
		contracts,
	)
}

func loadEscrowGetLeaderboardTransaction(contracts Contracts) []byte {
	return replaceAddresses(
		readFile(EscrowGetLeaderboardPath),
		contracts,
	)
}

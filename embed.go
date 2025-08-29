package nfl

import (
	_ "embed"
)

// scripts is a list of all the scripts we export with imports mapped
var (
	//go:embed scripts/user/account_is_all_setup.cdc
	UserAccountIsAllSetup []byte
	//go:embed scripts/user/account_is_setup.cdc
	UserAccountIsSetup []byte
)

// Transactions is a list of all the transactions we export with imports mapped
var (
	//go:embed transactions/admin/editions/close_edition.cdc
	EditionsCloseEdition []byte
	//go:embed transactions/admin/editions/create_edition.cdc
	EditionsCreateEdition []byte
	//go:embed transactions/admin/nfts/mint_moment_nft.cdc
	NftsMintMomentNft []byte
	//go:embed transactions/admin/nfts/mint_moment_nfts_multi.cdc
	NftsBatchMintMomentNfts []byte
	//go:embed transactions/admin/plays/create_play.cdc
	PlaysCreatePlay []byte
	//go:embed transactions/admin/plays/update_play_description.cdc
	PlaysUpdatePlayDescription []byte
	//go:embed transactions/admin/plays/update_play_dynamic_metadata.cdc
	PlaysUpdatePlayDynamicMetadata []byte
	//go:embed transactions/admin/series/close_series.cdc
	SeriesCloseSeries []byte
	//go:embed transactions/admin/series/create_series.cdc
	SeriesCreateSeries []byte
	//go:embed transactions/admin/sets/create_set.cdc
	SetsCreateSet []byte
	//go:embed transactions/admin/badges/create_badge.cdc
	CreateBadge []byte
	//go:embed transactions/admin/badges/update_badge.cdc
	UpdateBadge []byte
	//go:embed transactions/admin/badges/delete_badge.cdc
	DeleteBadge []byte
	//go:embed transactions/admin/badges/add_badge_to_entity.cdc
	AddBadgeToEntity []byte
	//go:embed transactions/admin/badges/remove_badge_from_entity.cdc
	RemoveBadgeFromEntity []byte

	//go:embed transactions/user/setup_allday_account.cdc
	UserSetupAllDayAccount []byte
	//go:embed transactions/user/transfer_moment_nft.cdc
	UserTransferMomentNft []byte
	//go:embed transactions/user/batch_transfer_moment_nfts.cdc
	UserBatchTransferMomentNfts []byte
	//go:embed transactions/user/setup_all_collections.cdc
	UserSetUpAllCollections []byte
)

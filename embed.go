package nfl

import (
	"context"
	"embed"
	"fmt"
	"github.com/onflow/flow-go-sdk"
	"os"
	"reflect"

	"github.com/onflow/cadence"
	"github.com/onflow/flowkit/v2/config"
	"github.com/onflow/flowkit/v2/config/json"
	"github.com/onflow/flowkit/v2/project"
)

//go:embed transactions/*
//go:embed scripts/*
var transactionScripts embed.FS

//go:embed flow.json
var configFS embed.FS

var ResolvedScripts Scripts
var ResolvedTransactions Transactions

// Config can be created to have Transactions that use arbitrary addresses
type Config struct {
	AllDayAddress                string `default:"0x4dfd62c88d1b6462"`
	PackNFTAddress               string `default:"0x4dfd62c88d1b6462"`
	NFTAddress                   string `default:"0x631e88ae7f1d7c20"`
	FungibleTokenContractAddress string `default:"0x9a0766d93b6608b7"`
	//DapperUtilityCoinContractAddress string `default:"0x82ec283f88a62e65"`
	MetadataViewsAddress    string `default:"0x631e88ae7f1d7c20"`
	ViewResolverAddress     string `default:"0x631e88ae7f1d7c20"`
	BurnerAddress           string `default:"0x9a0766d93b6608b7"`
	IPackNFTContractAddress string `default:"0xd8f6346999b983f5"`
}

// Scripts is a list of all the scripts we export with imports mapped
// Good candidate for autogeneration in the future
type Scripts struct {
	UserAccountIsAllSetup []byte `src:"scripts/user/account_is_all_setup.cdc"`
	UserAccountIsSetup    []byte `src:"scripts/user/account_is_setup.cdc"`
}

// Transactions is a list of all the transactions we export with imports mapped
type Transactions struct {
	EditionsCloseEdition           []byte `src:"transactions/admin/editions/close_edition.cdc"`
	EditionsCreateEdition          []byte `src:"transactions/admin/editions/create_edition.cdc"`
	NftsMintMomentNft              []byte `src:"transactions/admin/nfts/mint_moment_nft.cdc"`
	NftsBatchMintMomentNfts        []byte `src:"transactions/admin/nfts/mint_moment_nfts_multi.cdc"`
	PlaysCreatePlay                []byte `src:"transactions/admin/plays/create_play.cdc"`
	PlaysUpdatePlayDescription     []byte `src:"transactions/admin/plays/update_play_description.cdc"`
	PlaysUpdatePlayDynamicMetadata []byte `src:"transactions/admin/plays/update_play_dynamic_metadata.cdc"`
	SeriesCloseSeries              []byte `src:"transactions/admin/series/close_series.cdc"`
	SeriesCreateSeries             []byte `src:"transactions/admin/series/create_series.cdc"`
	SetsCreateSet                  []byte `src:"transactions/admin/sets/create_set.cdc"`

	UserSetupAllDayAccount  []byte `src:"transactions/user/setup_allday_account.cdc"`
	UserTransferMomentNft   []byte `src:"transactions/user/transfer_moment_nft.cdc"`
	UserSetUpAllCollections []byte `src:"transactions/user/setup_all_collections.cdc"`
}

func init() {
	networkSpecificConfig, err := ConfigForEnv(getNetwork())
	if err != nil {
		panic(fmt.Errorf("unable to resolve config for network: %w", err))
	}
	ResolvedTransactions, err = NewTransactions(context.Background(), networkSpecificConfig)
	if err != nil {
		panic(fmt.Errorf("unable to build transactions: %w", err))
	}

	ResolvedScripts, err = NewScripts(context.Background(), networkSpecificConfig)
	if err != nil {
		panic(fmt.Errorf("unable to build scripts: %w", err))
	}
}

// NewTransactions returns a struct with all the transactions
func NewTransactions(ctx context.Context, conf Config) (Transactions, error) {
	// resolve addresses and expose templates
	txs := Transactions{}
	scriptProcessor := NewScriptProcessor(conf)

	// to avoid a lot of boiler plate we use struct tags to map the file names
	// this isn't strictly necessary, and we could just perpetually add to a list here
	// but this would result in more "shot gun surgery" every time a new tx is added
	// tradeoff is legibility vs ease of maintenance
	// another alternative that would likely be superior is to autogenerate the code
	// by crawling the file tree
	txType := reflect.TypeOf(txs)
	txVal := reflect.ValueOf(&txs)
	for i := 0; i < txType.NumField(); i++ {
		f := txType.Field(i)
		path := f.Tag.Get("src")
		script, err := scriptProcessor.Process(ctx, conf, path)
		if err != nil {
			return Transactions{}, err
		}
		txVal.Elem().Field(i).SetBytes(script)
	}

	return txs, nil
}

// NewScripts returns a struct with all the transactions
func NewScripts(ctx context.Context, conf Config) (Scripts, error) {
	// resolve addresses and expose templates
	scripts := Scripts{}
	scriptProcessor := NewScriptProcessor(conf)

	// to avoid a lot of boiler plate we use struct tags to map the file names
	// this isn't strictly necessary, and we could just perpetually add to a list here
	// but this would result in more "shot gun surgery" every time a new tx is added
	// tradeoff is legibility vs ease of maintenance
	// another alternative that would likely be superior is to autogenerate the code
	// by crawling the file tree
	scriptType := reflect.TypeOf(scripts)
	scriptVal := reflect.ValueOf(&scripts)
	for i := 0; i < scriptType.NumField(); i++ {
		f := scriptType.Field(i)
		path := f.Tag.Get("src")
		script, err := scriptProcessor.Process(ctx, conf, path)
		if err != nil {
			return Scripts{}, err
		}
		scriptVal.Elem().Field(i).SetBytes(script)
	}

	return scripts, nil
}

type ScriptProcessor struct {
	*project.ImportReplacer
}

// ConfigFromEnv will return config appropriate for the network
func ConfigForEnv(network string) (Config, error) {
	loader := config.NewLoader(&embedReaderWriter{configFS})
	loader.AddConfigParser(json.NewParser())
	flowConf, err := loader.Load([]string{"flow.json"})
	if err != nil {
		return Config{}, err
	}

	allDayAddress, err := contractAddressHex(flowConf, "AllDay", network)
	if err != nil {
		return Config{}, err
	}
	nftAddress, err := contractAddressHex(flowConf, "NonFungibleToken", network)
	if err != nil {
		return Config{}, err
	}
	ftAddress, err := contractAddressHex(flowConf, "FungibleToken", network)
	if err != nil {
		return Config{}, err
	}
	//dapperUtilityCoinAddress, err := contractAddressHex(flowConf, "DapperUtilityCoin", network)
	//if err != nil {
	//	return Config{}, err
	//}
	metadataViewsAddress, err := contractAddressHex(flowConf, "MetadataViews", network)
	if err != nil {
		return Config{}, err
	}
	viewResolverAddress, err := contractAddressHex(flowConf, "ViewResolver", network)
	if err != nil {
		return Config{}, err
	}
	burnerAddress, err := contractAddressHex(flowConf, "Burner", network)
	if err != nil {
		return Config{}, err
	}
	iPackNftContractAddress, err := contractAddressHex(flowConf, "IPackNFT", network)
	if err != nil {
		return Config{}, err
	}
	packNftContractAddress, err := contractAddressHex(flowConf, "PackNFT", network)
	if err != nil {
		return Config{}, err
	}

	return Config{
		// if we have other addresses to be pulled from config, add them here
		AllDayAddress:                allDayAddress,
		NFTAddress:                   nftAddress,
		FungibleTokenContractAddress: ftAddress,
		//DapperUtilityCoinContractAddress: dapperUtilityCoinAddress,
		MetadataViewsAddress:    metadataViewsAddress,
		ViewResolverAddress:     viewResolverAddress,
		BurnerAddress:           burnerAddress,
		IPackNFTContractAddress: iPackNftContractAddress,
		PackNFTAddress:          packNftContractAddress,
	}, nil
}

func contractAddressHex(flowConf *config.Config, name, network string) (string, error) {
	s, err := flowConf.Contracts.ByName(name)
	if err != nil {
		return "", err
	}
	return s.Aliases.ByNetwork(network).Address.String(), nil
}

func NewScriptProcessor(conf Config) *ScriptProcessor {
	return &ScriptProcessor{newImportReplacer(conf)}
}

func newImportReplacer(conf Config) *project.ImportReplacer {
	// map addresses used as imports
	allDayContract := project.NewContract("AllDay", "./contracts/AllDay.cdc", nil,
		flow.HexToAddress(conf.AllDayAddress), "AllDayDeployer", nil)
	nftContract := project.NewContract("NonFungibleToken", "./contracts/imports/NonFungibleToken.cdc", nil,
		flow.HexToAddress(conf.NFTAddress), "NonFungibleTokenDeployer", nil)
	metadataViewsContract := project.NewContract("MetadataViews", "./contracts/imports/MetadataViews.cdc", nil,
		flow.HexToAddress(conf.MetadataViewsAddress), "MetadataViewsDeployer", nil)
	viewResolverContract := project.NewContract("ViewResolver", "./contracts/imports/ViewResolver.cdc", nil,
		flow.HexToAddress(conf.ViewResolverAddress), "ViewResolverDeployer", nil)
	burnerContract := project.NewContract("Burner", "./contracts/imports/Burner.cdc", nil,
		flow.HexToAddress(conf.BurnerAddress), "BurnerDeployer", nil)
	ftContract := project.NewContract("FungibleToken", "./contracts/imports/FungibleToken.cdc", nil,
		flow.HexToAddress(conf.FungibleTokenContractAddress), "FungibleTokenDeployer", nil)
	//ducContract := project.NewContract("DapperUtilityCoin", "./contracts/imports/DapperUtilityCoin.cdc", nil,
	//	flow.HexToAddress(conf.NFTAddress), "DapperUtilityCoinDeployer", nil)
	iPackNftContract := project.NewContract("IPackNFT", "./contracts/imports/IPackNFT.cdc", nil,
		flow.HexToAddress(conf.IPackNFTContractAddress), "IPackNFTDeployer", nil)
	packNftContract := project.NewContract("PackNFT", "./contracts/PackNFT.cdc", nil,
		flow.HexToAddress(conf.PackNFTAddress), "PackNFTDeployer", nil)

	return project.NewImportReplacer(
		[]*project.Contract{
			allDayContract,
			nftContract,
			metadataViewsContract,
			ftContract,
			//ducContract,
			iPackNftContract,
			packNftContract,
			viewResolverContract,
			burnerContract,
		},
		map[string]string{
			"AllDay":           conf.AllDayAddress,
			"NonFungibleToken": conf.NFTAddress,
			"MetadataViews":    conf.MetadataViewsAddress,
			//"DapperUtilityCoin": conf.DapperUtilityCoinContractAddress,
			"FungibleToken": conf.FungibleTokenContractAddress,
			"IPackNFT":      conf.IPackNFTContractAddress,
			"PackNFT":       conf.PackNFTAddress,
			"ViewResolver":  conf.ViewResolverAddress,
			"Burner":        conf.BurnerAddress,
		})
}

func (s *ScriptProcessor) Process(ctx context.Context, conf Config, path string) ([]byte, error) {
	raw, err := transactionScripts.ReadFile(path)
	if err != nil {
		return nil, err
	}
	program, err := project.NewProgram(raw, []cadence.Value{}, path)
	if err != nil {
		return nil, err
	}

	replaced, err := s.Replace(program)
	if err != nil {
		return nil, err
	}
	return replaced.Code(), nil
}

type embedReaderWriter struct {
	embed.FS
}

func (e *embedReaderWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	// no op
	return nil
}

func getNetwork() string {
	network, has := os.LookupEnv("FLOW_NETWORK")
	if !has {
		network = "testnet"
	}
	return network
}

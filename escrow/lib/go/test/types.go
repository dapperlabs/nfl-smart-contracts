package test

import (
	"fmt"
	"github.com/onflow/cadence"
)

type SeriesData struct {
	ID     uint64
	Name   string
	Active bool
}
type SetData struct {
	ID   uint64
	Name string
}
type PlayData struct {
	ID             uint64
	Classification string
	Metadata       map[string]string
}
type EditionData struct {
	ID          uint64
	SeriesID    uint64
	SetID       uint64
	PlayID      uint64
	MaxMintSize *uint64
	Tier        string
}
type OurNFTData struct {
	ID           uint64
	EditionID    uint64
	SerialNumber uint64
	// A UFix64 in uint64 form
	MintingDate uint64
}
type LeaderboardInfo struct {
	Name          string
	NftType       string
	EntriesLength uint64
}

func cadenceStringDictToGo(cadenceDict cadence.Dictionary) map[string]string {
	goDict := make(map[string]string)
	for _, pair := range cadenceDict.Pairs {
		goDict[pair.Key.ToGoValue().(string)] = pair.Value.ToGoValue().(string)
	}
	return goDict
}

func parseSeriesData(value cadence.Value) SeriesData {
	fields := value.(cadence.Struct).Fields
	return SeriesData{
		fields[0].ToGoValue().(uint64),
		fields[1].ToGoValue().(string),
		fields[2].ToGoValue().(bool),
	}
}

func parseSetData(value cadence.Value) SetData {
	fields := value.(cadence.Struct).Fields
	return SetData{
		fields[0].ToGoValue().(uint64),
		fields[1].ToGoValue().(string),
	}
}

func parsePlayData(value cadence.Value) PlayData {
	fields := value.(cadence.Struct).Fields
	return PlayData{
		fields[0].ToGoValue().(uint64),
		fields[1].ToGoValue().(string),
		cadenceStringDictToGo(fields[2].(cadence.Dictionary)),
	}
}

func parseEditionData(value cadence.Value) EditionData {
	fields := value.(cadence.Struct).Fields
	var maxMintSize uint64
	if fields[4] != nil && fields[4].ToGoValue() != nil {
		maxMintSize = fields[4].ToGoValue().(uint64)
	}
	return EditionData{
		fields[0].ToGoValue().(uint64),
		fields[1].ToGoValue().(uint64),
		fields[2].ToGoValue().(uint64),
		fields[3].ToGoValue().(uint64),
		&maxMintSize,
		fields[5].ToGoValue().(string),
	}
}

func parseNFTProperties(value cadence.Value) OurNFTData {
	array := value.(cadence.Array).Values
	return OurNFTData{
		array[0].ToGoValue().(uint64),
		array[1].ToGoValue().(uint64),
		array[2].ToGoValue().(uint64),
		array[3].ToGoValue().(uint64),
	}
}

func parseLeaderboardInfo(value cadence.Value) (LeaderboardInfo, error) {
	optionalVal, ok := value.(cadence.Optional)
	if !ok {
		return LeaderboardInfo{}, fmt.Errorf("expected value to be of type cadence.Optional, got %T", value)
	}

	if optionalVal.Value == nil {
		return LeaderboardInfo{}, fmt.Errorf("optional value is nil")
	}

	structVal, ok := optionalVal.Value.(cadence.Struct)
	if !ok {
		return LeaderboardInfo{}, fmt.Errorf("inner value of the optional is not a Struct, got %T", optionalVal.Value)
	}

	fields := structVal.Fields
	if len(fields) < 3 {
		return LeaderboardInfo{}, fmt.Errorf("struct does not contain enough fields")
	}

	name, ok := fields[0].(cadence.String)
	if !ok {
		return LeaderboardInfo{}, fmt.Errorf("field 0 is not a String")
	}

	nftType := fields[1]

	entriesLength := fields[2].(cadence.Int).Value.Uint64()

	return LeaderboardInfo{
		Name:          name.String(),
		NftType:       nftType.String(),
		EntriesLength: entriesLength,
	}, nil
}

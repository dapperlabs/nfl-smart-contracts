package test

import (
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

func cadenceStringDictToGo(cadenceDict cadence.Dictionary) map[string]string {
	goDict := make(map[string]string)
	for _, pair := range cadenceDict.Pairs {
		goDict[string(pair.Key.(cadence.String))] = string(pair.Value.(cadence.String))
	}
	return goDict
}

func parseSeriesData(value cadence.Value) SeriesData {
	fields := value.(cadence.Struct).FieldsMappedByName()
	return SeriesData{
		uint64(fields["id"].(cadence.UInt64)),
		string(fields["name"].(cadence.String)),
		bool(fields["active"].(cadence.Bool)),
	}
}

func parseSetData(value cadence.Value) SetData {
	fields := value.(cadence.Struct).FieldsMappedByName()
	return SetData{
		uint64(fields["id"].(cadence.UInt64)),
		string(fields["name"].(cadence.String)),
	}
}

func parsePlayData(value cadence.Value) PlayData {
	fields := value.(cadence.Struct).FieldsMappedByName()
	return PlayData{
		uint64(fields["id"].(cadence.UInt64)),
		string(fields["classification"].(cadence.String)),
		cadenceStringDictToGo(fields["metadata"].(cadence.Dictionary)),
	}
}

func parseEditionData(value cadence.Value) EditionData {
	fields := value.(cadence.Struct).FieldsMappedByName()
	var maxMintSize uint64
	if fields["maxMintSize"].(cadence.Optional).Value != nil {
		maxMintSize = uint64(fields["maxMintSize"].(cadence.Optional).Value.(cadence.UInt64))
	}
	return EditionData{
		uint64(fields["id"].(cadence.UInt64)),
		uint64(fields["seriesID"].(cadence.UInt64)),
		uint64(fields["setID"].(cadence.UInt64)),
		uint64(fields["playID"].(cadence.UInt64)),
		&maxMintSize,
		string(fields["tier"].(cadence.String)),
	}
}

func parseNFTProperties(value cadence.Value) OurNFTData {
	array := value.(cadence.Array).Values
	return OurNFTData{
		uint64(array[0].(cadence.UInt64)),
		uint64(array[1].(cadence.UInt64)),
		uint64(array[2].(cadence.UInt64)),
		uint64(array[3].(cadence.UFix64)),
	}
}

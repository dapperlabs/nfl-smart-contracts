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
	ID           uint64
	SeriesID     uint64
	SetID        uint64
	PlayID       uint64
	MaxMintSize  *uint64
	Tier         string
	Deserialized *bool
	PresetSerial *uint64
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

	var deserialized bool
	if fields[7] != nil && fields[7].ToGoValue() != nil {
		deserialized = fields[7].ToGoValue().(bool)
	}

	var presetSerial uint64
	if fields[8] != nil && fields[8].ToGoValue() != nil {
		presetSerial = fields[8].ToGoValue().(uint64)
	}
	return EditionData{
		fields[0].ToGoValue().(uint64),
		fields[1].ToGoValue().(uint64),
		fields[2].ToGoValue().(uint64),
		fields[3].ToGoValue().(uint64),
		&maxMintSize,
		fields[5].ToGoValue().(string),
		&deserialized,
		&presetSerial,
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

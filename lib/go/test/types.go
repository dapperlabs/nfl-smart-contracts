package test

import (
	"github.com/onflow/cadence"
)

type SeriesData struct {
	ID              uint32
	Name            string
	Metadata        map[string]string
	Active          bool
	CollectionIDs   []uint32
	CollectionsOpen uint32
}

type GeniesCollectionData struct {
	ID             uint32
	SeriesID       uint32
	Name           string
	Metadata       map[string]string
	Open           bool
	EditionIDs     []uint32
	EditionsActive uint32
}

type EditionData struct {
	ID           uint32
	CollectionID uint32
	Name         string
	Metadata     map[string]string
	Open         bool
	NumMinted    uint32
}

type OurNFTData struct {
	ID           uint64
	EditionID    uint32
	SerialNumber uint32
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

func cadenceUInt32ArrToGo(cadenceArr cadence.Array) []uint32 {
	goArr := make([]uint32, len(cadenceArr.Values))
	for i, val := range cadenceArr.Values {
		goArr[i] = val.ToGoValue().(uint32)
	}
	return goArr
}

func parseSeriesData(value cadence.Value) SeriesData {
	fields := value.(cadence.Struct).Fields
	return SeriesData{
		fields[0].ToGoValue().(uint32),
		fields[1].ToGoValue().(string),
		cadenceStringDictToGo(fields[2].(cadence.Dictionary)),
		fields[3].ToGoValue().(bool),
		cadenceUInt32ArrToGo(fields[4].(cadence.Array)),
		fields[5].ToGoValue().(uint32),
	}
}

func parseGeniesCollectionData(value cadence.Value) GeniesCollectionData {
	fields := value.(cadence.Struct).Fields
	return GeniesCollectionData{
		fields[0].ToGoValue().(uint32),
		fields[1].ToGoValue().(uint32),
		fields[2].ToGoValue().(string),
		cadenceStringDictToGo(fields[3].(cadence.Dictionary)),
		fields[4].ToGoValue().(bool),
		cadenceUInt32ArrToGo(fields[5].(cadence.Array)),
		fields[6].ToGoValue().(uint32),
	}
}

func parseEditionData(value cadence.Value) EditionData {
	fields := value.(cadence.Struct).Fields
	return EditionData{
		fields[0].ToGoValue().(uint32),
		fields[1].ToGoValue().(uint32),
		fields[2].ToGoValue().(string),
		cadenceStringDictToGo(fields[3].(cadence.Dictionary)),
		fields[4].ToGoValue().(bool),
		fields[5].ToGoValue().(uint32),
	}
}

func parseNFTProperties(value cadence.Value) OurNFTData {
	array := value.(cadence.Array).Values
	return OurNFTData{
		array[0].ToGoValue().(uint64),
		array[1].ToGoValue().(uint32),
		array[2].ToGoValue().(uint32),
		array[3].ToGoValue().(uint64),
	}
}

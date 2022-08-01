package test

import (
	"github.com/onflow/cadence"
)

type SeasonalEditionData struct {
	ID       uint64
	Metadata map[string]string
	Active   bool
}

type SeasonalNFTData struct {
	ID        uint64
	EditionID uint64
}

//func cadenceStringDictToGo(cadenceDict cadence.Dictionary) map[string]string {
//	goDict := make(map[string]string)
//	for _, pair := range cadenceDict.Pairs {
//		goDict[pair.Key.ToGoValue().(string)] = pair.Value.ToGoValue().(string)
//	}
//	return goDict
//}

func parseSeasonalEditionData(value cadence.Value) SeasonalEditionData {
	fields := value.(cadence.Struct).Fields
	return SeasonalEditionData{
		fields[0].ToGoValue().(uint64),
		cadenceStringDictToGo(fields[2].(cadence.Dictionary)),
		true,
	}
}

func parseSeasonalNFTProperties(value cadence.Value) SeasonalNFTData {
	array := value.(cadence.Array).Values
	return SeasonalNFTData{
		array[0].ToGoValue().(uint64),
		array[1].ToGoValue().(uint64),
	}
}

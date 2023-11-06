import AllDay from "../../contracts/AllDay.cdc"

// This script returns a Series struct for the given name,
// if it exists

pub fun main(seriesName: String): AllDay.SeriesData {
    return AllDay.getSeriesDataByName(name: seriesName)
}


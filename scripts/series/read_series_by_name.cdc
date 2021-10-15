import Showdown from "../../contracts/Showdown.cdc"

// This script returns a Series struct for the given name,
// if it exists

pub fun main(seriesName: String): Showdown.SeriesData {
    return Showdown.getSeriesDataByName(name: seriesName)
}

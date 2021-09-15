import Genies from "../../contracts/Genies.cdc"

// This script returns a Series struct for the given name,
// if it exists

pub fun main(seriesName: String): Genies.SeriesData {
    return Genies.getSeriesDataByName(name: seriesName)
}

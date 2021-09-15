import Genies from "../../contracts/Genies.cdc"

// This script returns the current series struct

pub fun main(): Genies.SeriesData {
    return Genies.getSeriesData(id: Genies.currentSeriesID)
}

import Genies from "../../contracts/Genies.cdc"

// This script returns a Series struct for the given id,
// if it exists

pub fun main(id: UInt32): Genies.SeriesData {
    return Genies.getSeriesData(id: id)
}

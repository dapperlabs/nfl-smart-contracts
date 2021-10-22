import Showdown from "../../contracts/Showdown.cdc"

// This script returns a Series struct for the given id,
// if it exists

pub fun main(id: UInt32): Showdown.SeriesData {
    return Showdown.getSeriesData(id: id)
}


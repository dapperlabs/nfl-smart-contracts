import AllDay from "../../contracts/AllDay.cdc"

// This script returns a Series struct for the given id,
// if it exists

pub fun main(id: UInt32): AllDay.SeriesData {
    return AllDay.getSeriesData(id: id)
}


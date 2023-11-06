import AllDay from "../../contracts/AllDay.cdc"

// This script returns all the names for Series.
// These can be related to Series structs via AllDay.getSeriesByName() .

pub fun main(): [String] {
    return AllDay.getAllSeriesNames()
}


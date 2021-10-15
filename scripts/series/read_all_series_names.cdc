import Showdown from "../../contracts/Showdown.cdc"

// This script returns all the names for Series.
// These can be related to Series structs via Showdown.getSeriesByName() .

pub fun main(): [String] {
    return Showdown.getAllSeriesNames()
}

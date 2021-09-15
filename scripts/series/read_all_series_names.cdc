import Genies from "../../contracts/Genies.cdc"

// This script returns all the names for Series.
// These can be related to Series structs via Genies.getSeriesByName() .

pub fun main(): [String] {
    return Genies.getAllSeriesNames()
}

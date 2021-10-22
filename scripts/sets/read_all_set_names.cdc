import AllDay from "../../contracts/AllDay.cdc"

// This script returns all the names for Set.
// These can be related to Set structs via AllDay.getSetByName() .

pub fun main(): [String] {
    return AllDay.getAllSetNames()
}


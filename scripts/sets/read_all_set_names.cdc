import Showdown from "../../contracts/Showdown.cdc"

// This script returns all the names for Set.
// These can be related to Set structs via Showdown.getSetByName() .

pub fun main(): [String] {
    return Showdown.getAllSetNames()
}


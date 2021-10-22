import Showdown from "../../contracts/Showdown.cdc"

// This script returns an Edition for an id number, if it exists.

pub fun main(editionID: UInt32): Showdown.EditionData {
    return Showdown.getEditionData(id: editionID)
}


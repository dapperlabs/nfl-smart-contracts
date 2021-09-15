import Genies from "../../contracts/Genies.cdc"

// This script returns an Edition for an id number, if it exists.

pub fun main(editionID: UInt32): Genies.EditionData {
    return Genies.getEditionData(id: editionID)
}

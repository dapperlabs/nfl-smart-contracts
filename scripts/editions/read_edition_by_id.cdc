import AllDay from "../../contracts/AllDay.cdc"

// This script returns an Edition for an id number, if it exists.

pub fun main(editionID: UInt64): AllDay.EditionData {
    return AllDay.getEditionData(id: editionID)
}


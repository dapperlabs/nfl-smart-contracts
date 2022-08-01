import AllDaySeasonal from "../../contracts/AllDaySeasonal.cdc"

// This script returns an Edition for an id number, if it exists.

pub fun main(editionID: UInt64): AllDaySeasonal.EditionData {
    return AllDaySeasonal.getEditionData(id: editionID)
}
import Showdown from "../../contracts/Showdown.cdc"

// This script returns all the Edition structs.
// This will be *long*.

pub fun main(): [Showdown.EditionData] {
    let editions: [Showdown.EditionData] = []
    var id: UInt32 = 1
    // Note < , as nextEditionID has not yet been used
    while id < Showdown.nextEditionID {
        editions.append(Showdown.getEditionData(id: id))
        id = id + 1
    }
    return editions
}


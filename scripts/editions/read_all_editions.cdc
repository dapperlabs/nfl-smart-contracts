import AllDay from "../../contracts/AllDay.cdc"

// This script returns all the Edition structs.
// This will be *long*.

pub fun main(): [AllDay.EditionData] {
    let editions: [AllDay.EditionData] = []
    var id: UInt32 = 1
    // Note < , as nextEditionID has not yet been used
    while id < AllDay.nextEditionID {
        editions.append(AllDay.getEditionData(id: id))
        id = id + 1
    }
    return editions
}


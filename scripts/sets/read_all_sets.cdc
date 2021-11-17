import AllDay from "../../contracts/AllDay.cdc"

// This script returns all the Set structs.
// This will eventually be *long*.

pub fun main(): [AllDay.SetData] {
    let sets: [AllDay.SetData] = []
    var id: UInt64 = 1
    // Note < , as nextSetID has not yet been used
    while id < AllDay.nextSetID {
        sets.append(AllDay.getSetData(id: id))
        id = id + 1
    }
    return sets
}


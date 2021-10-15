import Showdown from "../../contracts/Showdown.cdc"

// This script returns all the Set structs.
// This will eventually be *long*.

pub fun main(): [Showdown.SetData] {
    let sets: [Showdown.SetData] = []
    var id: UInt32 = 1
    // Note < , as nextSetID has not yet been used
    while id < Showdown.nextSetID {
        sets.append(Showdown.getSetData(id: id))
        id = id + 1
    }
    return sets
}

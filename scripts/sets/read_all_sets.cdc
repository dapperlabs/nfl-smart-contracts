import AllDay from "AllDay"

// This script returns all the Set structs.
// This will eventually be *long*.

access(all) fun main(): [AllDay.SetData] {
    let sets: [AllDay.SetData] = []
    var id: UInt64 = 1
    // Note < , as nextSetID has not yet been used
    while id < AllDay.nextSetID {
        sets.append(AllDay.getSetData(id: id))
        id = id + 1
    }
    return sets
}


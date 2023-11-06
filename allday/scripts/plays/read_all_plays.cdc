import AllDay from "../../contracts/AllDay.cdc"

// This script returns all the Set structs.
// This will eventually be *long*.

pub fun main(): [AllDay.PlayData] {
    let plays: [AllDay.PlayData] = []
    var id: UInt64 = 1
    // Note < , as nextPlayID has not yet been used
    while id < AllDay.nextPlayID {
        plays.append(AllDay.getPlayData(id: id))
        id = id + 1
    }
    return plays
}


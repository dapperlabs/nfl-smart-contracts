import Showdown from "../../contracts/Showdown.cdc"

// This script returns all the Set structs.
// This will eventually be *long*.

pub fun main(): [Showdown.PlayData] {
    let plays: [Showdown.PlayData] = []
    var id: UInt32 = 1
    // Note < , as nextPlayID has not yet been used
    while id < Showdown.nextPlayID {
        plays.append(Showdown.getPlayData(id: id))
        id = id + 1
    }
    return plays
}

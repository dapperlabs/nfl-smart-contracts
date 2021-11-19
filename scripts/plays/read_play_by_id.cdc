import AllDay from "../../contracts/AllDay.cdc"

// This script returns a Play struct for the given id,
// if it exists

pub fun main(id: UInt64): AllDay.PlayData {
    return AllDay.getPlayData(id: id)
}


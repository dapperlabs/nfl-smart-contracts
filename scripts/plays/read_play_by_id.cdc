import Showdown from "../../contracts/Showdown.cdc"

// This script returns a Play struct for the given id,
// if it exists

pub fun main(id: UInt32): Showdown.PlayData {
    return Showdown.getPlayData(id: id)
}


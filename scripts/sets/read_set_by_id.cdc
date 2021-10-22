import Showdown from "../../contracts/Showdown.cdc"

// This script returns a Set struct for the given id,
// if it exists

pub fun main(id: UInt32): Showdown.SetData {
    return Showdown.getSetData(id: id)
}


import AllDay from "../../contracts/AllDay.cdc"

// This script returns a Set struct for the given id,
// if it exists

pub fun main(id: UInt32): AllDay.SetData {
    return AllDay.getSetData(id: id)
}


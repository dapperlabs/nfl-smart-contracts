import AllDay from "../../contracts/AllDay.cdc"

// This script returns a Set struct for the given name,
// if it exists

pub fun main(setName: String): AllDay.SetData {
    return AllDay.getSetDataByName(name: setName)
}


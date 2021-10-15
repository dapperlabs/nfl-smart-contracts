import Showdown from "../../contracts/Showdown.cdc"

// This script returns a Set struct for the given name,
// if it exists

pub fun main(setName: String): Showdown.SetData {
    return Showdown.getSetDataByName(name: setName)
}

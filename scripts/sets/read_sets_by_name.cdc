import AllDay from "AllDay"

// This script returns a Set struct for the given name,
// if it exists

access(all) fun main(setName: String): AllDay.SetData {
    return AllDay.getSetDataByName(name: setName)
}


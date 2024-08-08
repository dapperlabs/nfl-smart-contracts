import AllDay from "AllDay"

// This script returns a Play struct for the given id,
// if it exists

access(all) fun main(id: UInt64): AllDay.PlayData {
    return AllDay.getPlayData(id: id)
}


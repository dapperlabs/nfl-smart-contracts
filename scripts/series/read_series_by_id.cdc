import AllDay from "AllDay"

// This script returns a Series struct for the given id,
// if it exists

access(all) fun main(id: UInt64): AllDay.SeriesData {
    return AllDay.getSeriesData(id: id)
}


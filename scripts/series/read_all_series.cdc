import Showdown from "../../contracts/Showdown.cdc"

// This script returns all the Series structs.
// This will eventually be *long*.

pub fun main(): [Showdown.SeriesData] {
    let series: [Showdown.SeriesData] = []
    var id: UInt32 = 1
    // Note < , as nextSeriesID has not yet been used
    while id < Showdown.nextSeriesID {
        series.append(Showdown.getSeriesData(id: id))
        id = id + 1
    }
    return series
}


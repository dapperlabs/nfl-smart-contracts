import AllDay from "../../contracts/AllDay.cdc"

// This script returns all the Series structs.
// This will eventually be *long*.

pub fun main(): [AllDay.SeriesData] {
    let series: [AllDay.SeriesData] = []
    var id: UInt32 = 1
    // Note < , as nextSeriesID has not yet been used
    while id < AllDay.nextSeriesID {
        series.append(AllDay.getSeriesData(id: id))
        id = id + 1
    }
    return series
}


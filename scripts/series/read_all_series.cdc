import Genies from "../../contracts/Genies.cdc"

// This script returns all the Series structs.
// This will eventually be *long*.

pub fun main(): [Genies.SeriesData] {
    let geniesNFTs: [Genies.SeriesData] = []
    var id: UInt32 = 0
    // Note <= , as currentSeriesID is inclusive
    while id <= Genies.currentSeriesID {
        geniesNFTs.append(Genies.getSeriesData(id: id))
        id = id + 1
    }
    return geniesNFTs
}

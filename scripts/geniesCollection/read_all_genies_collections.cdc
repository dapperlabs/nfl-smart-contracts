import Genies from "../../contracts/Genies.cdc"

// This script returns all the GeniesCollection structs.
// This will be *long*.

pub fun main(): [Genies.GeniesCollectionData] {
    let geniesNFTs: [Genies.GeniesCollectionData] = []
    var id: UInt32 = 0
    // Note < , as nextCollectionID has not yet been used
    while id < Genies.nextCollectionID{
        geniesNFTs.append(Genies.getGeniesCollectionData(id: id))
        id = id + 1
    }
    return geniesNFTs
}
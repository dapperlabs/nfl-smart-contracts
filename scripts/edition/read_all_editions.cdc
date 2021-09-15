import Genies from "../../contracts/Genies.cdc"

// This script returns all the Edition structs.
// This will be *long*.

pub fun main(): [Genies.EditionData] {
    let geniesNFTs: [Genies.EditionData] = []
    var id: UInt32 = 0
    // Note < , as nextEditionID has not hyet been used
    while id < Genies.nextEditionID {
        geniesNFTs.append(Genies.getEditionData(id: id))
        id = id + 1
    }
    return geniesNFTs
}

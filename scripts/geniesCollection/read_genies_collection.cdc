import Genies from "../../contracts/Genies.cdc"

// This script returns a GeniesCollection (not a NonFungibleToken.Collection!)
// for an id number, if it exists.

pub fun main(id: UInt32): Genies.GeniesCollectionData {
    return Genies.getGeniesCollectionData(id: id)
}

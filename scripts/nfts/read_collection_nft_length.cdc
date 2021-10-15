import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Showdown from "../../contracts/Showdown.cdc"

// This script returns the size of an account's Showdown collection.

pub fun main(address: Address): Int {
    let account = getAccount(address)

    let collectionRef = account.getCapability(Showdown.CollectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.getIDs().length
}

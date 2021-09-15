import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import Genies from "../../contracts/Genies.cdc"

// This script returns the size of an account's Genies collection.

pub fun main(address: Address): Int {
    let account = getAccount(address)

    let collectionRef = account.getCapability(Genies.CollectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.getIDs().length
}

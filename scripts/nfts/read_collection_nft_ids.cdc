import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

// This script returns an array of all the NFT IDs in an account's collection.

access(all) fun main(address: Address): [UInt64] {
    let account = getAccount(address)

    let collectionRef = getAccount(address).capabilities.borrow<&AllDay.Collection>(AllDay.CollectionPublicPath)
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.getIDs()
}


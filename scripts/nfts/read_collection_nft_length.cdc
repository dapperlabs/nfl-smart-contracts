import NonFungibleToken from "NonFungibleToken"
import AllDay from "AllDay"

// This script returns the size of an account's AllDay collection.

access(all) fun main(address: Address): Int {
    let account = getAccount(address)

    let collectionRef = getAccount(address).capabilities.borrow<&AllDay.Collection>(AllDay.CollectionPublicPath)
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.getIDs().length
}


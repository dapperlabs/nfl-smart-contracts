import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import AllDay from "../../contracts/AllDay.cdc"

// This script returns the size of an account's AllDay collection.

pub fun main(address: Address): Int {
    let account = getAccount(address)

    let collectionRef = account.getCapability(AllDay.CollectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")
    
    return collectionRef.getIDs().length
}

